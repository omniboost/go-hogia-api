package hogia_api_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	hogia_api "github.com/omniboost/go-hogia-api"
)

func TestGetCurrenciesRequest(t *testing.T) {
	req := client.NewGetCurrenciesRequest()
	req.PathParams().OrgID = "3cb44744-8fb9-4750-89b0-b29700871843"
	req.QueryParams().ApiVersion = hogia_api.Date{time.Date(2025, 06, 01, 0, 0, 0, 0, time.UTC)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
