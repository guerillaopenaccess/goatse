// HTTP for Otto VM in Go. 1 function and args is easy to embed. Use method name and JSON args.
package httpplus

import (
	"golang.org/x/net/publicsuffix"
	"net/http"
	"net/http/cookiejar"
	"time"
)

// timeout_secs == 0 means no timeout.
// Cookies not working.
func makeClient(timeout_secs int64, cookies bool) (client *http.Client, err error) {
	var (
		cookieArg http.CookieJar
	)
	if cookies {
		cookieArg, err = cookiejar.New(&cookiejar.Options{publicsuffix.List})
		if err != nil {
			return nil, err
		}
	} else {
		cookieArg = nil
	}
	client = &http.Client{
		Jar:     cookieArg,
		Timeout: time.Second * time.Duration(timeout_secs),
	}
	return client, nil
}
