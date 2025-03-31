package hogia_api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/omniboost/go-hogia-api/utils"
)

func (c *Client) NewPostCustomerInvoicesV2Request() PostCustomerInvoicesV2Request {
	return PostCustomerInvoicesV2Request{
		client:      c,
		queryParams: c.NewPostCustomerInvoicesV2QueryParams(),
		pathParams:  c.NewPostCustomerInvoicesV2PathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostCustomerInvoicesV2RequestBody(),
	}
}

type PostCustomerInvoicesV2Request struct {
	client      *Client
	queryParams *PostCustomerInvoicesV2QueryParams
	pathParams  *PostCustomerInvoicesV2PathParams
	method      string
	headers     http.Header
	requestBody PostCustomerInvoicesV2RequestBody
}

func (c *Client) NewPostCustomerInvoicesV2QueryParams() *PostCustomerInvoicesV2QueryParams {
	return &PostCustomerInvoicesV2QueryParams{}
}

type PostCustomerInvoicesV2QueryParams struct {
	ApiVersion Date `schema:"api-version"`
}

func (p PostCustomerInvoicesV2QueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostCustomerInvoicesV2Request) QueryParams() *PostCustomerInvoicesV2QueryParams {
	return r.queryParams
}

func (c *Client) NewPostCustomerInvoicesV2PathParams() *PostCustomerInvoicesV2PathParams {
	return &PostCustomerInvoicesV2PathParams{}
}

type PostCustomerInvoicesV2PathParams struct {
	OrgID string `schema:"org_id"`
}

func (p *PostCustomerInvoicesV2PathParams) Params() map[string]string {
	return map[string]string{
		"org_id": p.OrgID,
	}
}

func (r *PostCustomerInvoicesV2Request) PathParams() *PostCustomerInvoicesV2PathParams {
	return r.pathParams
}

func (r *PostCustomerInvoicesV2Request) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostCustomerInvoicesV2Request) SetMethod(method string) {
	r.method = method
}

func (r *PostCustomerInvoicesV2Request) Method() string {
	return r.method
}

func (s *Client) NewPostCustomerInvoicesV2RequestBody() PostCustomerInvoicesV2RequestBody {
	return PostCustomerInvoicesV2RequestBody{}
}

type PostCustomerInvoicesV2RequestBody CustomerInvoicesV2Request

func (r *PostCustomerInvoicesV2Request) RequestBody() *PostCustomerInvoicesV2RequestBody {
	return &r.requestBody
}

func (r *PostCustomerInvoicesV2Request) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostCustomerInvoicesV2Request) SetRequestBody(body PostCustomerInvoicesV2RequestBody) {
	r.requestBody = body
}

func (r *PostCustomerInvoicesV2Request) NewResponseBody() *PostCustomerInvoicesV2ResponseBody {
	return &PostCustomerInvoicesV2ResponseBody{}
}

type PostCustomerInvoicesV2ResponseBody CustomerInvoicesV2Response

func (v *PostCustomerInvoicesV2ResponseBody) UnmarshalJSON(data []byte) error {
	strValue := string(data)
	strValue = strings.Trim(strValue, "\"")

	v.CustomerInvoiceID = strValue
	return nil
}

func (r *PostCustomerInvoicesV2Request) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.org_id}}/customerinvoices", r.PathParams())
	return &u
}

func (r *PostCustomerInvoicesV2Request) Do() (PostCustomerInvoicesV2ResponseBody, error) {
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
