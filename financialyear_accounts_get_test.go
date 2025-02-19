package hogia_api_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	hogia_api "github.com/omniboost/go-hogia-api"
)

func TestGetFinancialYearAccountsRequest(t *testing.T) {
	req := client.NewGetFinancialYearAccountsRequest()
	req.PathParams().OrgID = "b135d637-1688-4c55-9061-b0a40094461b"
	req.PathParams().FinanceyearID = "63c83542-f4ba-40d5-865e-b0a400c0cc19"
	req.QueryParams().ApiVersion = hogia_api.Date{time.Date(2025, 02, 13, 0, 0, 0, 0, time.UTC)}
	req.QueryParams().Number = 8999
	// req.QueryParams().MinNumber = 8000
	// req.QueryParams().MaxNumber = 9000
	// req.QueryParams().AccountPlanId = "64fa2e81-3f77-403b-adc0-9badfb22b827"
	// req.QueryParams().Start = 12
	// req.QueryParams().Count = 4
	// req.QueryParams().Page = 3
	// req.QueryParams().PageSize = 2
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
