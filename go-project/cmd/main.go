package main

import (
    "fmt"
    "go-project/pkg/mypackage"
    "go-project/api/routes"
)

func main() {
    fmt.Println("Starting my go project...")

    // Using functions from the pkg package
    mypackage.MyFunction()

    // Start API server
    routes.StartServer()
}
