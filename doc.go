// Package recaptcha defines the functionality for verification of Recaptcha tokens (all versions)
// generated on the client site.
//
// To make this project independent to any web server there is not an executable file/script, but only the library is
// performing the basic evaluation of provided token and result is captured in the Response structure.
//
// User of this library can define endpoints as the project requires. The only one required task is to define one
// environmental variable SECRET_KEY which has to hold the generated secret key value.
//
// Example of basic usage:
//		package main
//
//		import (
//			"fmt"
//			"github.com/lukasaron/recaptcha"
//		)
//
//		func main() {
//			token := "--- accept the token from the client side--- "
//			remoteIP := "--- fill the remote IP or leave empty --- "
//			r, err := recaptcha.VerifyToken(token, remoteIP)
//			fmt.Printf("Response: %+v\n, error: %v", r, err)
//		}
package recaptcha
