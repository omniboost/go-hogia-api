package hogia_api_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	hogia_api "github.com/omniboost/go-hogia-api"
)

func TestGetDimensionsRequest(t *testing.T) {
	req := client.NewGetDimensionsRequest()
	req.PathParams().OrgID = "b135d637-1688-4c55-9061-b0a40094461b"
	req.QueryParams().ApiVersion = hogia_api.Date{time.Date(2025, 02, 13, 0, 0, 0, 0, time.UTC)}
	// req.QueryParams().Page = 3
	// req.QueryParams().DimensionType = hogia_api.DimensionTypeTwo
	// req.QueryParams().Version = 1
	// req.QueryParams().PageSize = 1
	// req.QueryParams().Number = "1"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
