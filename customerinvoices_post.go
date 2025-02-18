package hogia_v2

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-hogia-v2/utils"
)

func (c *Client) NewPostCustomerInvoicesRequest() PostCustomerInvoicesRequest {
	return PostCustomerInvoicesRequest{
		client:      c,
		queryParams: c.NewPostCustomerInvoicesQueryParams(),
		pathParams:  c.NewPostCustomerInvoicesPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostCustomerInvoicesRequestBody(),
	}
}

type PostCustomerInvoicesRequest struct {
	client      *Client
	queryParams *PostCustomerInvoicesQueryParams
	pathParams  *PostCustomerInvoicesPathParams
	method      string
	headers     http.Header
	requestBody PostCustomerInvoicesRequestBody
}

func (c *Client) NewPostCustomerInvoicesQueryParams() *PostCustomerInvoicesQueryParams {
	return &PostCustomerInvoicesQueryParams{}
}

type PostCustomerInvoicesQueryParams struct {
	ApiVersion Date `schema:"api-version"`
}

func (p PostCustomerInvoicesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostCustomerInvoicesRequest) QueryParams() *PostCustomerInvoicesQueryParams {
	return r.queryParams
}

func (c *Client) NewPostCustomerInvoicesPathParams() *PostCustomerInvoicesPathParams {
	return &PostCustomerInvoicesPathParams{}
}

type PostCustomerInvoicesPathParams struct {
	OrgID string `schema:"org_id"`
}

func (p *PostCustomerInvoicesPathParams) Params() map[string]string {
	return map[string]string{
		"org_id": p.OrgID,
	}
}

func (r *PostCustomerInvoicesRequest) PathParams() *PostCustomerInvoicesPathParams {
	return r.pathParams
}

func (r *PostCustomerInvoicesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostCustomerInvoicesRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostCustomerInvoicesRequest) Method() string {
	return r.method
}

func (s *Client) NewPostCustomerInvoicesRequestBody() PostCustomerInvoicesRequestBody {
	return PostCustomerInvoicesRequestBody{}
}

type PostCustomerInvoicesRequestBody CustomerInvoicesRequest

func (r *PostCustomerInvoicesRequest) RequestBody() *PostCustomerInvoicesRequestBody {
	return &r.requestBody
}

func (r *PostCustomerInvoicesRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostCustomerInvoicesRequest) SetRequestBody(body PostCustomerInvoicesRequestBody) {
	r.requestBody = body
}

func (r *PostCustomerInvoicesRequest) NewResponseBody() *PostCustomerInvoicesResponseBody {
	return &PostCustomerInvoicesResponseBody{}
}

type PostCustomerInvoicesResponseBody CustomerInvoicesResponse

func (r *PostCustomerInvoicesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.org_id}}/customerinvoices", r.PathParams())
	return &u
}

func (r *PostCustomerInvoicesRequest) Do() (PostCustomerInvoicesResponseBody, error) {
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
