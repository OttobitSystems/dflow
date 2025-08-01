package main

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type LogModel struct {
	gorm.Model
	LogText string
}

func log_text(textToStore string) {
	fmt.Println("Logging ", textToStore)
	db, err := gorm.Open(sqlite.Open("dflow.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&LogModel{})

	db.Create(&LogModel{LogText: textToStore})
}

func listLogs() {
	fmt.Println("showing logs..")

	db, err := gorm.Open(sqlite.Open("dflow.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&LogModel{})

	var logs []LogModel

	result := db.Find(&logs)

	if result.Error != nil {
		log.Fatal(err)
	}

	fmt.Println("count: ", len(logs))

	for _, logValue := range logs {
		fmt.Println(logValue.CreatedAt, " ", logValue.LogText)
	}

}

var rootCmd = &cobra.Command{
	Use:   "dflow",
	Short: "An flow application to follow the flow",
}

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Logs some text",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			log_text(args[0])
		} else {
			log.Fatal("args != 1, count: ", len(args))
		}
	},
}

var showLogsCmd = &cobra.Command{
	Use:   "list-logs",
	Short: "Shows log in database",
	Run: func(cmd *cobra.Command, args []string) {
		listLogs()
	},
}

func init() {
	rootCmd.AddCommand(logCmd, showLogsCmd)
}

func main() {
	err := rootCmd.Execute()

	if err != nil {
		log.Fatal(err)
	}

	/*
	   fmt.Println("Hello")
	   db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	   	if err != nil {
	   		log.Fatal(err)
	   	}

	   db.AutoMigrate(&TestModel{})

	   db.Create(&TestModel{SomeText: "SomeText"})
	*/
}
