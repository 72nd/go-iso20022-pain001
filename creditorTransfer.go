package pain001

import (
	"fmt"
	"strings"
)

// CreditorTransferInformation represents one payment.
type CreditorTransferInformation struct {
	InstructionId string  `xml:"PmtId>InstrId"`
	EndToEndId    string  `xml:"PmtId>EndToEndId"`
	PaymentType   string  `xml:"PmtTpInf>LclInstrm>Prtry"`
	Amount        Amount  `xml:"Amt>InstdAmt"`
	Address       Address `xml:"Cdtr"`
	Iban          string  `xml:"CdtrAcct>Id>IBAN"`
	Reference     string  `xml:"RmtInf>Ustrd"`
}

// NewCreditorTransferInformation returns a new CreditorTransferInformation. Needs to know, which number of payment it is.
func NewCreditorTransferInformation(transaction Transaction, paymentNumber int) CreditorTransferInformation {
	return CreditorTransferInformation{
		InstructionId: fmt.Sprintf("INSTRID-01-%d", paymentNumber),
		EndToEndId:    fmt.Sprintf("ENDTOENDID-%d", paymentNumber),
		PaymentType:   "CH03",
		Amount: Amount{
			Currency: transaction.Currency,
			Value:    transaction.Amount,
		},
		Address: Address{
			Name: transaction.Name,
			PostalAddress: PostalAddress{
				StreetName:   transaction.Street,
				StreetNumber: transaction.StreetNr,
				PostalCode:   transaction.Postcode,
				TownName:     transaction.Place,
				Country:      transaction.Country,
			},
		},
		Iban:      strings.Replace(transaction.IBAN, " ", "", -1),
		Reference: transaction.Reference,
	}
}

type Amount struct {
	Currency string `xml:"Ccy,attr"`
	Value    string `xml:",chardata"`
}
