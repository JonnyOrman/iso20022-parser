package iso20022

type Iso20022Record struct {
	InternationalBankSettlementAmount InternationalBankSettlementAmount `xml:"IntrBkSttlmAmt"`
	Debtor                            Debtor                            `xml:"Dbtr"`
	DebtorAccount                     Account                           `xml:"DbtrAcct"`
	DebtorAgent                       Agent                             `xml:"DbtrAgt"`
}
