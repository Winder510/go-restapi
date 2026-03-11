Lab 1 — Goroutine cơ bản
Bài toán

In số từ 1 → 5 bằng goroutine.

Yêu cầu

tạo goroutine

main phải đợi goroutine chạy xong

Gợi ý
go printNumbers()
Output
1
2
3
4
5
🧪 Lab 2 — Goroutine + sleep
Bài toán

Tạo 3 goroutine chạy song song.

worker A

worker B

worker C

Mỗi worker:

Worker A: step 1
Worker A: step 2
Worker A: step 3
Gợi ý
go worker("A")
go worker("B")
go worker("C")

Mỗi step sleep:

500ms

👉 Mục tiêu: thấy output interleave

🧪 Lab 3 — Channel cơ bản
Bài toán

Goroutine tính square của số.

Input

[1,2,3,4]

Output

1
4
9
16
Yêu cầu

goroutine tính toán

gửi kết quả qua channel

Gợi ý
resultChan := make(chan int)

worker:

resultChan <- num*num
🧪 Lab 4 — Fan-out / Fan-in pattern

Đây là pattern concurrency rất phổ biến trong Go.

Bài toán

Cho list numbers:

1..10

Tạo 3 worker goroutine xử lý song song.

Worker sẽ:

square number
Input
jobs = [1..10]
Output
1
4
9
16
25
...
Gợi ý
jobs channel
results channel
🧪 Lab 5 — Worker Pool

Đây là pattern production rất hay dùng.

Bài toán

Giả lập job queue.

jobs = 1..20

Tạo 5 worker goroutine

Worker sẽ:

process job
sleep random

Output ví dụ:

Worker 1 processing job 3
Worker 2 processing job 4
Worker 5 processing job 7
Gợi ý
jobs := make(chan int, 20)

worker

func worker(id int, jobs <-chan int)
🧪 Lab 6 — Rate Limiter

Bài này giống production API.

Bài toán

API chỉ cho phép:

5 request / second

Giả lập:

10 request
Gợi ý

dùng

time.Tick()

hoặc

time.NewTicker()
🧪 Lab 7 — Concurrent Web Crawler

Bài này xuất hiện trong tour Go nổi tiếng.

Bài toán

crawl website.

Ví dụ

example.com

Crawler sẽ:

fetch page

lấy các link

crawl tiếp

Yêu cầu

dùng goroutine

tránh crawl trùng URL

dùng mutex hoặc map

🧪 Lab 8 — Concurrent File Downloader
Bài toán

Cho list file:

file1
file2
file3
file4

Download song song.

Yêu cầu

3 goroutine

worker pool

Output

Downloading file1
Downloading file2
Downloading file3
Done file1
Done file2
🧪 Lab 9 — Pipeline Pattern

Một pattern concurrency rất Go-idiomatic.

Pipeline:

generate numbers
     ↓
square
     ↓
filter even
     ↓
print
Ví dụ

Input

1..10

Output

4
16
36
64
100
📚 Lab roadmap tốt nhất để master goroutine
1 goroutine
2 goroutine + wait
3 channel
4 fan-out fan-in
5 worker pool
6 rate limiter
7 pipeline
8 crawler

Nếu làm hết 8 bài này thì bạn gần như hiểu 80% concurrency Go.

💡 Nếu bạn muốn, tôi có thể đưa thêm:

10 bài goroutine giống câu hỏi phỏng vấn Go (khó hơn nhiều)

bài concurrency debug race condition (rất nhiều dev Go fail chỗ này)

**5 pattern concurrency production của các hệ thống như Uber và Cloudflare dùng trong Go.