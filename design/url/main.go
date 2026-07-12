package url

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// สร้าง root context พร้อม timeout รวม 10 วินาที
	// ทุก goroutine ที่รับ ctx นี้ไปจะถูกยกเลิกพร้อมกันเมื่อครบเวลา
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // คืน resource ของ timer เสมอ แม้งานเสร็จก่อน timeout

	urls := []string{"https://example.com", "https://google.com", "https://github.com"}
	results := FetchAll(ctx, urls)

	for _, r := range results {
		if r.Err != nil {
			fmt.Printf("FAIL %s: %v\n", r.URL, r.Err)
			continue
		}
		fmt.Printf("OK   %s: %d bytes\n", r.URL, len(r.Body))
	}
}

type Result struct {
	URL  string
	Body []byte
	Err  error
}

// job พก index ติดไปด้วย เพื่อให้วางผลลัพธ์กลับตำแหน่งเดิมได้
// (ใช้ index แทน map[url] → รองรับกรณี URL ซ้ำกันใน input)
type job struct {
	index int
	url   string
}

// indexedResult ห่อ Result พร้อม index เดิม
// เพราะลำดับที่ผลออกจาก channel ไม่การันตีว่าตรงกับลำดับ input
type indexedResult struct {
	index  int
	result Result
}

const maxWorkers = 3 // จำกัด concurrency สูงสุด 3 requests พร้อมกัน

func FetchAll(ctx context.Context, urls []string) []Result {
	// สายพานงาน — unbuffered ก็ได้ เพราะเราส่งงานใน goroutine แยก (ไม่ block main)
	jobs := make(chan job)

	// สายพานผลลัพธ์ — buffer เท่าจำนวนงาน
	// → worker ส่งผลได้ทันทีไม่ต้องรอ main มาอ่าน → ไม่มีทาง deadlock
	resultCh := make(chan indexedResult, len(urls))

	// จอง slice ตามขนาด input → วางผลตาม index ได้เลย เรียงลำดับอัตโนมัติ
	results := make([]Result, len(urls))

	// สร้าง worker pool ตามจำนวนที่จำกัดไว้
	for range maxWorkers {
		go worker(ctx, jobs, resultCh)
	}

	// ส่งงานใน goroutine แยก → main ไม่ block ตอนส่งงาน
	// สำคัญ: กัน deadlock กรณี urls มากกว่า workers
	go func() {
		defer close(jobs) // ปิดสายพานงานเมื่อส่งครบ → เป็น signal ให้ worker หยุด
		for i, url := range urls {
			select {
			case jobs <- job{index: i, url: url}:
			case <-ctx.Done():
				return // ctx หมดเวลา → เลิกส่งงานที่เหลือ
			}
		}
	}()

	// เก็บผลด้วย "วิธีนับ" — รู้จำนวนล่วงหน้า (len(urls))
	// → ไม่ต้องปิด resultCh และไม่ต้องใช้ WaitGroup
	received := 0
	for received < len(urls) {
		select {
		case <-ctx.Done():
			// timeout → หยุดรอ แล้วไปเติม error ให้งานที่ไม่เสร็จด้านล่าง
			fillUnfinished(results, urls, ctx.Err())
			return results
		case r := <-resultCh:
			results[r.index] = r.result // วางผลตามตำแหน่ง index เดิม
			received++
		}
	}

	return results
}

// fillUnfinished เติม error ให้ URL ที่ยังไม่มีผล (กรณี timeout)
// → caller แยกได้ชัดว่าตัวไหนสำเร็จ ตัวไหนโดนยกเลิก ไม่มีผลหายเงียบ
func fillUnfinished(results []Result, urls []string, err error) {
	for i := range results {
		if results[i].URL == "" { // zero value = ยังไม่เคยได้รับผล
			results[i] = Result{URL: urls[i], Err: err}
		}
	}
}

// worker ดึงงานจากสายพานมาทำจนกว่างานหมดหรือ ctx ถูกยกเลิก
func worker(ctx context.Context, jobs <-chan job, results chan<- indexedResult) {
	for {
		select {
		case <-ctx.Done():
			return // ctx หมดเวลา → หยุดทันที
		case j, ok := <-jobs:
			if !ok {
				return // channel ถูกปิด = งานหมดแล้ว → หยุด
				// สำคัญ: ถ้าไม่เช็ค ok จะได้ zero value วนไม่รู้จบ!
			}
			// ส่งผลได้เสมอไม่ block เพราะ resultCh buffer = len(urls)
			results <- indexedResult{index: j.index, result: fetch(ctx, j.url)}
		}
	}
}

// fetch ยิง HTTP GET จริง โดยผูก request กับ ctx
// → ถ้า ctx หมดเวลากลางทาง request จะถูกยกเลิกทันที ไม่ค้าง
func fetch(ctx context.Context, url string) Result {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Result{URL: url, Err: err}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Result{URL: url, Err: err}
	}
	defer resp.Body.Close() // ปิด body เสมอ → คืน connection กลับ pool

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{URL: url, Err: err}
	}

	return Result{URL: url, Body: body}
}
