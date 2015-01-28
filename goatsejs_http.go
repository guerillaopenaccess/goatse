// HTTP interface to goatse.
package main

import (
	"github.com/guerillaopenaccess/goatse/httpplus"
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	"net/http"
	"strings"
)

// Interface httpplus.
// :SIGJS: (method string, uri string, HttpArg string) -> (result object)
// :SIGJS:TYPE: HTTPArg: httpplus.HttpArgs struct, as JSON.
// :SIGJS:TYPE: Result: {body: string, status: string, statusCode: int, contentLength: int, header: {string:string}, trailer: {string:string}}
func goatsejs_httpmethod(call otto.FunctionCall) (retval otto.Value) {
	var (
		method         string
		uri            string
		hplus_args     *httpplus.HttpArgs
		response       *http.Response
		body           string
		responseObject *otto.Object
		responseValue  otto.Value
		err            error
	)
	method, uri, hplus_args, err = goatsejs_http_getargs(&call)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	response, body, err = httpplus.HTTPMethod(method, uri, hplus_args)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	responseObject, err = goatsejs_http_construct_return(call.Otto, response, body)
	responseValue = responseObject.Value()
	return returnCapsule(call.Otto, responseValue, otto.Value{})
}

// Main keys of http.Response to JS object.
func goatsejs_http_construct_return(ovm *otto.Otto, resp *http.Response, body string) (respObj *otto.Object, err error) {
	var (
		headerObj  *otto.Object
		trailerObj *otto.Object
	)
	respObj, _ = ovm.Object("({})")
	respObj.Set("body", body)
	respObj.Set("status", resp.Status)
	respObj.Set("statusCode", resp.StatusCode)
	respObj.Set("contentLength", resp.ContentLength)
	respObj.Set("url", resp.Request.URL.String())
	headerObj, _ = ovm.Object("({})")
	for k, v := range resp.Header {
		headerObj.Set(k, v)
	}
	respObj.Set("header", headerObj)
	trailerObj, _ = ovm.Object("({})")
	for k, v := range resp.Trailer {
		trailerObj.Set(k, v)
	}
	respObj.Set("trailer", trailerObj)
	return respObj, nil
}

// get args from otto call.
func goatsejs_http_getargs(call *otto.FunctionCall) (method, uri string, hplus_args *httpplus.HttpArgs, err error) {
	var (
		args_json       string
		_getargs_errout func(string) (string, string, *httpplus.HttpArgs, error)
	)

	_getargs_errout = func(str string) (string, string, *httpplus.HttpArgs, error) {
		return "", "", nil, fmt.Errorf(str)
	}

	hplus_args = &httpplus.HttpArgs{}

	method = call.Argument(0).String()
	if method == "" {
		return _getargs_errout("Method (string) not given.")
	}
	method = strings.TrimSpace(strings.ToUpper(method))
	if !(method == "HEAD" || method == "GET" || method == "POST") {
		return _getargs_errout("Bad HTTP method; must be GET/POST/HEAD: " + method)
	}
	uri = call.Argument(1).String()
	if uri == "" {
		return _getargs_errout("URI (string) not given.")
	}
	args_json = call.Argument(2).String()
	if (args_json != "") && (args_json != "undefined") {
		err = json.Unmarshal([]byte(args_json), hplus_args)
		if err != nil {
			return _getargs_errout("Failed to parse HttpArgs object to struct: " + err.Error() + " - JSON was: " + args_json)
		}
	} else {
		hplus_args = nil
	}
	return method, uri, hplus_args, nil
}
