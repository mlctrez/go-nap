package generator

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/dave/jennifer/jen"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GenerateDir(path string) (err error) {
	return filepath.Walk(path, walk)
}

var lineRe = regexp.MustCompile(" line ([0-9]+):")

func walk(path string, info fs.FileInfo, walkErr error) (err error) {

	if walkErr != nil {
		return walkErr
	}
	if strings.HasSuffix(info.Name(), ".html") {
		err = Generate(path)
		if err != nil {
			subMatch := lineRe.FindAllStringSubmatch(err.Error(), 1)
			if len(subMatch) > 0 {
				path += ":" + subMatch[0][1]
			}
			err = fmt.Errorf("%s : %s", path, err)
		}
	}
	return err
}

const NapPkg = "github.com/mlctrez/go-nap/nap"

const CPrefix = "E"

func Generate(html string) (err error) {

	var open *os.File
	if open, err = os.Open(html); err != nil {
		return err
	}
	defer func() { _ = open.Close() }()

	var root *element
	if root, err = parse(open); err != nil {
		return err
	}

	var pfx string
	if pfx = root.DataNap("prefix"); pfx == "" {
		return errors.New("missing prefix on first element")
	}

	f := jen.NewFile(filepath.Base(filepath.Dir(html)))

	f.HeaderComment("This code is auto generated. DO NOT EDIT.")

	allDataNap := root.allDataNap()
	for _, dn := range allDataNap {
		f.Const().Id(CPrefix + uc(pfx) + uc(dn.DataNap())).Op("=").Lit(pfx + "/" + dn.DataNap())
	}

	rRouter := jen.Id("r").Qual(NapPkg, "Router")
	goFilePath := strings.TrimSuffix(html, ".html") + ".go"

	if root.DataNap("override") != "" {
		overrideFile := filepath.Join(filepath.Dir(goFilePath), root.DataNap("override"))
		if _, err = os.Stat(overrideFile); err != nil {
			return err
		}
	}

	f.Func().Id(uc(pfx)).Params(rRouter).
		BlockFunc(func(group *jen.Group) {
			for _, dn := range allDataNap {
				methodName := jen.Id(uc(pfx) + uc(dn.DataNap()))

				//dnPage := dn.DataNap("page")
				//if dnPage != "" {
				//	group.Id("r").Dot("ElmFunc").Params(
				//		jen.Id("page-"+dnPage), methodName,
				//	)
				//}
				//
				//dnBody := dn.DataNap("body")
				//if dnBody != "" {
				//	group.Id("r").Dot("ElmFunc").Params(
				//		jen.Id("body-"+dnBody), methodName,
				//	)
				//}

				group.Id("r").Dot("ElmFunc").Params(
					jen.Id(CPrefix+uc(pfx)+uc(dn.DataNap())), methodName,
				)
			}
			if root.DataNap("override") != "" {
				group.Id(uc(pfx) + "Override").Params(jen.Id("r"))
			}
		})

	for _, dn := range allDataNap {
		f.Line()
		f.Func().Id(uc(pfx)+uc(dn.DataNap())).
			Params(rRouter).Qual(NapPkg, "Elm").
			Block(jen.ReturnFunc(dn.declaration))
	}

	var outFile *os.File
	if outFile, err = os.Create(goFilePath); err != nil {
		return err
	}
	defer func() { _ = outFile.Close() }()

	if err = f.Render(outFile); err != nil {
		return err
	}

	return nil
}

