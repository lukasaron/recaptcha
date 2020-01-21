// Recaptcha package defines the basic functionality for verification of Recaptcha tokens generated on the
// client site (front-end).
//
// To make this project independent to any web server there is not an executable file/script, but only the library is
// performing the basic evaluation of provided token and result is captured in the VerifyResponse structure.
//
// User of this library can define endpoints as the project requires. The only one required task is to define one
// environmental variable `SECRET_KEY` which has to hold the generated secret key value.
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
//			vr, err := recaptcha.VerifyToken(token, remoteIP)
//			fmt.Printf("verifyResponse: %+v\n, error: %v", vr, err)
//		}
package recaptcha
