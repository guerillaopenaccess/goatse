package main

import (
	"github.com/guerillaopenaccess/goatse/doiutils"
	"github.com/robertkrimen/otto"
)

// Metadata for DOI
// :SIGJS: (string)->[object]
func goatsejs_getDOIMeta(call otto.FunctionCall) otto.Value {
	var (
		doi     string
		meta    *doiutils.DoiMeta
		retmeta otto.Value
		err     error
	)
	doi, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	meta, err = doiutils.GetDOIMeta(doi)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	retmeta, err = call.Otto.ToValue(meta)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, retmeta, otto.Value{})
}

// List of dois for member. CR Top 20 only.
// :SIGJS: (int, string)->[object]
func goatsejs_memberSample(call otto.FunctionCall) otto.Value {
	var (
		s_num    int64
		m_name   string
		m_sample []doiutils.DoiMeta
		retv     otto.Value
		err      error
	)
	s_num, err = call.Argument(0).ToInteger()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	m_name, err = call.Argument(1).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	m_sample, err = doiutils.GetMemberSampleByName(int(s_num), m_name)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	retv, err = call.Otto.ToValue(m_sample)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, retv, otto.Value{})
}

// Sample list dois not open & not libgen
// :SIGJS: (int)->[object]
func goatsejs_notOpenSample(call otto.FunctionCall) otto.Value {
	var (
		s_num   int64
		notopen []doiutils.DoiMeta
		retv    otto.Value
		err     error
	)
	s_num, err = call.Argument(0).ToInteger()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	notopen, err = doiutils.GetNotOpenSample(int(s_num))
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	retv, err = call.Otto.ToValue(notopen)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, retv, otto.Value{})
}
