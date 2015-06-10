package lye

import "encoding/xml"

const (
	// ContentType is the SOAP XML MIME type. Because acronyms!
	ContentType = "application/soap+xml"

	// Namespace is the SOAP XML namespace.
	Namespace = "http://www.w3.org/2003/05/soap-envelope"
)

// Request encapsualates a SOAP 1.2 request, used with encoding/xml
type Request struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	Body    struct {
		XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Body"`
		Cmds    []interface{}
	}
}

// NewRequest returns an initialized SOAP request.
func NewRequest(cmds ...interface{}) *Request {
	r := &Request{}
	r.Body.Cmds = append(r.Body.Cmds, cmds...)
	return r
}

// Marshal marshals the SOAP request using Go’s XML package.
func (r *Request) Marshal() ([]byte, error) {
	return xml.Marshal(r)
}

// MarshalIndent marshals the SOAP request with indentation.
func (r *Request) MarshalIndent() ([]byte, error) {
	return xml.MarshalIndent(r, "  ", "  ")
}
