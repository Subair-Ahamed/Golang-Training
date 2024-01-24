package main

import (
	"fmt"
	"time"
	"log"
	"os"
)

func main() {
	//sleep:
	fmt.Println("Starting the golang task...")
    
	var s int
	fmt.Println("Enter the time in secs to complete the task:")
	fmt.Scanf("%d",&s)

	time.Sleep(time.Duration(s) * time.Second) //delays the execution for the given time
	fmt.Printf("Task completed after sleeping for %d seconds\n",s)

    //Log:
	//log.New() to create a logger that writes to standard output (STDOUT) with the prefix INFO with date or time.
	logger := log.New(os.Stdout,"INFO:",log.Ldate|log.Ltime) 

	// Print to the default logger
	logger.Print("This is a log message")

	// Print with a prefix and a newline
	logger.Println("Another log message")

	// Fatal will log the message and exit the program with a non-zero status
	logger.Fatal("Fatal log message")

	// Panic will log the message and cause a panic
	logger.Panic("Panic log message")
}



