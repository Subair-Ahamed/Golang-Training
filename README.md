Package name should be:

    good: short, concise, evocative.
    by convention, packages are given lower case, single-word names;
    no need for underscores or mixedCaps
    common practise is to use package name as the base name of its source directory

Another convention is that the package name is the base name of its source directory; the package in src/encoding/base64 is imported as "encoding/base64" but has name base64, not encoding_base64 and not encodingBase64.

Interface naming:

    Use a noun or noun phrase as the name
    One-method interfaces are named by the method name plus an -er suffix or similar modification to construct a noun. Eg. Reader, Writer, etc.
    use MixedCaps or mixedCaps rather than underscores to write multiword names.
    Use prefix I Eg. IUser, IRead, etc. One con - searching for interfaces across codebase is harder due to I
    Use suffix Interface. Eg. UserInterface, ReadInterface, etc.

File Naming:

    Generally, file names are single lowercase words.
    Go follows a convention where source files are all lower case with underscore separating multiple words.
    Compound file names are separated with _
    File names that begin with “.” or “_” are ignored by the go tool
    Test files in Go come with suffix _test.go . These are only compiled and run by the go test tool.
    Files with os and architecture specific suffixes automatically follow those same constraints, e.g. name_linux.go will only build on linux, name_amd64.go will only build on amd64.

Functions and Variables:

    Use camel case, exported functions and variables start with uppercase. (unexported/private are lowercase)
    Constant should use all capital letters and use underscore _ to separate words. Eg. INT_MAX
    If variable type is bool, its name should start with Has, Is, Can or Allow, etc.
    Avoid cryptic abbreviations. Eg: usrAge := 25

REST-URLS:

    -> Singleton and Collection Resources:
    
        A resource can be a singleton or a collection.
        For example, “customers” is a collection resource and “customer” is a singleton resource (in a banking domain).
        We can identify “customers” collection resource using the URI “/customers“. We can identify a single “customer” resource using the URI “/customers/{customerId}“.
        
                    /customers			//is a collection resource
                    
                    /customers/{id}		// is a singleton resource
                
    -> Collection and Sub-collection Resources
    
    A resource may contain sub-collection resources also.
    For example, sub-collection resource “accounts” of a particular “customer” can be identified using the URN “/customers/{customerId}/accounts” (in a banking domain).
    Similarly, a singleton resource “account” inside the sub-collection resource “accounts” can be identified as follows: “/customers/{customerId}/accounts/{accountId}“.
    
                    /customers						//is a collection resource
                    
                    /customers/{id}/accounts		// is a sub-collection resource
