package builder

import "fmt"

type HTTPRequest struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    string
}

type HTTPRequestBuilder struct {
	request HTTPRequest
}

func NewHTTPRequest() *HTTPRequestBuilder {
	return &HTTPRequestBuilder{}
}

func (b *HTTPRequestBuilder) SetMethod(method string) *HTTPRequestBuilder {
	b.request.Method = method
	return b
}

func (b *HTTPRequestBuilder) SetURL(url string) *HTTPRequestBuilder {
	b.request.URL = url
	return b
}

func (b *HTTPRequestBuilder) SetHeaders(headers map[string]string) *HTTPRequestBuilder {
	b.request.Headers = headers
	return b
}

func (b *HTTPRequestBuilder) SetBody(body string) *HTTPRequestBuilder {
	b.request.Body = body
	return b
}

func (b *HTTPRequestBuilder) Build() *HTTPRequest {
	return &b.request
}

func main() {
	builder := NewHTTPRequest()
	builder.SetMethod("GET")
	builder.SetURL("https://example.com")
	builder.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})
	builder.SetBody("{\"name\":\"John\",\"age\":30}")
	request := builder.Build()
	fmt.Println(request)
}
