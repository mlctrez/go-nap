package enc

import (
	"bytes"
	"encoding/xml"
	"io"
)

type Encodable interface {
	Encode(encoder *xml.Encoder) error
}

type BytesEncoder interface {
	Indent(depth string) BytesEncoder
	Encode(el Encodable) error
	EncodePage(el Encodable) error
	Write(w io.Writer) error
	Content() string
}

var _ BytesEncoder = (*bytesEncoder)(nil)

type bytesEncoder struct {
	buf *bytes.Buffer
	enc *xml.Encoder
}

func Dump(el Encodable) string {
	enc := New()
	enc.Indent("  ")
	err := enc.Encode(el)
	if err != nil {
		return err.Error()
	}
	return enc.Content()
}

func New() BytesEncoder {
	buf := &bytes.Buffer{}
	return &bytesEncoder{buf: buf, enc: xml.NewEncoder(buf)}
}

func (b *bytesEncoder) Indent(indent string) BytesEncoder {
	b.enc.Indent("", indent)
	return b
}

func (b *bytesEncoder) Encode(en Encodable) (err error) {
	if err = en.Encode(b.enc); err != nil {
		return err
	}
	if err = b.enc.Flush(); err != nil {
		return err
	}
	err = b.enc.Close()
	return err
}

func (b *bytesEncoder) EncodePage(en Encodable) (err error) {
	if err = b.enc.EncodeToken(xml.Directive("DOCTYPE html")); err != nil {
		return err
	}
	if err = b.enc.EncodeToken(xml.CharData("\n")); err != nil {
		return err
	}
	return b.Encode(en)
}

func (b *bytesEncoder) Write(w io.Writer) (err error) {
	_, err = w.Write(b.buf.Bytes())
	return err
}

func (b *bytesEncoder) Content() string {
	return b.buf.String()
}
