package main

import (
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"crypto/md5"
	"encoding/hex"
	"os"
)

// :SIGJS: (filepath string, filecontents string)->(result string)
// :SIGJS:TYPE: filecontents: binary string
func goatsejs_saveToFile(call otto.FunctionCall) otto.Value {
	var (
		file_name     string
		file_contents string
		err           error
	)
	file_name, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	file_contents, err = call.Argument(1).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	err = ioutil.WriteFile(file_name, []byte(file_contents), 0644)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, otto.Value{}, otto.Value{})
}

// :SIGJS: (filepath string)->(filecontents string)
// :SIGJS:TYPE: filecontents: binary string
func goatsejs_loadFromFile(call otto.FunctionCall) otto.Value {
	var (
		file_name     string
		file_contents []byte
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
	retval, err = vm.ToValue(string(file_contents))
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, retval, otto.Value{})
}

// :SIGJS: (filepath string)->(filesize int)
func goatsejs_fileSize(call otto.FunctionCall) otto.Value {
	var (
		file_name     string
		file_stat os.FileInfo
		file_size int64
		retval        otto.Value
		err           error
	)
	file_name, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	file_stat, err = os.Stat(file_name)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	file_size = file_stat.Size()
	retval, err = vm.ToValue(file_size)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, retval, otto.Value{})
}

// :SIGJS: (filepath string)->(md5sum string)
func goatsejs_md5SumFile(call otto.FunctionCall) otto.Value {
	var (
		file_name     string
		file_contents []byte
		file_md5	 [16]byte
		hexmd5			string
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
    file_md5 = md5.Sum(file_contents)
    hexmd5 = hex.EncodeToString(file_md5[:])
	retval, err = vm.ToValue(hexmd5)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, retval, otto.Value{})
}
