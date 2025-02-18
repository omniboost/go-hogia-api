package hogia_v2

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-hogia-v2/utils"
)

func (c *Client) NewGetVatSettingsRequest() GetVatSettingsRequest {
	return GetVatSettingsRequest{
		client:      c,
		queryParams: c.NewGetVatSettingsQueryParams(),
		pathParams:  c.NewGetVatSettingsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetVatSettingsRequestBody(),
	}
}

type GetVatSettingsRequest struct {
	client      *Client
	queryParams *GetVatSettingsQueryParams
	pathParams  *GetVatSettingsPathParams
	method      string
	headers     http.Header
	requestBody GetVatSettingsRequestBody
}

func (c *Client) NewGetVatSettingsQueryParams() *GetVatSettingsQueryParams {
	return &GetVatSettingsQueryParams{}
}

type GetVatSettingsQueryParams struct {
	ApiVersion Date `schema:"api-version"`
	Page       int  `schema:"page"`
	PageSize   int  `schema:"pageSize"`
	Version    int  `schema:"version"`
}

func (p GetVatSettingsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetVatSettingsRequest) QueryParams() *GetVatSettingsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetVatSettingsPathParams() *GetVatSettingsPathParams {
	return &GetVatSettingsPathParams{}
}

type GetVatSettingsPathParams struct {
	OrgID string `schema:"org_id"`
}

func (p *GetVatSettingsPathParams) Params() map[string]string {
	return map[string]string{
		"org_id": p.OrgID,
	}
}

func (r *GetVatSettingsRequest) PathParams() *GetVatSettingsPathParams {
	return r.pathParams
}

func (r *GetVatSettingsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetVatSettingsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetVatSettingsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetVatSettingsRequestBody() GetVatSettingsRequestBody {
	return GetVatSettingsRequestBody{}
}

type GetVatSettingsRequestBody struct {
}

func (r *GetVatSettingsRequest) RequestBody() *GetVatSettingsRequestBody {
	return nil
}

func (r *GetVatSettingsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetVatSettingsRequest) SetRequestBody(body GetVatSettingsRequestBody) {
	r.requestBody = body
}

func (r *GetVatSettingsRequest) NewResponseBody() *GetVatSettingsResponseBody {
	return &GetVatSettingsResponseBody{}
}

type GetVatSettingsResponseBody []VatSettingsResponse

func (r *GetVatSettingsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.org_id}}/vatsettings", r.PathParams())
	return &u
}

func (r *GetVatSettingsRequest) Do() (GetVatSettingsResponseBody, error) {
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
