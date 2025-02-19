package hogia_api

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/omniboost/go-hogia-api/utils"
)

func (c *Client) NewGetVatCodesRequest() GetVatCodesRequest {
	return GetVatCodesRequest{
		client:      c,
		queryParams: c.NewGetVatCodesQueryParams(),
		pathParams:  c.NewGetVatCodesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetVatCodesRequestBody(),
	}
}

type GetVatCodesRequest struct {
	client      *Client
	queryParams *GetVatCodesQueryParams
	pathParams  *GetVatCodesPathParams
	method      string
	headers     http.Header
	requestBody GetVatCodesRequestBody
}

func (c *Client) NewGetVatCodesQueryParams() *GetVatCodesQueryParams {
	return &GetVatCodesQueryParams{}
}

type GetVatCodesQueryParams struct {
	ApiVersion Date `schema:"api-version"`
	Number     int  `schema:"number,omitempty"`
}

func (p GetVatCodesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetVatCodesRequest) QueryParams() *GetVatCodesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetVatCodesPathParams() *GetVatCodesPathParams {
	return &GetVatCodesPathParams{}
}

type GetVatCodesPathParams struct {
	OrgID string `schema:"org_id"`
}

func (p *GetVatCodesPathParams) Params() map[string]string {
	return map[string]string{
		"org_id": p.OrgID,
	}
}

func (r *GetVatCodesRequest) PathParams() *GetVatCodesPathParams {
	return r.pathParams
}

func (r *GetVatCodesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetVatCodesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetVatCodesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetVatCodesRequestBody() GetVatCodesRequestBody {
	return GetVatCodesRequestBody{}
}

type GetVatCodesRequestBody struct {
}

func (r *GetVatCodesRequest) RequestBody() *GetVatCodesRequestBody {
	return nil
}

func (r *GetVatCodesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetVatCodesRequest) SetRequestBody(body GetVatCodesRequestBody) {
	r.requestBody = body
}

func (r *GetVatCodesRequest) NewResponseBody() *GetVatCodesResponseBody {
	return &GetVatCodesResponseBody{}
}

type GetVatCodesResponseBody []VatCodesResponse

func (v *GetVatCodesResponseBody) UnmarshalJSON(b []byte) (err error) {
	vc, vcs := VatCodesResponse{}, []VatCodesResponse{}
	if err = json.Unmarshal(b, &vc); err == nil {
		*v = make([]VatCodesResponse, 1)
		[]VatCodesResponse(*v)[0] = VatCodesResponse(vc)
		return
	}
	if err = json.Unmarshal(b, &vcs); err == nil {
		*v = GetVatCodesResponseBody(vcs)
		return
	}
	return
}

func (r *GetVatCodesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.org_id}}/vatcodes", r.PathParams())
	return &u
}

func (r *GetVatCodesRequest) Do() (GetVatCodesResponseBody, error) {
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
