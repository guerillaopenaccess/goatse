package httpplus

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path"
	"strings"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func newfileUploadRequest(uri string, params map[string]string, paramName, filePath string) (req *http.Request, err error) {
	var (
		file          *os.File
		first512      []byte
		contentType   string
		nameField     string
		filenameField string
		contentDisp   string
		body          *bytes.Buffer
		writer        *multipart.Writer
		fheader       textproto.MIMEHeader
		part          io.Writer
	)
	file, err = os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// For mimetype
	first512 = make([]byte, 512)
	file.Read(first512)
	file.Seek(0, 0)

	contentType = http.DetectContentType(first512)
	nameField = escapeQuotes(paramName)
	filenameField = escapeQuotes(path.Base(filePath))
	contentDisp = fmt.Sprintf(`form-data; name="%s"; filename="%s"`, nameField, filenameField)

	body = &bytes.Buffer{}
	writer = multipart.NewWriter(body)
	fheader = make(textproto.MIMEHeader)
	fheader.Set("Content-Disposition", contentDisp)
	fheader.Set("Content-Type", contentType)

	part, err = writer.CreatePart(fheader)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}

	// Not set
	req.Header.Set("Content-Type", "multipart/form-data; boundary="+writer.Boundary())
	return req, nil
}

func makeMultiparamsPOST(uri string, params map[string]string) (req *http.Request, err error) {
	var (
		body   *bytes.Buffer
		writer *multipart.Writer
	)
	body = &bytes.Buffer{}
	writer = multipart.NewWriter(body)

	for key, val := range params {
		writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "multipart/form-data; boundary="+writer.Boundary())

	return req, nil
}
