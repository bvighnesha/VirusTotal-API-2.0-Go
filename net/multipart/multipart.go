package multipart

import (
	"bytes"
	"io"
)

const (
	CONTENT_TYPE_FORM_DATA_WITH_BOUNDARY string = "multipart/form-data; boundary=" + BOUNDARY
	CONTENT_TYPE_BINARY                  string = "Content-Type: application/octet-stream"
	CONTENT_TYPE_TEXT                    string = "Content-Type: text/plain"
	CONTENT_TRANSFER_ENCODING            string = "Content-Transfer-Encoding: chunked"
	LINE_FEED                            string = "\r\n"
	CHARSET                              string = "UTF-8"
	BOUNDARY                             string = "e2a540ab4e6c5ed79c01157c255a2b5007e157d7"
	BOUNDARY_FIX                         string = "--"
	USER_AGENT                           string = "This Library Created by Bojja Vighneswar Rao, vighnesh.org"
)

type MultipartEntity struct {
	bytes.Buffer
}

func (me *MultipartEntity) AddHeader(name, value string) {

	me.WriteString(name)
	me.WriteString(":")
	me.WriteString(value)
	me.WriteString(LINE_FEED)
}

func (me *MultipartEntity) AddTextBody(typ, value string) {
	me.WriteString(BOUNDARY_FIX)
	me.WriteString(BOUNDARY)
	me.WriteString(LINE_FEED)
	me.WriteString("Content-Disposition: form-data; name=\"" + typ + "\"")
	me.WriteString(LINE_FEED)
	me.WriteString(CONTENT_TYPE_TEXT)
	me.WriteString("; charset=")
	me.WriteString(CHARSET)
	me.WriteString(LINE_FEED)
	me.WriteString(LINE_FEED)
	me.WriteString(value)
	me.WriteString(LINE_FEED)
}

func (me *MultipartEntity) AddBinaryBody(typ, filename string, body io.Reader) {

	me.WriteString(BOUNDARY_FIX)
	me.WriteString(BOUNDARY)
	me.WriteString(LINE_FEED)
	me.WriteString("Content-Disposition: form-data; name=\"" + typ + "\"; filename=\"" + filename + "\"")
	me.WriteString(LINE_FEED)
	me.WriteString(CONTENT_TYPE_BINARY)
	me.WriteString(LINE_FEED)
	me.WriteString(CONTENT_TRANSFER_ENCODING)
	me.WriteString(LINE_FEED)
	me.WriteString(LINE_FEED)

	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	me.Write(buf.Bytes())
}

func (me *MultipartEntity) Build() []byte {
	me.WriteString(LINE_FEED)
	me.WriteString(BOUNDARY_FIX)
	me.WriteString(BOUNDARY)
	me.WriteString(BOUNDARY_FIX)
	me.WriteString(LINE_FEED)
	return me.Bytes()
}
