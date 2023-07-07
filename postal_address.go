package iso20022

type PostalAddress struct {
	StreetName     string `xml:"StrtNm"`
	BuildingNumber int    `xml:"BldgNb"`
	TownName       string `xml:"TwnNm"`
	Country        string `xml:"Ctry"`
}
