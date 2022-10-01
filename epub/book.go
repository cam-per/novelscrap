package epub

import (
	"fmt"
	"io"

	"github.com/beevik/etree"
)

type container struct {
	Version  string
	Encoding string

	Rootfiles []struct {
		Path      string
		MediaType string
	}
}

func (c *container) write(w io.Writer) (int64, error) {
	doc := etree.NewDocument()

	doc.CreateProcInst("xml", fmt.Sprintf(`version="%s" encoding="%s"`, c.Version, c.Encoding))

	cont := doc.CreateElement("container")
	cont.CreateAttr("version", c.Version)
	cont.CreateAttr("xmlns", "urn:oasis:names:tc:opendocument:xmlns:container")

	rootfiles := cont.CreateElement("rootfiles")
	for _, rf := range c.Rootfiles {
		rootfile := rootfiles.CreateElement("rootfile")
		rootfile.CreateAttr("full-path", rf.Path)
		rootfile.CreateAttr("media-type", rf.MediaType)
	}
	return doc.WriteTo(w)
}

type Book struct {
	Container container
}
