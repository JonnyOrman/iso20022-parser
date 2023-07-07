package iso20022

import "encoding/xml"

type Parser struct {
}

func NewParser() *Parser {
	this := new(Parser)

	return this
}

func (this Parser) Parse(iso20022RecordXml string) Iso20022Record {
	var iso20022Record Iso20022Record
	if err := xml.Unmarshal([]byte(iso20022RecordXml), &iso20022Record); err != nil {
		panic(err)
	}

	return iso20022Record
}
