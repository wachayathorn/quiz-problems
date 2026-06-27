package main

import "github.com/wachayathorn/quiz-problems/design"

func main() {
	design.NewHttpRequest().
		SetMethod("GET").
		SetUrl("https://example.com").
		SetHeaders(map[string]string{
			"Authorization": "Bearer 1234567890",
		}).
		SetBody("{\"name\":\"John\",\"age\":30}").
		Build()
}
