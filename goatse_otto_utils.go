package main

import (
	"github.com/robertkrimen/otto"
)

// Success or error in JS object
func returnCapsule(vm *otto.Otto, retval, err otto.Value) otto.Value {
	var combined *otto.Object
	//	combined, _ = vm.Object("_mostrecentreturncapsule = {}")
	combined, _ = vm.Object("({})")
	combined.Set("value", retval)
	combined.Set("error", err)
	return combined.Value()
}

// Error in JS object.
func returnErrCapsule(ovm *otto.Otto, err error) otto.Value {
	var (
		errval otto.Value
	)
	errval, _ = ovm.ToValue(err.Error())
	return returnCapsule(ovm, otto.Value{}, errval)
}
