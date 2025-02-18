package hogia_v2

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/omniboost/go-hogia-v2/utils"
)

func (c *Client) NewPostVoucherDraftsRequest() PostVoucherDraftsRequest {
	return PostVoucherDraftsRequest{
		client:      c,
		queryParams: c.NewPostVoucherDraftsQueryParams(),
		pathParams:  c.NewPostVoucherDraftsPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostVoucherDraftsRequestBody(),
	}
}

type PostVoucherDraftsRequest struct {
	client      *Client
	queryParams *PostVoucherDraftsQueryParams
	pathParams  *PostVoucherDraftsPathParams
	method      string
	headers     http.Header
	requestBody PostVoucherDraftsRequestBody
}

func (c *Client) NewPostVoucherDraftsQueryParams() *PostVoucherDraftsQueryParams {
	return &PostVoucherDraftsQueryParams{}
}

type PostVoucherDraftsQueryParams struct {
	ApiVersion Date `schema:"api-version"`
}

func (p PostVoucherDraftsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostVoucherDraftsRequest) QueryParams() *PostVoucherDraftsQueryParams {
	return r.queryParams
}

func (c *Client) NewPostVoucherDraftsPathParams() *PostVoucherDraftsPathParams {
	return &PostVoucherDraftsPathParams{}
}

type PostVoucherDraftsPathParams struct {
	OrgID string `schema:"org_id"`
}

func (p *PostVoucherDraftsPathParams) Params() map[string]string {
	return map[string]string{
		"org_id": p.OrgID,
	}
}

func (r *PostVoucherDraftsRequest) PathParams() *PostVoucherDraftsPathParams {
	return r.pathParams
}

func (r *PostVoucherDraftsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostVoucherDraftsRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostVoucherDraftsRequest) Method() string {
	return r.method
}

func (s *Client) NewPostVoucherDraftsRequestBody() PostVoucherDraftsRequestBody {
	return PostVoucherDraftsRequestBody{}
}

type PostVoucherDraftsRequestBody VoucherDraftsRequest

func (r *PostVoucherDraftsRequest) RequestBody() *PostVoucherDraftsRequestBody {
	return &r.requestBody
}

func (r *PostVoucherDraftsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostVoucherDraftsRequest) SetRequestBody(body PostVoucherDraftsRequestBody) {
	r.requestBody = body
}

func (r *PostVoucherDraftsRequest) NewResponseBody() *PostVoucherDraftsResponseBody {
	return &PostVoucherDraftsResponseBody{}
}

type PostVoucherDraftsResponseBody VoucherDraftsResponse

func (v *PostVoucherDraftsResponseBody) UnmarshalJSON(data []byte) error {
	strValue := string(data)
	strValue = strings.Trim(strValue, "\"")

	v.VoucherID = strValue
	return nil
}

func (r *PostVoucherDraftsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.org_id}}/voucherdrafts", r.PathParams())
	return &u
}

func (r *PostVoucherDraftsRequest) Do() (PostVoucherDraftsResponseBody, error) {
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
