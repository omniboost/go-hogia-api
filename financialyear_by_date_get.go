package hogia_api

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-hogia-api/utils"
)

func (c *Client) NewGetFinancialYearByDateRequest() GetFinancialYearByDateRequest {
	return GetFinancialYearByDateRequest{
		client:      c,
		queryParams: c.NewGetFinancialYearByDateQueryParams(),
		pathParams:  c.NewGetFinancialYearByDatePathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetFinancialYearByDateRequestBody(),
	}
}

type GetFinancialYearByDateRequest struct {
	client      *Client
	queryParams *GetFinancialYearByDateQueryParams
	pathParams  *GetFinancialYearByDatePathParams
	method      string
	headers     http.Header
	requestBody GetFinancialYearByDateRequestBody
}

func (c *Client) NewGetFinancialYearByDateQueryParams() *GetFinancialYearByDateQueryParams {
	return &GetFinancialYearByDateQueryParams{}
}

type GetFinancialYearByDateQueryParams struct {
	ApiVersion Date `schema:"api-version"`
}

func (p GetFinancialYearByDateQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetFinancialYearByDateRequest) QueryParams() *GetFinancialYearByDateQueryParams {
	return r.queryParams
}

func (c *Client) NewGetFinancialYearByDatePathParams() *GetFinancialYearByDatePathParams {
	return &GetFinancialYearByDatePathParams{}
}

type GetFinancialYearByDatePathParams struct {
	OrgID string `schema:"org_id"`
	Date  Date   `schema:"date"`
}

func (p *GetFinancialYearByDatePathParams) Params() map[string]string {
	return map[string]string{
		"org_id": p.OrgID,
		"date":   p.Date.Format("2006-01-02"),
	}
}

func (r *GetFinancialYearByDateRequest) PathParams() *GetFinancialYearByDatePathParams {
	return r.pathParams
}

func (r *GetFinancialYearByDateRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetFinancialYearByDateRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetFinancialYearByDateRequest) Method() string {
	return r.method
}

func (s *Client) NewGetFinancialYearByDateRequestBody() GetFinancialYearByDateRequestBody {
	return GetFinancialYearByDateRequestBody{}
}

type GetFinancialYearByDateRequestBody struct {
}

func (r *GetFinancialYearByDateRequest) RequestBody() *GetFinancialYearByDateRequestBody {
	return nil
}

func (r *GetFinancialYearByDateRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetFinancialYearByDateRequest) SetRequestBody(body GetFinancialYearByDateRequestBody) {
	r.requestBody = body
}

func (r *GetFinancialYearByDateRequest) NewResponseBody() *GetFinancialYearByDateResponseBody {
	return &GetFinancialYearByDateResponseBody{}
}

type GetFinancialYearByDateResponseBody FinancialYearByDateResponse

func (r *GetFinancialYearByDateRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.org_id}}/financialyears/bydate/{{.date}}", r.PathParams())
	return &u
}

func (r *GetFinancialYearByDateRequest) Do() (GetFinancialYearByDateResponseBody, error, *http.Response) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err, nil
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, true)
	if err != nil {
		return *r.NewResponseBody(), err, nil
	}

	responseBody := r.NewResponseBody()
	resp, err := r.client.Do(req, responseBody)
	return *responseBody, err, resp
}
