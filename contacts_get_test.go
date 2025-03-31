package hogia_api_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	hogia_api "github.com/omniboost/go-hogia-api"
)

func TestGetContactsRequest(t *testing.T) {
	req := client.NewGetContactsRequest()
	req.PathParams().OrgID = "3cb44744-8fb9-4750-89b0-b29700871843"
	// req.PathParams().ContactID = "df52c06b-a79d-4212-a38b-b2a700cf113c"
	req.QueryParams().ApiVersion = hogia_api.Date{time.Date(2025, 03, 21, 0, 0, 0, 0, time.UTC)}
	// req.QueryParams().Page = 0
	req.QueryParams().Version = 0
	// req.QueryParams().Embed = []string{"ContactChannels.Addresses"}
	// req.QueryParams().PageSize = 100
	req.QueryParams().Number = "1"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
