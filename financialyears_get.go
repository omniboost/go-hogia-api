package hogia_api

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-hogia-api/utils"
)

func (c *Client) NewGetFinancialYearsRequest() GetFinancialYearsRequest {
	return GetFinancialYearsRequest{
		client:      c,
		queryParams: c.NewGetFinancialYearsQueryParams(),
		pathParams:  c.NewGetFinancialYearsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetFinancialYearsRequestBody(),
	}
}

type GetFinancialYearsRequest struct {
	client      *Client
	queryParams *GetFinancialYearsQueryParams
	pathParams  *GetFinancialYearsPathParams
	method      string
	headers     http.Header
	requestBody GetFinancialYearsRequestBody
}

func (c *Client) NewGetFinancialYearsQueryParams() *GetFinancialYearsQueryParams {
	return &GetFinancialYearsQueryParams{}
}

type GetFinancialYearsQueryParams struct {
	ApiVersion Date `schema:"api-version"`
	Page       int  `schema:"Page"`
	PageSize   int  `schema:"PageSize"`
	// Sort           string `schema:"Sort"`
	// SortProperties string `schema:"SortProperties"`
}

func (p GetFinancialYearsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetFinancialYearsRequest) QueryParams() *GetFinancialYearsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetFinancialYearsPathParams() *GetFinancialYearsPathParams {
	return &GetFinancialYearsPathParams{}
}

type GetFinancialYearsPathParams struct {
	OrgID string `schema:"org_id"`
}

func (p *GetFinancialYearsPathParams) Params() map[string]string {
	return map[string]string{
		"org_id": p.OrgID,
	}
}

func (r *GetFinancialYearsRequest) PathParams() *GetFinancialYearsPathParams {
	return r.pathParams
}

func (r *GetFinancialYearsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetFinancialYearsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetFinancialYearsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetFinancialYearsRequestBody() GetFinancialYearsRequestBody {
	return GetFinancialYearsRequestBody{}
}

type GetFinancialYearsRequestBody struct {
}

func (r *GetFinancialYearsRequest) RequestBody() *GetFinancialYearsRequestBody {
	return nil
}

func (r *GetFinancialYearsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetFinancialYearsRequest) SetRequestBody(body GetFinancialYearsRequestBody) {
	r.requestBody = body
}

func (r *GetFinancialYearsRequest) NewResponseBody() *GetFinancialYearsResponseBody {
	return &GetFinancialYearsResponseBody{}
}

type GetFinancialYearsResponseBody []FinancialYearsResponse

func (r *GetFinancialYearsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.org_id}}/financialyears/paged", r.PathParams())
	return &u
}

func (r *GetFinancialYearsRequest) Do() (GetFinancialYearsResponseBody, error, *http.Response) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err, nil
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err, nil
	}

	responseBody := r.NewResponseBody()
	resp, err := r.client.Do(req, responseBody)
	return *responseBody, err, resp
}

func (r *GetFinancialYearsRequest) All() ([]FinancialYearsResponse, error) {
	allYears := []FinancialYearsResponse{}
	for {
		body, err, resp := r.Do()
		if err != nil {
			return allYears, err
		}

		allYears = append(allYears, body...)

		// Get pagination info from headers
		currentPage := resp.Header.Get("X-Pagination-Currentpage")
		pageCount := resp.Header.Get("X-Pagination-Pagecount")

		// Convert header values to integers
		currentPageNum, _ := strconv.Atoi(currentPage)
		totalPages, _ := strconv.Atoi(pageCount)

		if currentPageNum == totalPages {
			break
		}
		r.QueryParams().Page = r.QueryParams().Page + 1
	}

	return allYears, nil
}