func (d *element) declaration(group *jen.Group) {

	if d.name == "#text" {
		group.Qual(NapPkg, "Text").Parens(jen.Lit(d.Get("data")))
		return
	}

	var ret *jen.Statement

	if d.newLine {
		ret = group.Line().Id("r").Dot("E").Params(jen.Lit(d.name))
	} else {
		ret = group.Id("r").Dot("E").Params(jen.Lit(d.name))
	}

	filterAttributes := d.FilterAttributes("data-nap")
	newLine := len(filterAttributes) > 2
	for _, attr := range filterAttributes {
		if newLine {
			ret.Op(".").Line().Id("Set").Params(jen.Lit(attr.Name.Local), jen.Lit(attr.Value))
		} else {
			ret.Dot("Set").Params(jen.Lit(attr.Name.Local), jen.Lit(attr.Value))
		}
	}

	prefix := d.ParentDataNap("prefix")
	if len(d.children) > 0 {
		filteredChildren := d.FilterChildren("omit")
		if len(filteredChildren) > 0 {
			ret.Op(".").Line().Id("Append").ParamsFunc(func(group *jen.Group) {
				var newLine = len(filteredChildren) > 1
				for _, child := range filteredChildren {
					if child.DataNap() != "" {
						id := jen.Id(CPrefix + uc(prefix) + uc(child.DataNap()))
						if newLine {
							group.Line().Id("r").Dot("Elm").Params(id)
						} else {
							group.Id("r").Dot("Elm").Params(id)
						}
					} else {
						child.newLine = newLine
						child.declaration(group)
					}
				}
			})
		}
	}
}

func uc(in string) string {
	switch len(in) {
	case 0:
		return in
	case 1:
		return strings.ToUpper(in)
	default:
		return strings.ToUpper(in[0:1]) + in[1:]
	}
}

func parse(file io.Reader) (el *element, err error) {
	var parents []*element
	charBuffer := bytes.NewBufferString("")
	decoder := xml.NewDecoder(file)

	appendText := func() {
		data := strings.TrimSpace(charBuffer.String())
		charBuffer.Reset()
		if len(data) > 0 && len(parents) > 0 {
			parents[len(parents)-1].Text(data)
		}
	}

	var token xml.Token
	for ; err == nil; token, err = decoder.Token() {
		switch t := token.(type) {
		case xml.StartElement:
			appendText()
			ne := &element{name: t.Name.Local, attributes: t.Attr}
			parents = append(parents, ne)
			if len(parents) > 1 {
				parents[len(parents)-2].appendChild(ne)
			}
		case xml.EndElement:
			appendText()
			if len(parents) > 1 {
				parents = parents[:len(parents)-1]
			}
		case xml.CharData:
			charBuffer.Write(t)
		}
	}
	if len(parents) == 0 {
		err = errors.New("no parent elements")
	}
	if err != io.EOF && err != nil {
		return el, err
	}

	return parents[0], nil
}

type element struct {
	name       string
	attributes []xml.Attr
	parent     *element
	children   []*element
	newLine    bool
}

func (d *element) appendChild(el *element) {
	d.children = append(d.children, el)
	el.parent = d
}

func (d *element) Set(name, value string) {
	for _, attribute := range d.attributes {
		if attribute.Name.Local == name {
			attribute.Value = value
			return
		}
	}
	d.attributes = append(d.attributes, xml.Attr{
		Name:  xml.Name{Local: name},
		Value: value,
	})
}

func (d *element) Get(name string) string {
	for _, attribute := range d.attributes {
		if attribute.Name.Local == name {
			return attribute.Value
		}
	}
	return ""
}

func (d *element) Text(data string) {
	el := &element{name: "#text"}
	el.Set("data", data)
	d.appendChild(el)
}

func (d *element) allDataNap() (result []*element) {
	if d.DataNap() != "" {
		result = append(result, d)
	}
	for _, child := range d.children {
		result = append(result, child.allDataNap()...)
	}
	return result
}

func (d *element) DataNap(suffixes ...string) string {
	key := "data-nap"
	if len(suffixes) > 0 {
		key += "-" + strings.Join(suffixes, "-")
	}
	return d.Get(key)
}

func (d *element) ParentDataNap(suffixes ...string) string {
	for c := d; c != nil; c = c.parent {
		if nap := c.DataNap(suffixes...); nap != "" {
			return nap
		}
	}
	return ""
}

func (d *element) FilterAttributes(prefix string) []xml.Attr {
	var result []xml.Attr
	for _, attribute := range d.attributes {
		if !strings.HasPrefix(attribute.Name.Local, prefix) {
			result = append(result, attribute)
		}
	}
	return result
}

func (d *element) FilterChildren(suffix string) []*element {
	if len(d.children) == 0 {
		return nil
	}
	var result []*element
	for _, child := range d.children {
		if child.DataNap(suffix) == "" {
			result = append(result, child)
		}
	}
	return result
}
