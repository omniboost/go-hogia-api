package hogia_api

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-hogia-api/utils"
)

func (c *Client) NewGetCurrenciesRequest() GetCurrenciesRequest {
	return GetCurrenciesRequest{
		client:      c,
		queryParams: c.NewGetCurrenciesQueryParams(),
		pathParams:  c.NewGetCurrenciesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetCurrenciesRequestBody(),
	}
}

type GetCurrenciesRequest struct {
	client      *Client
	queryParams *GetCurrenciesQueryParams
	pathParams  *GetCurrenciesPathParams
	method      string
	headers     http.Header
	requestBody GetCurrenciesRequestBody
}

func (c *Client) NewGetCurrenciesQueryParams() *GetCurrenciesQueryParams {
	return &GetCurrenciesQueryParams{}
}

type GetCurrenciesQueryParams struct {
	ApiVersion Date `schema:"api-version"`
}

func (p GetCurrenciesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetCurrenciesRequest) QueryParams() *GetCurrenciesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetCurrenciesPathParams() *GetCurrenciesPathParams {
	return &GetCurrenciesPathParams{}
}

type GetCurrenciesPathParams struct {
	OrgID string `schema:"org_id"`
}

func (p *GetCurrenciesPathParams) Params() map[string]string {
	return map[string]string{
		"org_id": p.OrgID,
	}
}

func (r *GetCurrenciesRequest) PathParams() *GetCurrenciesPathParams {
	return r.pathParams
}

func (r *GetCurrenciesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetCurrenciesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCurrenciesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetCurrenciesRequestBody() GetCurrenciesRequestBody {
	return GetCurrenciesRequestBody{}
}

type GetCurrenciesRequestBody struct {
}

func (r *GetCurrenciesRequest) RequestBody() *GetCurrenciesRequestBody {
	return nil
}

func (r *GetCurrenciesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetCurrenciesRequest) SetRequestBody(body GetCurrenciesRequestBody) {
	r.requestBody = body
}

func (r *GetCurrenciesRequest) NewResponseBody() *GetCurrenciesResponseBody {
	return &GetCurrenciesResponseBody{}
}

type GetCurrenciesResponseBody []CurrenciesResponse

func (r *GetCurrenciesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.org_id}}/currencies", r.PathParams())
	return &u
}

func (r *GetCurrenciesRequest) Do() (GetCurrenciesResponseBody, error) {
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
