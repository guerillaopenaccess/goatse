package main

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
	"golang.org/x/net/html"
	"strings"
)

// Turns a string into a goquery.Selection
func wrap_docfragment(fragment string, nobody bool) (sel *goquery.Selection, err error) {
	var (
		frag_reader *strings.Reader
		doc         *goquery.Document
	)
	frag_reader = strings.NewReader(fragment)
	doc, err = goquery.NewDocumentFromReader(frag_reader)
	if nobody {
		sel = doc.Find("body").Children().Eq(0)
		if len(sel.Nodes) == 0 {
			sel = doc.Find("head").Children().Eq(0)
		}
		if len(sel.Nodes) == 0 {
			panic("Failed to wrap HTML fragment without padding: " + fragment)
		}
	} else {
		sel = doc.Clone()
	}
	return sel, err
}

// Selection.Html returns html inside node, not node. This returns node.
func flatten_html_node(node *html.Node) string {
	var buf *bytes.Buffer
	buf = new(bytes.Buffer)
	html.Render(buf, node)
	return buf.String()
}

func flatten_selection_to_html(i int, sel *goquery.Selection) (retstr string) {
	return flatten_html_node(sel.Nodes[0])
}

func flatten_selection_to_text(i int, sel *goquery.Selection) (retstr string) {
	return sel.Text()
}

// Returns strings for each element by `query` in `frag`.
// :SIGJS: (string, string, bool) -> [string]
func goatsejs_query_find(call otto.FunctionCall) otto.Value {
	var (
		frag    string
		query   string
		nowrap  bool
		doc     *goquery.Selection
		sel     *goquery.Selection
		retstrs []string
		ottstrs otto.Value
		err     error
	)
	frag, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	query, err = call.Argument(1).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	nowrap, err = call.Argument(2).ToBoolean()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	doc, err = wrap_docfragment(frag, nowrap)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	sel = doc.Find(query)
	retstrs = sel.Map(flatten_selection_to_html)
	ottstrs, err = call.Otto.ToValue(retstrs)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, ottstrs, otto.Value{})
}

// String value for attribute.
// :SIGJS: (docfragment, attr string)->string
func goatsejs_query_attr(call otto.FunctionCall) otto.Value {
	var (
		frag        string
		attrname    string
		doc         *goquery.Selection
		attrval     string
		attrexists  bool
		emptystring otto.Value
		ottoattr    otto.Value
		err         error
	)
	frag, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	attrname, err = call.Argument(1).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	doc, err = wrap_docfragment(frag, true)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	attrval, attrexists = doc.Attr(attrname)
	if !attrexists {
		emptystring, err = call.Otto.ToValue("")
		if err != nil {
			return returnErrCapsule(call.Otto, err)
		}
		return returnCapsule(call.Otto, emptystring, otto.Value{})
	}
	attrval = strings.TrimSpace(attrval)
	ottoattr, err = call.Otto.ToValue(attrval)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, ottoattr, otto.Value{})
}

// List of strings of children.
// :SIGJS: (docfragment string) -> [string]
func goatsejs_query_children(call otto.FunctionCall) otto.Value {
	var (
		frag         string
		doc          *goquery.Selection
		children     *goquery.Selection
		flatchildren []string
		ottochildren otto.Value
		err          error
	)
	frag, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	doc, err = wrap_docfragment(frag, true)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	children = doc.Children()
	flatchildren = children.Map(flatten_selection_to_html)
	ottochildren, err = call.Otto.ToValue(flatchildren)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, ottochildren, otto.Value{})
}

// Text-only contents.
// :SIGJS: (docfragment string) -> string
func goatsejs_query_text(call otto.FunctionCall) otto.Value {
	var (
		frag    string
		doc     *goquery.Selection
		ottodoc otto.Value
		err     error
	)
	frag, err = call.Argument(0).ToString()
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	doc, err = wrap_docfragment(frag, true)
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	ottodoc, err = call.Otto.ToValue(doc.Text())
	if err != nil {
		return returnErrCapsule(call.Otto, err)
	}
	return returnCapsule(call.Otto, ottodoc, otto.Value{})
}
