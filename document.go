package pain001

import (
	"encoding/xml"
)

// DateTimeFormat represents the date and time format which is used in the standard.
const DateTimeFormat = "2006-01-02T15:04:05"

// Document is the root structure of the payment statement.
type Document struct {
	XMLName                          xml.Name `xml:"Document"`
	Xmlns                            string   `xml:"xmlns,attr"`
	XmlnsXid                         string   `xml:"xmlns:xsi,attr"`
	XsiSchemaLocation                string   `xml:"xsi:schemaLocation,attr"`
	CustomerCreditTransferInitiation CustomerCreditTransferInitiation
}

// NewDocument returns a new document.
func NewDocument(order Order) (Document, error) {
	initiation, err := NewCustomerCreditTransferInitiation(order)
	if err != nil {
		return Document{}, err
	}
	document := Document{
		Xmlns:                            "http://www.six-interbank-clearing.com/de/pain.001.001.03.ch.02.xsd",
		XmlnsXid:                         "http://www.w3.org/2001/XMLSchema-instance",
		XsiSchemaLocation:                "http://www.six-interbank-clearing.com/de/pain.001.001.03.ch.02.xsd  pain.001.001.03.ch.02.xsd",
		CustomerCreditTransferInitiation: initiation,
	}
	return document, nil
}

// toXml returns the document as XML.
func (d Document) toXml(indent bool) ([]byte, error) {
	var out []byte
	var err error
	if indent {
		out, err = xml.MarshalIndent(d, "  ", "    ")
	} else {
		out, err = xml.Marshal(d)
	}

	if err != nil {
		return []byte{}, err
	}
	return out, nil
}
