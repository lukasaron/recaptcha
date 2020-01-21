package recaptcha

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	contentType = "application/json"
)

var (
	VerifyURL, _ = url.Parse("https://www.google.com/recaptcha/api/siteverify")
)

type VerifyResponse struct {
	Success            bool      `json:"success"`          // status of the verification
	TimeStamp          time.Time `json:"challenge_ts"`     // timestamp of the challenge load (ISO format)
	HostName           string    `json:"hostname"`         // the hostname of the site where the reCAPTCHA was solved
	Score              float32   `json:"score"`            // the score for this request (0.0 - 1.0)
	Action             string    `json:"action"`           // the action name for this request
	ErrorCodes         []string  `json:"error-codes"`      // error codes
	AndroidPackageName string    `json:"apk_package_name"` // android related only
}

func VerifyToken(token, remoteIP string) (VerifyResponse, error) {
	q := VerifyURL.Query()
	q.Add("secret", os.Getenv("SECRET_KEY"))
	q.Add("response", token)
	q.Add("remoteip", remoteIP)
	VerifyURL.RawQuery = q.Encode()

	vr := VerifyResponse{}
	r, err := http.Post(VerifyURL.String(), contentType, nil)
	if err != nil {
		return vr, err
	}

	b, err := ioutil.ReadAll(r.Body)
	_ = r.Body.Close() // close immediately after reading finished
	if err != nil {
		return vr, err
	}

	return vr, json.Unmarshal(b, &vr)
}
