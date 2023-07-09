package iso20022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	iso20022RecordXml := `
	<CdtTrfTxInf>
		<IntrBkSttlmAmt Ccy="USD">12500</IntrBkSttlmAmt>
		<IntrBkSttlmDt>2022-04-06</IntrBkSttlmDt>
		<Dbtr>
			<Nm>ACME NV.</Nm>
			<PstlAdr>
				<StrtNm>Amstel</StrtNm>
				<BldgNb>344</BldgNb>
				<TwnNm>Amsterdam</TwnNm>
				<Ctry>NL</Ctry>
			</PstlAdr>
		</Dbtr>
		<DbtrAcct>
			<Id>
				<Othr>
					<Id>8754219990</Id>
				</Othr>
			</Id>
		</DbtrAcct>
		<DbtrAgt>
			<FinInstnId>
				<BIC>EXABNL2U</BIC>
			</FinInstnId>
		</DbtrAgt>
	</CdtTrfTxInf>`

	parser := NewParser()

	expected := Iso20022Record{
		InternationalBankSettlementAmount: InternationalBankSettlementAmount{
			Ccy:    "USD",
			Amount: 12500,
		},
		Debtor: Debtor{
			Name: "ACME NV.",
			PostalAddress: PostalAddress{
				StreetName:     "Amstel",
				BuildingNumber: 344,
				TownName:       "Amsterdam",
				Country:        "NL",
			},
		},
		DebtorAccount: Account{
			Id: 8754219990,
		},
		DebtorAgent: Agent{
			FinancialInstitutionId: FinancialInstitutionId{
				BIC: "EXABNL2U",
			},
		},
	}

	result := parser.Parse(iso20022RecordXml)

	assert.Equal(t, expected, result)
}
