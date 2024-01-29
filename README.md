# Golang-Training

# Dependency Management

# 1.Getting started with Modules:
		
  -> To get started using Modules in your project, you can simply enter the following command:
	
	         "go mod init <module-path>" or simply "go mod init"
	
# 2.Installing dependencies:

	-> To install dependencies, use the "go get" command, which will also update the go.mod file automatically:
		
            "go get <specified-path>"
              
# 3.Authenticating dependencies:

	-> The go.sum file acts as a dependency checker that authenticates your modules against unexpected or malicious changes that may break your entire codebase.
	
# 4.Removing unused dependencies:

	-> After running the following command, it’ll remove any unused dependencies in your project and update the go.mod file.Remember that the go.sum file will still contain the cryptographic hash of the package’s content even after you’ve removed it. 
	
	          "go mod tidy"


# 5.Installing missing dependencies:

	-> We can run "go build" or "go test" command to install all the missing dependencies automatically.
	
# 6.Upgrading dependency versions:

	   6.1.Download and update to a specific version:

  	          "go get <pathname@versiom>"
  	     
  	 6.2.The -u flag implies that the latest minor or patch releases are to be used):

              "go get -u <pathname>"
                 
# 7.Listing dependencies:

	-> To see a list of the current module and all its dependencies, run the following command:

              "go list -m all" or "go list -m <module-name>"   
