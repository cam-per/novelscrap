package epub

var DefaultContainer = container{
	Version:  "1.0",
	Encoding: "UTF-8",
	Rootfiles: []struct {
		Path      string
		MediaType string
	}{
		{
			Path:      "OPS/book.opf",
			MediaType: "application/oebps-package+xml",
		},
	},
}
