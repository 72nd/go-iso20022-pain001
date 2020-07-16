package pain001

import (
	"strings"
)

// PaymentInformation contains most of the payment information. For more information please refer to the standard.
type PaymentInformation struct {
	PaymentInformationId         string                        `xml:"PmtInfId"`
	PaymentMethod                string                        `xml:"PmtMtd"`
	BatchBooking                 bool                          `xml:"BtchBookg"`
	RequiredExecutionDate        string                        `xml:"ReqdExctnDt"`
	Debitor                      Address                       `xml:"Dbtr"`
	DebitorIban                  string                        `xml:"DbtrAcct>Id>IBAN"`
	DebitorBic                   string                        `xml:"DbtrAgt>FinInstnId>BIC"`
	CreditorTransferInformation []CreditorTransferInformation `xml:"CdtTrfTxInf"`
}

// NewPaymentInformation takes an Order as an argument and returns a new PaymentInformation.
func NewPaymentInformation(order Order) PaymentInformation {
	info := PaymentInformation{
		PaymentInformationId:  "PMTINFID-01",
		PaymentMethod:         "TRF",
		BatchBooking:          true,
		RequiredExecutionDate: order.ExecuteOn,
		Debitor: Address{
			Name: order.Debitor.Name,
			PostalAddress: PostalAddress{
				StreetName:   order.Debitor.Street,
				StreetNumber: order.Debitor.StreetNr,
				PostalCode:   order.Debitor.Postcode,
				TownName:     order.Debitor.Place,
				Country:      order.Debitor.Country,
			},
		},
		DebitorIban: strings.Replace(order.Debitor.IBAN, " ", "", -1),
		DebitorBic:  order.Debitor.BIC,
	}

	info.CreditorTransferInformation = make([]CreditorTransferInformation, len(order.Transactions))
	for i := range order.Transactions {
		info.CreditorTransferInformation[i] = NewCreditorTransferInformation(order.Transactions[i], i+1)
	}
	return info
}
