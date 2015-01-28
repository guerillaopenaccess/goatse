package main

import (
	"github.com/robertkrimen/otto"
	"io/ioutil"
)

// Like the nodeJS require. Scope is not clean. Returns module.exports.
// Module and module.exports are objects.
// :SIGJS: (filename string) -> (returnvalue Dynamic||Null)
func goatsejs_requirejs(call otto.FunctionCall) otto.Value {
	var (
		file_name     string
		file_contents []byte
		source        string
		script        *otto.Script
		retval        otto.Value
		err           error
	)
	file_name, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	file_contents, err = ioutil.ReadFile(file_name)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	source = `var module={}; module.exports={};new function(){` + string(file_contents) + `;return module.exports}()`
	script, err = call.Otto.Compile("", source)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	retval, err = call.Otto.Run(script)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, retval, otto.Value{})
}
