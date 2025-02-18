package hogia_v2_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	hogia_v2 "github.com/omniboost/go-hogia-v2"
)

func TestGetVatCodesRequest(t *testing.T) {
	req := client.NewGetVatCodesRequest()
	req.PathParams().OrgID = "b135d637-1688-4c55-9061-b0a40094461b"
	req.QueryParams().ApiVersion = hogia_v2.Date{time.Date(2025, 02, 13, 0, 0, 0, 0, time.UTC)}
	req.QueryParams().Number = 2
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
