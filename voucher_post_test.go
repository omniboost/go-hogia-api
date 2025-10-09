package hogia_api_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	hogia_api "github.com/omniboost/go-hogia-api"
)

func TestPostVoucherRequest(t *testing.T) {
	req := client.NewPostVouchersRequest()
	req.PathParams().OrgID = "3cb44744-8fb9-4750-89b0-b29700871843"
	req.PathParams().FinanceyearID = "c5020bfd-b4ab-4e84-8084-b297009f6dcd"
	req.QueryParams().ApiVersion = hogia_api.Date{time.Date(2025, 06, 01, 0, 0, 0, 0, time.UTC)}
	req.RequestBody().Date = hogia_api.DateTime{time.Date(2025, 10, 8, 0, 0, 0, 0, time.UTC)}
	req.RequestBody().Number = 1
	req.RequestBody().TypeDto = "CustomerInvoice"
	req.RequestBody().Text = "Test"
	req.RequestBody().Serie = "A"
	req.RequestBody().CurrencyID = "3fa85f64-5717-4562-b3fc-2c963f66afa6"
	req.RequestBody().InvoiceID = "3fa85f64-5717-4562-b3fc-2c963f66afa6"
	req.RequestBody().InvoiceNumber = "12345"
	req.RequestBody().ParentID = "3fa85f64-5717-4562-b3fc-2c963f66afa6"
	req.RequestBody().AccrualTypeDto = "NotSet"
	req.RequestBody().SequenceID = "78428961-8bb0-4c8e-88ef-9a3f190e5a70"
	req.RequestBody().VoucherRows = []hogia_api.VoucherRows{
		{AccountNumber: 1510, NumberOf: 1, RowNumber: 1, Amount: 100, Text: "Test", Specification: "Test", ProjectNumber: "1", VoucherRowDimensions: []hogia_api.VoucherRowDimensions{{DimensionLevel: 1, DimensionNumber: "1000"}}},
	}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
