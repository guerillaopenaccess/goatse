// Put JS and Go functions in goatse object
package main

import (
	"github.com/robertkrimen/otto"
	"log"
	"sort"
	"strings"
)

// Go functions here to attach to `goatse` object in JS environment.
func getGoatseGoFunctions() map[string]func(otto.FunctionCall) otto.Value {
	var prepared map[string]func(otto.FunctionCall) otto.Value
	prepared = map[string]func(otto.FunctionCall) otto.Value{}
	prepared["_log"] = goatsejs_log
	prepared["_sleep"] = goatsejs_sleep
	prepared["_randInt"] = goatsejs_randInt
	prepared["_exec"] = goatsejs_exec
	prepared["_getenv"] = goatsejs_getenv
	prepared["_argv"] = goatsejs_argv
	prepared["_where"] = goatsejs_where
	prepared["_exists"] = goatsejs_exists
	prepared["_rename"] = goatsejs_rename
	prepared["_mkdirs"] = goatsejs_mkdirs
	prepared["_rmtree"] = goatsejs_rmtree
	prepared["_require"] = goatsejs_requirejs
	prepared["_httpMethod"] = goatsejs_httpmethod
	prepared["_saveToFile"] = goatsejs_saveToFile
	prepared["_loadFromFile"] = goatsejs_loadFromFile
	prepared["_md5SumFile"] = goatsejs_md5SumFile
	prepared["_fileSize"] = goatsejs_fileSize
	prepared["_query_find"] = goatsejs_query_find
	prepared["_query_attr"] = goatsejs_query_attr
	prepared["_query_children"] = goatsejs_query_children
	prepared["_query_text"] = goatsejs_query_text
	prepared["_getDOIMeta"] = goatsejs_getDOIMeta
	prepared["_memberSample"] = goatsejs_memberSample
	prepared["_notOpenSample"] = goatsejs_notOpenSample
	prepared["_strings_entry"] = goatsejs_strings_entry
	prepared["_strings_urlescape"] = goatsejs_strings_urlescape
	prepared["_strings_urlunescape"] = goatsejs_strings_urlunescape
	return prepared
}

// Combined source in order "goatse_js_builtins"
func getGoatseJSFunctions() (script string) {
	var (
		scriptNames []string
		scripts     []string
		contents    []byte
		err         error
	)
	script = "// Sources JS\r\n"
	scriptNames = getAllSourceJS()
	scripts = make([]string, len(scriptNames))
	for i, scriptn := range scriptNames {
		contents, err = Asset("goatse_js_builtins/" + scriptn)
		if err != nil {
			panic("Failed to builtin JS: " + err.Error())
		}
		scripts[i] = "// " + scriptn + ":\r\n" + string(contents)
	}
	script = strings.Join(scripts, "\r\n\r\n")
	return script
}

// Iterates over JS source file names.
func getAllSourceJS() (allJSSource []string) {
	var (
		err error
	)
	allJSSource, err = AssetDir("goatse_js_builtins")
	if err != nil {
		panic(err.Error())
	}
	sort.Strings(allJSSource)
	return allJSSource
}

// Create JS `goatse` object. Add Go funcs. Run JS source to add JS built-ins.
func registerGoatseJS(ovm *otto.Otto) {
	var (
		goatse_gofuncs map[string]func(otto.FunctionCall) otto.Value
		goatse_obj     *otto.Object
		JSSource       string
		err            error
	)
	goatse_obj, err = ovm.Object("goatse = {}")
	logFatal(err, "Failed to make object: ")
	goatse_gofuncs = getGoatseGoFunctions()
	for name, function := range goatse_gofuncs {
		err = goatse_obj.Set(name, function)
		logFatal(err, "Failed to put "+name+" in goatse: ")
	}
	JSSource = getGoatseJSFunctions()
	_, err = ovm.Run(JSSource)
	logFatal(err, "Failed to put JS functions in object: ")
}

func logFatal(err error, msg string) {
	if err != nil {
		log.Fatal(msg + err.Error())
	}
}
