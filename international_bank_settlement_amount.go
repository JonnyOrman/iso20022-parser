package iso20022

type InternationalBankSettlementAmount struct {
	Ccy    string `xml:"Ccy,attr"`
	Amount int    `xml:",chardata"`
}
