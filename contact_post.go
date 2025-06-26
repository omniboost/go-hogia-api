package hogia_api

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-hogia-api/utils"
)

func (c *Client) NewPostContactRequest() PostContactRequest {
	return PostContactRequest{
		client:      c,
		queryParams: c.NewPostContactQueryParams(),
		pathParams:  c.NewPostContactPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostContactRequestBody(),
	}
}

type PostContactRequest struct {
	client      *Client
	queryParams *PostContactQueryParams
	pathParams  *PostContactPathParams
	method      string
	headers     http.Header
	requestBody PostContactRequestBody
}

func (c *Client) NewPostContactQueryParams() *PostContactQueryParams {
	return &PostContactQueryParams{}
}

type PostContactQueryParams struct {
	ApiVersion Date `schema:"api-version"`
	Version    int  `schema:"version"`
}

func (p PostContactQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostContactRequest) QueryParams() *PostContactQueryParams {
	return r.queryParams
}

func (c *Client) NewPostContactPathParams() *PostContactPathParams {
	return &PostContactPathParams{}
}

type PostContactPathParams struct {
	OrgID string `schema:"org_id"`
	ID string    `schema:"id,omitempty"`
}

func (p *PostContactPathParams) Params() map[string]string {
	return map[string]string{
		"org_id": p.OrgID,
        "id":     p.ID,
	}
}

func (r *PostContactRequest) PathParams() *PostContactPathParams {
	return r.pathParams
}

func (r *PostContactRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostContactRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostContactRequest) Method() string {
	return r.method
}

func (s *Client) NewPostContactRequestBody() PostContactRequestBody {
	return PostContactRequestBody{}
}

type PostContactRequestBody ContactsPost

func (r *PostContactRequest) RequestBody() *PostContactRequestBody {
	return &r.requestBody
}

func (r *PostContactRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostContactRequest) SetRequestBody(body PostContactRequestBody) {
	r.requestBody = body
}

func (r *PostContactRequest) NewResponseBody() *PostContactResponseBody {
	return &PostContactResponseBody{}
}

type PostContactResponseBody ContactsPostResponse

func (r *PostContactRequest) URL() *url.URL {
	if r.PathParams().ID != "" {
		u := r.client.GetEndpointURL("{{.org_id}}/contacts/{{.id}}", r.PathParams())
		return &u
	}

	u := r.client.GetEndpointURL("{{.org_id}}/contacts", r.PathParams())
	return &u
}

func (r *PostContactRequest) Do() (PostContactResponseBody, error) {
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
