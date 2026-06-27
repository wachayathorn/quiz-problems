package design

type HttpRequest struct {
	method  string
	url     string
	headers map[string]string
	body    string
}

type HttpRequestBuilder struct {
	httpRequest HttpRequest
}

func NewHttpRequest() *HttpRequestBuilder {
	return &HttpRequestBuilder{
		httpRequest: HttpRequest{},
	}
}

func (b *HttpRequestBuilder) SetMethod(method string) *HttpRequestBuilder {
	b.httpRequest.method = method
	return b
}

func (b *HttpRequestBuilder) SetUrl(url string) *HttpRequestBuilder {
	b.httpRequest.url = url
	return b
}

func (b *HttpRequestBuilder) SetHeaders(headers map[string]string) *HttpRequestBuilder {
	b.httpRequest.headers = headers
	return b
}

func (b *HttpRequestBuilder) SetBody(body string) *HttpRequestBuilder {
	b.httpRequest.body = body
	return b
}

func (b *HttpRequestBuilder) Build() *HttpRequest {
	return &b.httpRequest
}
