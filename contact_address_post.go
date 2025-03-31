package hogia_api

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-hogia-api/utils"
)

func (c *Client) NewPostContactAddressRequest() PostContactAddressRequest {
	return PostContactAddressRequest{
		client:      c,
		queryParams: c.NewPostContactAddressQueryParams(),
		pathParams:  c.NewPostContactAddressPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostContactAddressRequestBody(),
	}
}

type PostContactAddressRequest struct {
	client      *Client
	queryParams *PostContactAddressQueryParams
	pathParams  *PostContactAddressPathParams
	method      string
	headers     http.Header
	requestBody PostContactAddressRequestBody
}

func (c *Client) NewPostContactAddressQueryParams() *PostContactAddressQueryParams {
	return &PostContactAddressQueryParams{}
}

type PostContactAddressQueryParams struct {
	ApiVersion Date `schema:"api-version"`
	Version    int  `schema:"version"`
}

func (p PostContactAddressQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostContactAddressRequest) QueryParams() *PostContactAddressQueryParams {
	return r.queryParams
}

func (c *Client) NewPostContactAddressPathParams() *PostContactAddressPathParams {
	return &PostContactAddressPathParams{}
}

type PostContactAddressPathParams struct {
	OrgID     string `schema:"org_id"`
	ContactID string `schema:"contact_id"`
}

func (p *PostContactAddressPathParams) Params() map[string]string {
	return map[string]string{
		"org_id":     p.OrgID,
		"contact_id": p.ContactID,
	}
}

func (r *PostContactAddressRequest) PathParams() *PostContactAddressPathParams {
	return r.pathParams
}

func (r *PostContactAddressRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostContactAddressRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostContactAddressRequest) Method() string {
	return r.method
}

func (s *Client) NewPostContactAddressRequestBody() PostContactAddressRequestBody {
	return PostContactAddressRequestBody{}
}

type PostContactAddressRequestBody ContactAddresssPost

func (r *PostContactAddressRequest) RequestBody() *PostContactAddressRequestBody {
	return &r.requestBody
}

func (r *PostContactAddressRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostContactAddressRequest) SetRequestBody(body PostContactAddressRequestBody) {
	r.requestBody = body
}

func (r *PostContactAddressRequest) NewResponseBody() *PostContactAddressResponseBody {
	return &PostContactAddressResponseBody{}
}

type PostContactAddressResponseBody ContactAddresssPostResponse

func (r *PostContactAddressRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.org_id}}/contacts/{{.contact_id}}/addresses", r.PathParams())
	return &u
}

func (r *PostContactAddressRequest) Do() (PostContactAddressResponseBody, error) {
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
