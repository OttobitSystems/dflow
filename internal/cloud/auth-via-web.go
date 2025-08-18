// Package auth is responsible to work with cloud
package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)

var (
	authCodeChan     = make(chan string)
	serviceName      = "dflow"
	clientID         = "6l1pc4n102thk4ocsfcl4av77g"
	cognitoDomain    = "https://eu-south-1l5w4uze0p.auth.eu-south-1.amazoncognito.com"
	redirectedURI    = "http://localhost:5000/signedin"
	signinAuth       = "NmwxcGM0bjEwMnRoazRvY3NmY2w0YXY3N2c6MWliaGVsMHAzdGQ3ZDUxMTMzZ2dmOG9panB2ZmE2ZzhpOWQ5dGJrODNiNGtzNnNoaTBlaw=="
	cognitoURI       = "https://eu-south-1l5w4uze0p.auth.eu-south-1.amazoncognito.com/login?client_id=6l1pc4n102thk4ocsfcl4av77g&response_type=code&scope=aws.cognito.signin.user.admin+dflow-auth%2Fread&redirect_uri=http%3A%2F%2Flocalhost%3A5000%2Fsignedin"
	userLogedInCloud = false
)

func RefreshSession() string {
	accessToken, err := retriveTokenFromMemory()
	if err != nil {
		return ""
	}

	userLogedInCloud = true

	fmt.Println(accessToken)

	_, err = refreshSession(*accessToken)
	if err != nil {
		fmt.Println("Cloud session expired, please retry login using `dflow auth` command")
		clearTokenInMemory()
	}

	return accessToken.AccessToken
}

func LoginWeb() {
	if userLogedInCloud {
		fmt.Println("User already logged in.")
		return
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", cognitoURI)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", cognitoURI)
	case "darwin": // macOS
		cmd = exec.Command("open", cognitoURI)
	default:
		panic("Os not supported yet: " + runtime.GOOS)
	}
	cmd.Start()

	http.HandleFunc("/signedin", signedInHandler)

	go func() {
		err := http.ListenAndServe(":5000", nil)
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("Login page has been showed in browser, please sign in here and return here when you finished")
	authCode := <-authCodeChan

	_, err := getAuthToken(authCode)
	if err != nil {
		panic(err)
	}
	userLogedInCloud = true
}

func signedInHandler(w http.ResponseWriter, r *http.Request) {
	authCodeFromURL := r.URL.Query().Get("code")

	authCodeChan <- authCodeFromURL
}

func getAuthToken(authCode string) (*TokenResponse, error) {
	ctx := context.Background()

	tokenURL := fmt.Sprintf("%s/oauth2/token", cognitoDomain)

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", clientID)
	data.Set("code", authCode)
	data.Set("redirect_uri", redirectedURI)

	req, err := http.NewRequestWithContext(ctx, "POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+signinAuth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error form http: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("cognito error (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var tokenResponse TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, fmt.Errorf("reponse deserialization err: %w", err)
	}

	storeTokenInMemory(tokenResponse)

	return &tokenResponse, nil
}

func refreshSession(token TokenResponse) (*TokenResponse, error) {
	ctx := context.Background()

	tokenURL := fmt.Sprintf("%s/oauth2/token", cognitoDomain)

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("client_id", clientID)
	data.Set("refresh_token", token.RefreshToken)

	req, err := http.NewRequestWithContext(ctx, "POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+signinAuth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error form http: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("cognito error (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var tokenResponse TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, fmt.Errorf("reponse deserialization err: %w", err)
	}

	fmt.Println("at " + tokenResponse.AccessToken)
	fmt.Println("ei " + string(tokenResponse.ExpiresIn))
	fmt.Println("it " + tokenResponse.IDToken)
	fmt.Println("rt " + tokenResponse.RefreshToken)
	fmt.Println("tt " + tokenResponse.TokenType)

	storeTokenInMemory(tokenResponse)

	return &tokenResponse, nil
}
