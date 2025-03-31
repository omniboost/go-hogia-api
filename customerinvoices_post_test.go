package hogia_api_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPostCustomerInvoicesRequest(t *testing.T) {
	req := client.NewPostCustomerInvoicesRequest()
	req.PathParams().OrgID = "b135d637-1688-4c55-9061-b0a40094461b"
	// req.QueryParams().ApiVersion = hogia_api.Date{time.Date(2025, 02, 13, 0, 0, 0, 0, time.UTC)}
	// req.RequestBody().Dimensions = []hogia_api.Dimensions{{DimensionNumber: "1", DimensionLevel: 1}}
	// req.RequestBody().SequenceID = "6081157c-410e-42f0-8e72-67e970e4d1a5"
	// req.RequestBody().InvoiceHeader.InvoiceDate = hogia_api.DateTime{time.Date(2025, 02, 13, 0, 0, 0, 0, time.UTC)}
	// req.RequestBody().InvoiceHeader.InvoiceExpirationDate = hogia_api.DateTime{time.Date(2025, 02, 18, 0, 0, 0, 0, time.UTC)}
	// req.RequestBody().InvoiceHeader.ShippingAddress.Street = "Geraltgatan 12"
	// req.RequestBody().InvoiceHeader.ShippingAddress.StreetExtraLine = "2:B"
	// req.RequestBody().InvoiceHeader.ShippingAddress.ZipCode = "44453"
	// req.RequestBody().InvoiceHeader.ShippingAddress.City = "Stenungsund"
	// req.RequestBody().InvoiceHeader.ShippingAddress.Country = "Sverige"
	// req.RequestBody().InvoiceHeader.BillingAddress = hogia_api.BillingAddress{Street: "Geraltgatan 12", StreetExtraLine: "2:B", ZipCode: "44453", City: "Stenungsund", Country: "Sverige"}
	// req.RequestBody().InvoiceHeader.Customer = hogia_api.CustomerForInvoice{Name: "Kund AB", Number: "1234", IdentificationNumber: "112233-1234", VatNumber: "SE112233123401", CountryCode: "SE"}
	// req.RequestBody().InvoiceHeader.Comment = "This is a test comment"
	// req.RequestBody().InvoiceHeader.OurReference = "Pelle"
	// req.RequestBody().InvoiceHeader.YourReference = "Kalle"
	// req.RequestBody().InvoiceHeader.DistributionMethod = 2
	// req.RequestBody().InvoiceHeader.DistributionEmail = "example@example.se"
	// req.RequestBody().InvoiceHeader.Language = "sv-SE"
	// req.RequestBody().InvoiceHeader.CustomerOrderNumber = "1234"
	// req.RequestBody().InvoiceHeader.CustomerProjectNumber = "10"
	// req.RequestBody().InvoiceHeader.SupplierOrderNumber = "1234"
	// req.RequestBody().InvoiceHeader.InvoiceReference = "1234"
	// req.RequestBody().CurrencyInfo = hogia_api.CurrencyInfo{CurrencyProfitsAccountNumber: 3960, CurrencyLossesAccountNumber: 7960, ExchangeRate: 1, RoundingFactor: 0, Currency: hogia_api.Currency{Decimals: 2, Code: "SEK"}, AccountingCurrency: hogia_api.AccountingCurrency{Decimals: 2, Code: "SEK"}}
	// req.RequestBody().InvoiceLines = []hogia_api.InvoiceLines{{UnitPrice: 100, UnitType: "st", Quantity: 1, Vat: hogia_api.Vat{Percent: 0.25, AccountNumber: 1920}, SalesAccountNumber: 3011, ArticleType: "service", RutRotWorkTypeName: "no-category", StockAccountNumber: 1400, CostAccountNumber: 3590, CostAmount: 77, Dimensions: []hogia_api.Dimensions{{DimensionNumber: "1", DimensionLevel: 0, AccountType: 0}}, ProjectNumber: "10", DiscountPercent: 0, Name: "Service", Type: "InvoiceLine"}}
	// // req.RequestBody().RutRot.RutRotApplicants = []hogia_api.RutRotApplicants{{SequenceID: "c540e767-2887-4241-b0c1-f0d17736bfb5", IdentificationNumber: "690409-1839", Name: "Tord Österberg", Amount: 562, ApplicationNumber: 0, UtcExportedDate: hogia_api.DateTime{time.Date(2025, 02, 13, 0, 0, 0, 0, time.UTC)}}}
	// // req.RequestBody().RutRot.RutRotWorkTypeLines = []hogia_api.RutRotWorkTypeLines{{MaterialCost: 2000, WorkHours: 1, Name: "Service", CustomerInvoiceID: "00000000-0000-0000-0000-000000000000"}}
	// // req.RequestBody().RutRot.IsRutRot = true
	// // req.RequestBody().RutRot.DebtAccountNumber = 1513
	// // req.RequestBody().RutRot.Type = 1
	// // req.RequestBody().RutRot.ReductionAmount = 562
	// // req.RequestBody().RutRot.PropertyDescription = "Stenungsund Höga 80:1"
	// // req.RequestBody().RutRot.BrfOrganizationNumber = "546546-5465"
	// // req.RequestBody().RutRot.ApartmentNumber = "45"
	// // req.RequestBody().RutRot.InformationText = "Test 1"
	// // req.RequestBody().RutRot.Address = hogia_api.Address{Street: "Geraltgatan 12", StreetExtraLine: "2:B", ZipCode: "44453", City: "Stenungsund", Country: "Sverige"}
	// // req.RequestBody().RutRot.OtherCosts = 0
	// // req.RequestBody().RutRot.TotalWorkCost = 2375
	// req.RequestBody().PaymentChannels = []hogia_api.PaymentChannels{{Type: "BankGiro", Value: "123-4567", IsPrimary: true}}
	// req.RequestBody().Organization.Name = "Pelles mekaniska AB"
	// req.RequestBody().Organization.OrganizationNumber = "123412-1234"
	// req.RequestBody().Organization.VatRegNumber = "SE123412123"
	// req.RequestBody().Organization.Address = hogia_api.Address{Street: "Geraltgatan 12", StreetExtraLine: "2:B", ZipCode: "44453", City: "Stenungsund", Country: "Sverige"}
	// req.RequestBody().Organization.ContactChannels = []string{"0303-12345", "contact@organization.com"}
	// req.RequestBody().Organization.TypedContactChannels = []hogia_api.TypedContactChannels{{Value: "contact@organization.com", ContactChannelType: 1}}
	// req.RequestBody().Organization.IsFTaxApproved = true
	// req.RequestBody().Organization.SubjectToVAT = true
	// req.RequestBody().Organization.Headquarters = "Stenungsund"
	// req.RequestBody().IsReverseCharge = false
	// req.RequestBody().CustomerDebtAccountNumber = 1510
	// req.RequestBody().ProjectNumber = "10"
	// req.RequestBody().AccountingMethod = 1
	// req.RequestBody().RoundingAccountNumber = 3740
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
