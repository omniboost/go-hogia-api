package hogia_api_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	hogia_api "github.com/omniboost/go-hogia-api"
)

func TestPostCustomerInvoicesV2Request(t *testing.T) {
	req := client.NewPostCustomerInvoicesV2Request()
	req.PathParams().OrgID = "b135d637-1688-4c55-9061-b0a40094461b"
	req.QueryParams().ApiVersion = hogia_api.Date{time.Date(2025, 02, 13, 0, 0, 0, 0, time.UTC)}
	req.RequestBody().SequenceID = "6081157c-410e-42f0-8e72-67e970e4d1a5"
	req.RequestBody().InvoiceLines = []hogia_api.InvoiceLines{
		// {TextInvoiceLine: hogia_api.TextInvoiceLine{Description: "Test"},
		// 	ItemInvoiceLine: hogia_api.ItemInvoiceLine{
		// 		Quantity:        5,
		// 		ItemNumber:      "1",
		// 		DiscountPercent: 10,
		// 		Dimensions:      []hogia_api.Dimensions{{Type: "One", Number: "1", Name: "RE1"}},
		// 		ProjectNumber:   "100",
		// 	}},
	}
	req.RequestBody().CustomerInvoice = hogia_api.CustomerInvoice{
		IsReverseCharge:       false,
		CustomerNumber:        "12",
		ProjectNumber:         "10",
		Dimensions:            []hogia_api.Dimensions{{Type: "One", Number: "1", Name: "RE1"}},
		InvoiceDate:           hogia_api.Date{time.Date(2025, 02, 13, 0, 0, 0, 0, time.UTC)},
		InvoiceExpirationDate: hogia_api.Date{time.Date(2025, 02, 18, 0, 0, 0, 0, time.UTC)},
		InvoiceNumber:         101,
		Ocr:                   "10157",
		Comment:               "Test",
		OurReference:          "2",
		YourReference:         "1",
		CustomerOrderNumber:   "123",
		CustomerProjectNumber: "1",
		SupplierOrderNumber:   "123",
		InvoiceReference:      "123",
		Currency: hogia_api.Currency{
			Code: "SEK",
		},
	}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
