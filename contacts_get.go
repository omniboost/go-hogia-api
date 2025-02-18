package hogia_v2

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-hogia-v2/utils"
)

func (c *Client) NewGetContactsRequest() GetContactsRequest {
	return GetContactsRequest{
		client:      c,
		queryParams: c.NewGetContactsQueryParams(),
		pathParams:  c.NewGetContactsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetContactsRequestBody(),
	}
}

type GetContactsRequest struct {
	client      *Client
	queryParams *GetContactsQueryParams
	pathParams  *GetContactsPathParams
	method      string
	headers     http.Header
	requestBody GetContactsRequestBody
}

func (c *Client) NewGetContactsQueryParams() *GetContactsQueryParams {
	return &GetContactsQueryParams{}
}

type GetContactsQueryParams struct {
	ApiVersion Date     `schema:"api-version"`
	Version    int      `schema:"version"`
	Page       int      `schema:"page"`
	PageSize   int      `schema:"pageSize"`
	Embed      []string `schema:"embed"`
}

func (p GetContactsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetContactsRequest) QueryParams() *GetContactsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetContactsPathParams() *GetContactsPathParams {
	return &GetContactsPathParams{}
}

type GetContactsPathParams struct {
	OrgID string `schema:"org_id"`
}

func (p *GetContactsPathParams) Params() map[string]string {
	return map[string]string{
		"org_id": p.OrgID,
	}
}

func (r *GetContactsRequest) PathParams() *GetContactsPathParams {
	return r.pathParams
}

func (r *GetContactsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetContactsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetContactsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetContactsRequestBody() GetContactsRequestBody {
	return GetContactsRequestBody{}
}

type GetContactsRequestBody struct {
}

func (r *GetContactsRequest) RequestBody() *GetContactsRequestBody {
	return nil
}

func (r *GetContactsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetContactsRequest) SetRequestBody(body GetContactsRequestBody) {
	r.requestBody = body
}

func (r *GetContactsRequest) NewResponseBody() *GetContactsResponseBody {
	return &GetContactsResponseBody{}
}

type GetContactsResponseBody []ContactsResponse

func (r *GetContactsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.org_id}}/contacts", r.PathParams())
	return &u
}

func (r *GetContactsRequest) Do() (GetContactsResponseBody, error) {
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
