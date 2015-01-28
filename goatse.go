// REP Aaron
package main

import (
	"flag"
	"fmt"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
	"log"
)

var (
	modeflag   string
	scriptflag string
	vm         *otto.Otto
)

func init() {
	flag.StringVar(&modeflag, "mode", "auto", "Mode of operation; [script, doilist, source]")
	flag.StringVar(&scriptflag, "f", "NOT GIVEN", "script or list of DOI")
	flag.Parse()
	vm = otto.New()
	registerGoatseJS(vm)
}

func main() {
	var (
		err error
	)
	switch modeflag {
	case "auto":
		{
			log.Print("Auto mode: freeing articles")
			err = autoMode()
			if err != nil {
				log.Fatal("Error in Auto mode: " + err.Error())
			}
		}
	case "script":
		{
			log.Print("Scripted mode")
			err = scriptMode()
			if err != nil {
				log.Fatal("Error in script mode: " + err.Error())
			}
		}
	case "doilist":
		{
			err = listMode()
			if err != nil {
				log.Fatal("Error in list mode: " + err.Error())
			}
		}
	case "source":
		{
			err = sourceMode()
			if err != nil {
				log.Fatal("Error in source mode: " + err.Error())
			}
		}
	default:
		{
			log.Fatal("Bad mode. Modes are 'script', 'doilist', 'source'")
		}
	}
}

func autoMode() (err error) {
	_, err = vm.Run(`goatse.default_mode();`)
	if err != nil {
		return err
	}
	return nil
}

// Requires list "-f"
func listMode() (err error) {
	var goatse_object *otto.Object
	goatse_object, err = vm.Object(`goatse`)
	if err != nil {
		return err
	}
	_, err = goatse_object.Call("doilist_mode", scriptflag)
	if err != nil {
		return err
	}
	return nil
}

// Requires script "-f".
func scriptMode() (err error) {
	var (
		script *otto.Script
	)
	script, err = vm.Compile(scriptflag, nil)
	if err != nil {
		if err.Error() == "open "+scriptflag+": no such file or directory" {
			return fmt.Errorf("File '%s' not found, cannot run specified script.", scriptflag)
		} else {
			return err
		}
	}
	_, err = vm.Run(script)
	if err != nil {
		return err
	}
	return nil
}

// Put source tree to "sourcetree".
func sourceMode() (err error) {
	log.Print("Source mode: dumping source tree to directory 'sourcetree'")
	err = RestoreAssets("goatse_source", "sourcetree")
	return err
}
