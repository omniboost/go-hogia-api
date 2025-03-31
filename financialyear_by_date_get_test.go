package hogia_api_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	hogia_api "github.com/omniboost/go-hogia-api"
)

func TestGetFinancialYearByDateRequest(t *testing.T) {
	req := client.NewGetFinancialYearByDateRequest()
	req.PathParams().OrgID = "b135d637-1688-4c55-9061-b0a40094461b"
	req.PathParams().Date = hogia_api.Date{time.Date(2025, 03, 24, 0, 0, 0, 0, time.UTC)}
	req.QueryParams().ApiVersion = hogia_api.Date{time.Date(2025, 03, 24, 0, 0, 0, 0, time.UTC)}
	// req.QueryParams().Page = 0
	// req.QueryParams().PageSize = 1
	resp, err, _ := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
