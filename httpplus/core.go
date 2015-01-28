package httpplus

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	defaultUserAgent string
	defaultClient    *http.Client
)

func init() {
	defaultUserAgent, _ = randomUserAgent()
	defaultClient, _ = makeClient(0, false)
}

func HTTPMethod(method, uri string, args *HttpArgs) (response *http.Response, body string, err error) {
	var (
		client    *http.Client
		bytesbody []byte
		request   *http.Request
	)
	if args == nil {
		args = &HttpArgs{}
	}

	if args.Timeout_seconds > 0 || args.Cookies {
		client, err = makeClient(args.Timeout_seconds, args.Cookies)
		if err != nil {
			return nil, "", err
		}
	} else {
		client = defaultClient
	}

	request, err = buildRequest(method, uri, args)
	if err != nil {
		return nil, "", err
	}

	if args.Headers["User-Agent"] == "" {
		request.Header.Set("User-Agent", defaultUserAgent)
	}

	response, err = client.Do(request)
	if err != nil {
		return nil, "", err
	}
	defer response.Body.Close()

	bytesbody, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, "", err
	}

	body = string(bytesbody)
	return response, body, nil
}

type HttpArgs struct {
	// Time for download.
	Timeout_seconds int64

	Cookies bool

	// If set ignore other Post args
	PostBody string

	// MIME name for file
	PostFileByNameName string
	// File upload with MIME
	PostFileByName string
	// MIME params ± File
	PostBodyParams map[string]string

	UrlParams map[string][]string

	// Headers. Can override default User-Agent. User-Agent is already fake.
	Headers map[string]string

	// Must be set with password
	BasicAuthUser     string
	// Must be set with user
	BasicAuthPassword string
}

func (a *HttpArgs) hasBasicAuth() bool {
	if a.BasicAuthUser != "" && a.BasicAuthPassword != "" {
		return true
	} else {
		return false
	}
}

// For when users have url params as a map[string]string but requires map[string][]string.
func (a *HttpArgs) AddQueryMap(m map[string]string) {
	for k, v := range m {
		a.UrlParams[k] = []string{v}
	}
}

// switch methods to functions
func buildRequest(method, uri string, args *HttpArgs) (req *http.Request, err error) {
	switch method {
	case "HEAD": {
			return buildGETorHEAD("HEAD", uri, args)
		}
	case "GET": {
			return buildGETorHEAD("GET", uri, args)
		}
	case "POST": {
			req, err = buildPOST(uri, args)
			if err != nil {
				return nil, err
			}
			if args.hasBasicAuth() {
				fmt.Println("Setting Basic Authentication: " + args.BasicAuthUser + args.BasicAuthPassword)
				req.SetBasicAuth(args.BasicAuthUser, args.BasicAuthPassword)
			}
			setHeaders(req, args.Headers)
			return req, nil
		}
	default: {
			return nil, fmt.Errorf("Invalid or unsupported method: %s", method)
		}
	}
	return nil, fmt.Errorf("Switch bypassed somehow, this shouldn't be possible.")
}

func buildGETorHEAD(method, uri string, args *HttpArgs) (req *http.Request, err error) {
	req, err = http.NewRequest(method, uriWithParams(uri, args), nil)
	if err != nil {
		return nil, err
	}
	if args.hasBasicAuth() {
		fmt.Println("Setting Basic Authentication: " + args.BasicAuthUser + args.BasicAuthPassword)
		req.SetBasicAuth(args.BasicAuthUser, args.BasicAuthPassword)
	}
	setHeaders(req, args.Headers)
	return req, nil
}

func buildPOST(uri string, args *HttpArgs) (req *http.Request, err error) {
	if args.PostBody != "" {
		return http.NewRequest("POST", uriWithParams(uri, args), strings.NewReader(args.PostBody))
	}
	if args.PostFileByName != "" {
		if args.PostFileByNameName == "" {
			args.PostFileByNameName = "files"
		}
		return newfileUploadRequest(uri, args.PostBodyParams, args.PostFileByNameName, args.PostFileByName)
	}
	if len(args.PostBodyParams) > 0 {
		return makeMultiparamsPOST(uri, args.PostBodyParams)
	}
	return http.NewRequest("POST", uriWithParams(uri, args), nil)
}

func setHeaders(req *http.Request, headers map[string]string) (err error) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return nil
}

func uriWithParams(uri string, args *HttpArgs) (newUri string) {
	var (
		vals url.Values
	)
	if len(args.UrlParams) > 0 {
		vals = url.Values(args.UrlParams)
		newUri = strings.TrimRight(uri, "?") + "?" + vals.Encode()
	} else {
		newUri = uri
	}
	return newUri
}
