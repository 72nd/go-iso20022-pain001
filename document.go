package pain001

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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
func NewDocument(data convert.Data) Document {
	document := Document{
		Xmlns:                            "http://www.six-interbank-clearing.com/de/pain.001.001.03.ch.02.xsd",
		XmlnsXid:                         "http://www.w3.org/2001/XMLSchema-instance",
		XsiSchemaLocation:                "http://www.six-interbank-clearing.com/de/pain.001.001.03.ch.02.xsd  pain.001.001.03.ch.02.xsd",
		CustomerCreditTransferInitiation: NewCustomerCreditTransferInitiation(data),
	}
	return document
}

