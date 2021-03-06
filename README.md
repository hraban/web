#  web.go

web.go is the simplest way to write web applications in the Go programming
language. It's ideal for writing simple, performant backend web services. 


## Features

web.go should be familiar to people who've developed websites with higher-level
web frameworks like sinatra or web.py. It is designed to be a lightweight web
framework that doesn't impose any scaffolding on the user. Some features
include:

* Routing to url handlers based on regular expressions
* Secure cookies
* Support for fastcgi and scgi
* Web applications are compiled to native code. This means very fast execution
  and page render speed
* Efficiently serving static files
* TLS server support
* automatic marshaling of JSON and XML content

(A security expert should probably sign off on the TLS support before it becomes
a standard.)

## Installation

Make sure you have the a working Go environment. See the [install
instructions](http://golang.org/doc/install.html).

To install web.go, simply run:

    go get github.com/hraban/web

Note that you need to set the environment variable GOPATH before using go
get. 

## Example
    
```go
    package main

    import (
        "github.com/hraban/web"
    )

    func hello(val string) string {
        return "hello " + val
    } 

    func main() {
        web.Get("/(.*)", hello)
        web.Run("0.0.0.0:9999")
    }
```

To run the application, put the code in a file called hello.go and run:

    go run hello.go

You can point your browser to http://localhost:9999/world . 

See http://www.godoc.org/github.com/hraban/web for a tutorial.  There are more
example applications in the examples directory.
