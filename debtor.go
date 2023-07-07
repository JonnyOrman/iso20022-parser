package iso20022

type Debtor struct {
	Name          string        `xml:"Nm"`
	PostalAddress PostalAddress `xml:"PstlAdr"`
}
