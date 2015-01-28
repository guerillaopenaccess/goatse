package main

import (
	"github.com/robertkrimen/otto"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Sleeps for X seconds
// :SIGJS: (time_ms int) -> Null
func goatsejs_sleep(call otto.FunctionCall) otto.Value {
	var (
		msarg int64
		err   error
	)
	msarg, err = call.Argument(0).ToInteger()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	time.Sleep(time.Duration(msarg) * time.Millisecond)
	return returnCapsule(call.Otto, otto.Value{}, otto.Value{})
}

// Return random number 0->n
// :SIGJS: (int)->int
func goatsejs_randInt(call otto.FunctionCall) otto.Value {
	var (
		upp    int64
		retint otto.Value
		err    error
	)
	upp, err = call.Argument(0).ToInteger()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	retint, err = call.Otto.ToValue(rand.Intn(int(math.Abs(float64(upp)))))
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, retint, otto.Value{})
}

// Logs to stderr
// :SIGJS: (message string) -> Null
func goatsejs_log(call otto.FunctionCall) otto.Value {
	var (
		logmsg string
		err    error
	)
	logmsg, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	log.Print(logmsg)
	return returnCapsule(call.Otto, otto.Value{}, otto.Value{})
}

// Returns JS array containing os.Args
// :SIGJS: () -> (argv [string])
func goatsejs_argv(call otto.FunctionCall) otto.Value {
	var (
		outval otto.Value
		err    error
	)
	outval, err = call.Otto.ToValue(os.Args)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, outval, otto.Value{})
}

// Returns an environment variable
// :SIGJS: (envvar string) -> (envval string)
func goatsejs_getenv(call otto.FunctionCall) otto.Value {
	var (
		arg      string
		variable string
		outval   otto.Value
		err      error
	)
	arg, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	variable = os.Getenv(arg)
	outval, err = call.Otto.ToValue(variable)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, outval, otto.Value{})
}

// exec call returning stdout
// :SIGJS: (command string, arguments [string]) -> (output string)
func goatsejs_exec(call otto.FunctionCall) otto.Value {
	var (
		command string
		args    []string
		strarg  string
		out     []byte
		outstr  string
		outval  otto.Value
		err     error
	)
	command, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	args = []string{}
	for _, arg := range call.ArgumentList[1:] {
		strarg = arg.String()
		args = append(args, strarg)
	}
	out, err = exec.Command(command, args...).Output()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	outstr = string(out)
	outval, err = call.Otto.ToValue(outstr)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, outval, otto.Value{})
}

// Returns path to command
// :SIGJS: (command string) -> abspath string
func goatsejs_where(call otto.FunctionCall) otto.Value {
	var (
		command string
		path    string
		outval  otto.Value
		err     error
	)
	command, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	path, err = exec.LookPath(command)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	outval, err = call.Otto.ToValue(path)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, outval, otto.Value{})
}

// does file/path exist
func _exists(path string) (bool, error) {
	var err error
	_, err = os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Test if file exist
// :SIGJS: (filename string) -> bool
func goatsejs_exists(call otto.FunctionCall) otto.Value {
	var (
		filepath string
		exists   bool
		ret      otto.Value
		err      error
	)
	filepath, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	exists, err = _exists(filepath)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	ret, err = call.Otto.ToValue(exists)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, ret, otto.Value{})
}

// Rename or move
// :SIGJS: (pathname string) -> undefined
func goatsejs_rename(call otto.FunctionCall) otto.Value {
	var (
		src string
		dst string
		err error
	)
	src, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	dst, err = call.Argument(1).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	err = os.Rename(src, dst)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, otto.Value{}, otto.Value{})
}

// Make a path.
// :SIGJS: (path string) -> undefined
func goatsejs_mkdirs(call otto.FunctionCall) otto.Value {
	var (
		path string
		err  error
	)
	path, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	err = os.MkdirAll(path, 0644)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, otto.Value{}, otto.Value{})
}

// Remove a file or folder recursively.
// :SIGJS: (path string) -> undefined
func goatsejs_rmtree(call otto.FunctionCall) otto.Value {
	var (
		path string
		err  error
	)
	path, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	err = os.RemoveAll(path)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, otto.Value{}, otto.Value{})
}
