package main

import (
	"fmt"
	"os"
)

func main() {

	os.Setenv("GOPATH","2023") //Setenv -> to set a value
	goPath := os.Getenv("GOPATH") //Getenv -> to retrieve the value
	fmt.Println("GOPATH:", goPath) //Specifies the location of your Go workspace. This includes the src, bin, and pkg directories. If not set, Go assumes a default location ($HOME/go in Unix systems)
	
	shell := os.Getenv("SHELL")
	fmt.Println("Shell:", shell) //prints the name of the current user shell.

	goBin := os.Getenv("GOBIN")
	fmt.Println("GOBIN:", goBin)			

	goRoot := os.Getenv("GOROOT")
	fmt.Println("GOROOT:", goRoot)	

	path := os.Getenv("PATH")
	fmt.Println("PATH:", path)

	//os.LookupEnv:
	getEnv := func(key string) {
        val, ok := os.LookupEnv(key) //retrieves the value of the environment variable named by the key
        if !ok {
            fmt.Printf("%s not set\n", key)
        } else {
            fmt.Printf("%s=%s\n", key, val)
        }
        }
        getEnv("EDITOR") //not set
        getEnv("SHELL") //prints the value
    
	//os.ExpandEnv:
	os.Setenv("EDITOR", "emacs")
	expand := os.ExpandEnv("My editor is $EDITOR.") //replaces the '$var' inside a string into the value of the given variable
        fmt.Println(expand)       
        fmt.Println(os.ExpandEnv("My shell is $SHELL."))

	//Example-1:
	var a string
	fmt.Println("Enter the value:")
	fmt.Scanf("%v",&a)

	os.Setenv("MY_VARIABLE","a")
	value := os.Getenv("MY_VARIABLE")

	if value == "" {
		fmt.Println("MY_VARIABLE is not set.")
	} else {
		fmt.Println("Value of MY_VARIABLE:", a)
	}
    
}

/* 
1. GOPATH: Specifies the location of your Go workspace. This includes the src, bin, and pkg directories. 
           If not set, Go assumes a default location ($HOME/go in Unix systems).

2. GOBIN: Defines the directory where Go executables (binary files) are installed when you run go install.

3. GOOS and GOARCH: These specify the target operating system and architecture, respectively. 
                    For ex, GOOS=linux and GOARCH=amd64 would compile code for Linux on the AMD64 architecture.

4. GOROOT: Specifies the root directory where Go is installed. This variable is set by the Go installation itself 
           and doesn’t need to be set manually in most cases.

5. GO111MODULE: This variable sets the module behavior for Go 1.11 and higher. 
                Setting it to on enables the use of Go modules for dependency management.

6. CGO_ENABLED: Used to enable or disable CGO (C Go) support. 
                Setting it to 0 disables CGO, while setting it to 1 enables it.

7. PATH: Should include the Go binary directory ($GOPATH/bin or $GOBIN) to run Go executables without specifying their full paths. */
