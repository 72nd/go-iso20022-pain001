package pain001

import "testing"

func TestNewCreditorTransferInformation(t *testing.T) {
	expected := `<CreditorTransferInformation><PmtId><InstrId>INSTRID-01-0</InstrId><EndToEndId>ENDTOENDID-0</EndToEndId></PmtId><PmtTpInf><LclInstrm><Prtry>CH03</Prtry></LclInstrm></PmtTpInf><Amt><InstdAmt Ccy="CFN">17.17</InstdAmt></Amt><Cdtr><Nm>Joe Moon</Nm><PstlAdr><StrtNm>Muldon Street</StrtNm><BldgNb>17</BldgNb><PstCd>1717</PstCd><TwnNm>Steet Town</TwnNm><Ctry>FN</Ctry></PstlAdr></Cdtr><CdtrAcct><Id><IBAN>FN5604835012345678009</IBAN></Id></CdtrAcct><RmtInf><Ustrd>statuary</Ustrd></RmtInf></CreditorTransferInformation>`
	data := NewCreditorTransferInformation(Transaction{
		Name:      "Joe Moon",
		Street:    "Muldon Street",
		StreetNr:  17,
		Postcode:  1717,
		Place:     "Steet Town",
		Country:   "FN",
		IBAN:      "FN56 0483 5012 3456 7800 9",
		Reference: "statuary",
		Amount:    "17.17",
		Currency:  "CFN",
	}, 0)
	compareXmlResult(t, data, expected)
}

func TestAmount(t *testing.T) {
	expected := `<Amount Ccy="CFN">23.42</Amount>`
	data := Amount{
		Currency: "CFN",
		Value:    "23.42"}
	compareXmlResult(t, data, expected)
}
