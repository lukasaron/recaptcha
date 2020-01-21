# reCaptcha

[![GoDoc](https://godoc.org/github.com/lukasaron/recaptcha?status.svg)](https://godoc.org/github.com/lukasaron/recaptcha)
[![Build Status](https://travis-ci.com/lukasaron/recaptcha.svg?branch=master)](https://travis-ci.com/lukasaron/recaptcha)
[![Go Report Card](https://goreportcard.com/badge/github.com/lukasaron/recaptcha)](https://goreportcard.com/report/github.com/lukasaron/recaptcha)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

reCaptcha defines the functionality for verification of reCaptcha tokens (all version) generated on the client site.

To make this project independent to any web server there is not an executable file/script, but only the library is
performing the basic evaluation of provided token and result is captured in the VerifyResponse structure.

User of this library can define endpoints as the project requires. The only one required task is to define one
environmental variable `SECRET_KEY` which has to hold the generated secret key value.

## Installation
```go
go get module github.com/lukasaron/recaptcha
```

## Example of usage
```go
package main

import (
	"fmt"
	"github.com/lukasaron/recaptcha"
)

func main() {
    token := "--- accept the token from the client side--- "
    remoteIP := "--- fill the remote IP or leave empty --- "
	r, err := recaptcha.VerifyToken(token, remoteIP)
	fmt.Printf("Response: %+v\n, error: %v", r, err)
}
```
