package semaphore

import (
	"context"
)

func main() {
	urls := []string{"url1", "url2", "url3", "url4", "url5"}
	results := make([]string, len(urls))
	ctx := context.Background()
	sem := make(chan struct{}, 3)
	for i, url := range urls {
		sem <- struct{}{} // block ถ้ามี 3 ตัวกำลังรันอยู่
		go func(i int, url string) {
			defer func() { <-sem }()
			results[i] = fetch(ctx, url) // แต่ละ index เขียนช่องตัวเอง → ไม่ race
		}(i, url)
	}
}

func fetch(ctx context.Context, url string) string {
	// implementation
	return ""
}
