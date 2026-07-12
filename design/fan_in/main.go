package fanin

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx := context.Background()

	ch1 := producer(1, 2, 3)
	ch2 := producer(4, 5)
	ch3 := producer(6, 7, 8, 9)

	merged := merge(ctx, ch1, ch2, ch3)

	count := 0
	for v := range merged { // ← ต้องหลุด loop ได้เอง = ข้อพิสูจน์ว่า close ถูกต้อง
		fmt.Println(v)
		count++
	}
	fmt.Println("total:", count) // ต้องได้ 9
}

// producer สร้าง channel ที่ส่งค่าแล้วปิดตัวเอง
func producer(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, n := range nums {
			ch <- n
		}
	}()
	return ch
}

func merge(ctx context.Context, channels ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	// แตก goroutine 1 ตัวต่อ 1 input channel
	// → แต่ละตัว "block รอ" ข้อมูลของตัวเอง ไม่ต้อง poll ไม่กิน CPU
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()    // จบงาน (channel ปิด หรือ ctx cancel) → แจ้ง WaitGroup
			for v := range c { // อ่านจนกว่า c จะถูก close → หลุด loop เอง
				select {
				case out <- v:
				case <-ctx.Done():
					return // ctx cancel → เลิกส่ง ป้องกัน goroutine leak
				}
			}
		}(ch)
	}

	// Closer goroutine: รอทุก forwarder จบ → ปิด out
	// นี่คือหน้าที่ของ WaitGroup ที่เราคุยกัน:
	// "ผู้ส่ง out มีหลายคน — ปิดได้เมื่อรู้ว่าทุกคนจบแล้วเท่านั้น"
	go func() {
		wg.Wait()
		close(out) // → for range merged ใน main หลุด loop
	}()

	return out // ← return ทันที! งานทั้งหมดอยู่ใน background
}
