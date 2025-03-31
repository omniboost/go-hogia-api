package hogia_api_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	hogia_api "github.com/omniboost/go-hogia-api"
)

func TestPostVoucherDraftsRequest(t *testing.T) {
	req := client.NewPostVoucherDraftsRequest()
	req.PathParams().OrgID = "b135d637-1688-4c55-9061-b0a40094461b"
	req.QueryParams().ApiVersion = hogia_api.Date{time.Date(2025, 02, 13, 0, 0, 0, 0, time.UTC)}
	req.RequestBody().VoucherDraft.VoucherDate = hogia_api.DateTime{time.Date(2025, 02, 13, 0, 0, 0, 0, time.UTC)}
	req.RequestBody().VoucherDraft.VoucherType = "CustomerInvoice"
	req.RequestBody().VoucherDraft.Text = "Test"
	req.RequestBody().VoucherDraft.CurrencyCode = "SEK"
	req.RequestBody().VoucherDraft.Series = "A"
	// req.RequestBody().SequenceID = "6081157c-410e-42f0-8e72-67e970e4d1a5"
	req.RequestBody().VoucherDraftRows = []hogia_api.VoucherDraftRows{{Amount: 100, NumberOf: 0, ProjectNumber: "1", Specification: "Test", Text: "Test", AccountNumber: 1510, Dimensions: []hogia_api.DimensionsVoucher{{DimensionNumber: "1000", DimensionLevel: 1}}}}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
