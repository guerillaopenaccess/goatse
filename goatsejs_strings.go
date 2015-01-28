package main

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"net/url"
	"strings"
)

func _get_str_args(call *otto.FunctionCall) (strarg1, strarg2, strarg3 string, err error) {
	strarg1, err = call.Argument(0).ToString()
	if err != nil {
		return "", "", "", err
	}
	strarg2, err = call.Argument(1).ToString()
	if err != nil {
		return "", "", "", err
	}
	// argstring empty ok
	strarg3, _ = call.Argument(2).ToString()
	return strarg1, strarg2, strarg3, err
}

// :SIGJS: (string) -> string
func goatsejs_strings_urlescape(call otto.FunctionCall) otto.Value {
	var (
		strarg string
		retstr string
		retval otto.Value
		err    error
	)
	strarg, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	retstr = url.QueryEscape(strarg)
	retval, err = call.Otto.ToValue(retstr)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, retval, otto.Value{})
}

// :SIGJS: (string) -> string
func goatsejs_strings_urlunescape(call otto.FunctionCall) otto.Value {
	var (
		strarg string
		retstr string
		retval otto.Value
		err    error
	)
	strarg, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	retstr, err = url.QueryUnescape(strarg)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	retval, err = call.Otto.ToValue(retstr)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, retval, otto.Value{})
}

// :SIGJS: (funcname string, mainstring, argstring1, argstring2) -> string
func goatsejs_strings_entry(call otto.FunctionCall) otto.Value {
	var (
		funcname   string
		mainstring string
		argstring  string
		argrune    rune
		retstring  string
		retval     otto.Value
		err        error
	)
	funcname, mainstring, argstring, err = _get_str_args(&call)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	switch funcname {
	case "contains":
		{
			retstring = fmt.Sprintf("%t", strings.Contains(mainstring, argstring))
		}
	case "containsAny":
		{
			retstring = fmt.Sprintf("%t", strings.ContainsAny(mainstring, argstring))
		}
	case "containsRune":
		{
			argrune, _, err = strings.NewReader(argstring).ReadRune()
			if err != nil {
				return returnErrCapsule(call.Otto, err)
			}
			retstring = fmt.Sprintf("%t", strings.ContainsRune(mainstring, argrune))
		}
	case "count":
		{
			retstring = fmt.Sprintf("%d", strings.Count(mainstring, argstring))
		}
	case "equalFold":
		{
			retstring = fmt.Sprintf("%t", strings.EqualFold(mainstring, argstring))
		}
	case "hasPrefix":
		{
			retstring = fmt.Sprintf("%t", strings.HasPrefix(mainstring, argstring))
		}
	case "hasSuffix":
		{
			retstring = fmt.Sprintf("%t", strings.HasSuffix(mainstring, argstring))
		}
	case "indexRune":
		{
			argrune, _, err = strings.NewReader(argstring).ReadRune()
			if err != nil {
				return returnErrCapsule(call.Otto, err)
			}
			retstring = fmt.Sprintf("%d", strings.IndexRune(mainstring, argrune))
		}
	case "lastIndex":
		{
			retstring = fmt.Sprintf("%d", strings.LastIndex(mainstring, argstring))
		}
	case "lastIndexAny":
		{
			retstring = fmt.Sprintf("%d", strings.LastIndexAny(mainstring, argstring))
		}
	case "title":
		{
			retstring = strings.Title(mainstring)
		}
	case "trim":
		{
			retstring = strings.Trim(mainstring, argstring)
		}
	case "trimSpace":
		{
			retstring = strings.TrimSpace(mainstring)
		}
	case "trimLeft":
		{
			retstring = strings.TrimLeft(mainstring, argstring)
		}
	case "trimRight":
		{
			retstring = strings.TrimRight(mainstring, argstring)
		}
	case "trimPrefix":
		{
			retstring = strings.TrimPrefix(mainstring, argstring)
		}
	case "trimSuffix":
		{
			retstring = strings.TrimSuffix(mainstring, argstring)
		}
	}
	retval, err = call.Otto.ToValue(retstring)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, retval, otto.Value{})
}
