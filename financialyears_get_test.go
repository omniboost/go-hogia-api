package hogia_v2_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	hogia_v2 "github.com/omniboost/go-hogia-v2"
)

func TestGetFinancialYearsRequest(t *testing.T) {
	req := client.NewGetFinancialYearsRequest()
	req.PathParams().OrgID = "b135d637-1688-4c55-9061-b0a40094461b"
	req.QueryParams().ApiVersion = hogia_v2.Date{time.Date(2025, 02, 13, 0, 0, 0, 0, time.UTC)}
	// req.QueryParams().Page = 0
	// req.QueryParams().PageSize = 1
	resp, err, _ := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

func TestGetFinancialYearsAll(t *testing.T) {
	req := client.NewGetFinancialYearsRequest()
	req.PathParams().OrgID = "b135d637-1688-4c55-9061-b0a40094461b"
	req.QueryParams().ApiVersion = hogia_v2.Date{time.Date(2025, 02, 13, 0, 0, 0, 0, time.UTC)}
	req.QueryParams().Page = 1
	req.QueryParams().PageSize = 1
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
