package pain001

import "testing"

func TestNewDocument(t *testing.T) {
	doc, err := NewDocument(Order{
		ExecuteOn: "2020-01-02",
		Debitor: Debitor{
			Name:     "George Goodman",
			Street:   "Fnord Street",
			StreetNr: 23,
			Postcode: 2323,
			Place:    "Fnord Town",
			Country:  "FN",
			IBAN:     "FN56 1483 6012 4456 8800 0",
			BIC:      "EXP1234",
		},
		Transactions: []Transaction{{
			Name:      "Joe Moon",
			Street:    "Muldon Street",
			StreetNr:  17,
			Postcode:  1717,
			Place:     "Steet Town",
			Country:   "FN",
			IBAN:      "FN56 0483 5012 3456 7800 9",
			Reference: "statuary",
			Amount:    "17.17",
			Currency:  "CFN"}}})
	if err != nil {
		t.Error(err)
	}
	expected := "http://www.six-interbank-clearing.com/de/pain.001.001.03.ch.02.xsd"
	if doc.Xmlns != expected {
		t.Errorf("does not match, expected: \"%s\", actual: \"%s\"", expected, doc.Xmlns)
	}
	expected = "http://www.w3.org/2001/XMLSchema-instance"
	if doc.XmlnsXid != expected {
		t.Errorf("does not match, expected: \"%s\", actual: \"%s\"", expected, doc.XmlnsXid)
	}
	expected = "http://www.six-interbank-clearing.com/de/pain.001.001.03.ch.02.xsd  pain.001.001.03.ch.02.xsd"
	if doc.XsiSchemaLocation != expected {
		t.Errorf("does not match, expected: \"%s\", actual: \"%s\"", expected, doc.XsiSchemaLocation)
	}
	doc.CustomerCreditTransferInitiation.GroupHeader.MessageIdentification = "016dfe3b-6c81-428c"
	doc.CustomerCreditTransferInitiation.GroupHeader.CreationDateTime = "2020-07-16T19:59:25"
	expectedXml := `<Document xmlns="http://www.six-interbank-clearing.com/de/pain.001.001.03.ch.02.xsd" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.six-interbank-clearing.com/de/pain.001.001.03.ch.02.xsd  pain.001.001.03.ch.02.xsd"><CstmrCdtTrfInitn><GrpHdr><MsgId>016dfe3b-6c81-428c</MsgId><CreDtTm>2020-07-16T19:59:25</CreDtTm><NbOfTxs>1</NbOfTxs><CtrlSum>17.17</CtrlSum><InitgPty><Nm>George Goodman</Nm></InitgPty></GrpHdr><PmtInf><PmtInfId>PMTINFID-01</PmtInfId><PmtMtd>TRF</PmtMtd><BtchBookg>true</BtchBookg><ReqdExctnDt>2020-01-02</ReqdExctnDt><Dbtr><Nm>George Goodman</Nm><PstlAdr><StrtNm>Fnord Street</StrtNm><BldgNb>23</BldgNb><PstCd>2323</PstCd><TwnNm>Fnord Town</TwnNm><Ctry>FN</Ctry></PstlAdr></Dbtr><DbtrAcct><Id><IBAN>FN5614836012445688000</IBAN></Id></DbtrAcct><DbtrAgt><FinInstnId><BIC>EXP1234</BIC></FinInstnId></DbtrAgt><CdtTrfTxInf><PmtId><InstrId>INSTRID-01-1</InstrId><EndToEndId>ENDTOENDID-1</EndToEndId></PmtId><PmtTpInf><LclInstrm><Prtry>CH03</Prtry></LclInstrm></PmtTpInf><Amt><InstdAmt Ccy="CFN">17.17</InstdAmt></Amt><Cdtr><Nm>Joe Moon</Nm><PstlAdr><StrtNm>Muldon Street</StrtNm><BldgNb>17</BldgNb><PstCd>1717</PstCd><TwnNm>Steet Town</TwnNm><Ctry>FN</Ctry></PstlAdr></Cdtr><CdtrAcct><Id><IBAN>FN5604835012345678009</IBAN></Id></CdtrAcct><RmtInf><Ustrd>statuary</Ustrd></RmtInf></CdtTrfTxInf></PmtInf></CstmrCdtTrfInitn></Document>`
	compareXmlResult(t, doc, expectedXml)
}

func TestTransactionSum(t *testing.T) {
	order1 := Order{
		Transactions: []Transaction{
			{Amount: "10.10"},
			{Amount: "20.20"}},
	}
	order2 := Order{
		Transactions: []Transaction{
			{Amount: "10.10"}},
	}
	sum1, err := order1.transactionSum()
	if err != nil {
		t.Error(err)
	}
	sum2, err := order2.transactionSum()
	if err != nil {
		t.Error(err)
	}
	if sum1 != "30.30" {
		t.Errorf("sum of %s and %s was %s not 30.30", order1.Transactions[0].Amount, order1.Transactions[1].Amount, sum1)
	}
	if sum2 != "10.10" {
		t.Errorf("sum of %s (and nothing else) was %s not 10.10", order2.Transactions[0].Amount, sum2)
	}
}
