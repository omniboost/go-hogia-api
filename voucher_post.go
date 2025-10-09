package hogia_api

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-hogia-api/utils"
)

func (c *Client) NewPostVouchersRequest() PostVouchersRequest {
	return PostVouchersRequest{
		client:      c,
		queryParams: c.NewPostVouchersQueryParams(),
		pathParams:  c.NewPostVouchersPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostVouchersRequestBody(),
	}
}

type PostVouchersRequest struct {
	client      *Client
	queryParams *PostVouchersQueryParams
	pathParams  *PostVouchersPathParams
	method      string
	headers     http.Header
	requestBody PostVouchersRequestBody
}

func (c *Client) NewPostVouchersQueryParams() *PostVouchersQueryParams {
	return &PostVouchersQueryParams{}
}

type PostVouchersQueryParams struct {
	ApiVersion Date `schema:"api-version"`
}

func (p PostVouchersQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostVouchersRequest) QueryParams() *PostVouchersQueryParams {
	return r.queryParams
}

func (c *Client) NewPostVouchersPathParams() *PostVouchersPathParams {
	return &PostVouchersPathParams{}
}

type PostVouchersPathParams struct {
	OrgID         string `schema:"org_id"`
	FinanceyearID string `schema:"financeyear_id"`
}

func (p *PostVouchersPathParams) Params() map[string]string {
	return map[string]string{
		"org_id":         p.OrgID,
		"financeyear_id": p.FinanceyearID,
	}
}

func (r *PostVouchersRequest) PathParams() *PostVouchersPathParams {
	return r.pathParams
}

func (r *PostVouchersRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostVouchersRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostVouchersRequest) Method() string {
	return r.method
}

func (s *Client) NewPostVouchersRequestBody() PostVouchersRequestBody {
	return PostVouchersRequestBody{}
}

type PostVouchersRequestBody VouchersRequest

func (r *PostVouchersRequest) RequestBody() *PostVouchersRequestBody {
	return &r.requestBody
}

func (r *PostVouchersRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostVouchersRequest) SetRequestBody(body PostVouchersRequestBody) {
	r.requestBody = body
}

func (r *PostVouchersRequest) NewResponseBody() *PostVouchersResponseBody {
	return &PostVouchersResponseBody{}
}

type PostVouchersResponseBody VouchersResponse

// func (v *PostVouchersResponseBody) UnmarshalJSON(b []byte) (err error) {
// 	vc, vcs := AccountsResponse{}, []AccountsResponse{}
// 	if err = json.Unmarshal(b, &vc); err == nil {
// 		*v = make([]AccountsResponse, 1)
// 		[]AccountsResponse(*v)[0] = AccountsResponse(vc)
// 		return
// 	}
// 	if err = json.Unmarshal(b, &vcs); err == nil {
// 		*v = GetFinancialYearAccountsResponseBody(vcs)
// 		return
// 	}
// 	return
// }

func (r *PostVouchersRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.org_id}}/financialyears/{{.financeyear_id}}/vouchers", r.PathParams())
	return &u
}

func (r *PostVouchersRequest) Do() (PostVouchersResponseBody, error) {
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
