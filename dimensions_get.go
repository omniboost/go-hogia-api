package hogia_api

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-hogia-api/utils"
)

func (c *Client) NewGetDimensionsRequest() GetDimensionsRequest {
	return GetDimensionsRequest{
		client:      c,
		queryParams: c.NewGetDimensionsQueryParams(),
		pathParams:  c.NewGetDimensionsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetDimensionsRequestBody(),
	}
}

type GetDimensionsRequest struct {
	client      *Client
	queryParams *GetDimensionsQueryParams
	pathParams  *GetDimensionsPathParams
	method      string
	headers     http.Header
	requestBody GetDimensionsRequestBody
}

func (c *Client) NewGetDimensionsQueryParams() *GetDimensionsQueryParams {
	return &GetDimensionsQueryParams{}
}

type GetDimensionsQueryParams struct {
	ApiVersion     Date           `schema:"api-version"`
	DimensionType  DimensionType  `schema:"dimensionType"`
	Page           int            `schema:"page"`
	PageSize       int            `schema:"pageSize"`
	DimensionState DimensionState `schema:"dimensionState"`
	SearchText     string         `schema:"searchText"`
	Version        int            `schema:"version"`
	Type           string         `schema:"type"`
	Number         string         `schema:"number"`
	Sort           string         `schema:"sort"`
	SortProperties []string       `schema:"sortProperties"`
}

func (p GetDimensionsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetDimensionsRequest) QueryParams() *GetDimensionsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetDimensionsPathParams() *GetDimensionsPathParams {
	return &GetDimensionsPathParams{}
}

type GetDimensionsPathParams struct {
	OrgID string `schema:"org_id"`
}

func (p *GetDimensionsPathParams) Params() map[string]string {
	return map[string]string{
		"org_id": p.OrgID,
	}
}

func (r *GetDimensionsRequest) PathParams() *GetDimensionsPathParams {
	return r.pathParams
}

func (r *GetDimensionsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetDimensionsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetDimensionsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetDimensionsRequestBody() GetDimensionsRequestBody {
	return GetDimensionsRequestBody{}
}

type GetDimensionsRequestBody struct {
}

func (r *GetDimensionsRequest) RequestBody() *GetDimensionsRequestBody {
	return nil
}

func (r *GetDimensionsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetDimensionsRequest) SetRequestBody(body GetDimensionsRequestBody) {
	r.requestBody = body
}

func (r *GetDimensionsRequest) NewResponseBody() *GetDimensionsResponseBody {
	return &GetDimensionsResponseBody{}
}

type GetDimensionsResponseBody []DimensionsResponse

func (r *GetDimensionsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.org_id}}/dimensions", r.PathParams())
	return &u
}

func (r *GetDimensionsRequest) Do() (GetDimensionsResponseBody, error) {
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
