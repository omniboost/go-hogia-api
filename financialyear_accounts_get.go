package hogia_api

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/omniboost/go-hogia-api/utils"
)

func (c *Client) NewGetFinancialYearAccountsRequest() GetFinancialYearAccountsRequest {
	return GetFinancialYearAccountsRequest{
		client:      c,
		queryParams: c.NewGetFinancialYearAccountsQueryParams(),
		pathParams:  c.NewGetFinancialYearAccountsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetFinancialYearAccountsRequestBody(),
	}
}

type GetFinancialYearAccountsRequest struct {
	client      *Client
	queryParams *GetFinancialYearAccountsQueryParams
	pathParams  *GetFinancialYearAccountsPathParams
	method      string
	headers     http.Header
	requestBody GetFinancialYearAccountsRequestBody
}

func (c *Client) NewGetFinancialYearAccountsQueryParams() *GetFinancialYearAccountsQueryParams {
	return &GetFinancialYearAccountsQueryParams{}
}

type GetFinancialYearAccountsQueryParams struct {
	ApiVersion      Date         `schema:"api-version"`
	Number          int          `schema:"number,omitempty"`
	MinNumber       int          `schema:"minNumber,omitempty"`
	MaxNumber       int          `schema:"maxNumber,omitempty"`
	AccountPlanId   string       `schema:"accountPlanId,omitempty"`
	Start           int          `schema:"start,omitempty"`
	Count           int          `schema:"count,omitempty"`
	Page            int          `schema:"Page,omitempty"`
	PageSize        int          `schema:"PageSize,omitempty"`
	AccountState    AccountState `schema:"accountState,omitempty"`
	Sort            string       `schema:"sort,omitempty"`
	SortProperties  []string     `schema:"sortProperties,omitempty"`
	SearchParameter string       `schema:"searchParameter,omitempty"`
}

func (p GetFinancialYearAccountsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetFinancialYearAccountsRequest) QueryParams() *GetFinancialYearAccountsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetFinancialYearAccountsPathParams() *GetFinancialYearAccountsPathParams {
	return &GetFinancialYearAccountsPathParams{}
}

type GetFinancialYearAccountsPathParams struct {
	OrgID         string `schema:"org_id"`
	FinanceyearID string `schema:"financeyear_id"`
}

func (p *GetFinancialYearAccountsPathParams) Params() map[string]string {
	return map[string]string{
		"org_id":         p.OrgID,
		"financeyear_id": p.FinanceyearID,
	}
}

func (r *GetFinancialYearAccountsRequest) PathParams() *GetFinancialYearAccountsPathParams {
	return r.pathParams
}

func (r *GetFinancialYearAccountsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetFinancialYearAccountsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetFinancialYearAccountsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetFinancialYearAccountsRequestBody() GetFinancialYearAccountsRequestBody {
	return GetFinancialYearAccountsRequestBody{}
}

type GetFinancialYearAccountsRequestBody struct {
}

func (r *GetFinancialYearAccountsRequest) RequestBody() *GetFinancialYearAccountsRequestBody {
	return nil
}

func (r *GetFinancialYearAccountsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetFinancialYearAccountsRequest) SetRequestBody(body GetFinancialYearAccountsRequestBody) {
	r.requestBody = body
}

func (r *GetFinancialYearAccountsRequest) NewResponseBody() *GetFinancialYearAccountsResponseBody {
	return &GetFinancialYearAccountsResponseBody{}
}

type GetFinancialYearAccountsResponseBody []AccountsResponse

func (v *GetFinancialYearAccountsResponseBody) UnmarshalJSON(b []byte) (err error) {
	vc, vcs := AccountsResponse{}, []AccountsResponse{}
	if err = json.Unmarshal(b, &vc); err == nil {
		*v = make([]AccountsResponse, 1)
		[]AccountsResponse(*v)[0] = AccountsResponse(vc)
		return
	}
	if err = json.Unmarshal(b, &vcs); err == nil {
		*v = GetFinancialYearAccountsResponseBody(vcs)
		return
	}
	return
}

func (r *GetFinancialYearAccountsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.org_id}}/financialyears/{{.financeyear_id}}/accounts", r.PathParams())
	return &u
}

func (r *GetFinancialYearAccountsRequest) Do() (GetFinancialYearAccountsResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, true)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
