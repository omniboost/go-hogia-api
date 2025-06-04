package hogia_api_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	hogia_api "github.com/omniboost/go-hogia-api"
)

func TestPostContactRequest(t *testing.T) {
	// req := client.NewPostContactRequest()
	// req.PathParams().OrgID = "3cb44744-8fb9-4750-89b0-b29700871843"
	// req.QueryParams().ApiVersion = hogia_api.Date{time.Date(2025, 03, 12, 0, 0, 0, 0, time.UTC)}
	// req.RequestBody().ContactTypes = []int{2}
	// req.RequestBody().CountryCode = "SE"
	// req.RequestBody().Country = "Sverige"
	// req.RequestBody().CompanyName = "Test Company 2"
	// req.RequestBody().Number = "5"
	// req.RequestBody().GivenName = ""
	// req.RequestBody().SurName = ""
	// // req.RequestBody().IdentificationNumber = "112233-1234"
	// req.RequestBody().Customer = hogia_api.CustomerPost{GlnNumber: "7300009009558", PaymentTerms: "30", IsRutRot: false, EInvoiceReceiverEmail: "", InvoiceReference: "Faktura ref", IsActive: true, SelectedCurrencyCode: "SE", IsReverseCharge: false}
	// req.RequestBody().VatNumber = "SE112233123401"
	// req.RequestBody().Language = "sv-SE"
	// // req.RequestBody().AddressType = 2
	// // req.RequestBody().Street = "Test"
	// // req.RequestBody().StreetExtraLine = ""
	// // req.RequestBody().City = "Staden"
	// // req.RequestBody().ZipCode = "12345"
	// contactResp, err := req.Do()
	// if err != nil {
	// 	t.Error(err)
	// }

	// b, _ := json.MarshalIndent(contactResp, "", "  ")
	// fmt.Println(string(b))

	// contactID := contactResp.ID

	addressReq := client.NewPostContactAddressRequest()
	addressReq.PathParams().OrgID = "3cb44744-8fb9-4750-89b0-b29700871843"
	addressReq.QueryParams().ApiVersion = hogia_api.Date{time.Date(2025, 03, 21, 0, 0, 0, 0, time.UTC)}
	// addressReq.PathParams().ContactID = contactID
	addressReq.PathParams().ContactID = "df52c06b-a79d-4212-a38b-b2a700cf113c"

	addressReq.RequestBody().AddressType = 2
	addressReq.RequestBody().Street = "Test"
	addressReq.RequestBody().StreetExtraLine = ""
	addressReq.RequestBody().City = "Staden"
	addressReq.RequestBody().ZipCode = "12345"

	resp, err := addressReq.Do()
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
