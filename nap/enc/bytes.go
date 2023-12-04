package enc

import (
	"bytes"
	"encoding/xml"
	"io"
)

type Encode interface {
	Encode(encoder *xml.Encoder) error
}

type BytesEncoder interface {
	Indent(depth string) BytesEncoder
	Encode(el Encode) error
	EncodePage(el Encode) error
	Write(w io.Writer) error
	Content() string
}

var _ BytesEncoder = (*bytesEncoder)(nil)

type bytesEncoder struct {
	buf *bytes.Buffer
	enc *xml.Encoder
}

func New() BytesEncoder {
	buf := &bytes.Buffer{}
	return &bytesEncoder{buf: buf, enc: xml.NewEncoder(buf)}
}

func (b *bytesEncoder) Indent(indent string) BytesEncoder {
	b.enc.Indent("", indent)
	return b
}

func (b *bytesEncoder) Encode(el Encode) (err error) {
	if err = el.Encode(b.enc); err != nil {
		return err
	}
	if err = b.enc.Flush(); err != nil {
		return err
	}
	err = b.enc.Close()
	return err
}

func (b *bytesEncoder) EncodePage(el Encode) (err error) {
	if err = b.enc.EncodeToken(xml.Directive("DOCTYPE html")); err != nil {
		return err
	}
	return b.Encode(el)
}

func (b *bytesEncoder) Write(w io.Writer) (err error) {
	_, err = w.Write(b.buf.Bytes())
	return err
}

func (b *bytesEncoder) Content() string {
	return b.buf.String()
}
