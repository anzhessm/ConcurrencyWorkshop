package arithmetics

import (
	"sync"
	"sync/atomic"
	"time"
)

func SequentialSum(inputSize int) int {
	sum := 0
	for i := 1; i <= inputSize; i++ {
		sum += process(i)
	}
	return sum
}

// ParallelSum implement this method.
func ParallelSumReimplemented(inputSize int) int {
	wg := sync.WaitGroup{}
	sum := &atomic.Int64{}
	for i := 1; i <= inputSize; i++ {
		wg.Add(1)
		go func() {
			parallelProcess(sum, i)
			wg.Done()
		}()
	}
	wg.Wait()
	return int(sum.Load())
}

func ParallelSum(inputSize int) int {
	wg := sync.WaitGroup{}
	results := make(chan int, inputSize)
	for i := 1; i <= inputSize; i++ {
		wg.Add(1)
		go func(input int) {
			results <- process(input)
			wg.Done()
		}(i)
	}

	wg.Wait()

	close(results)
	var sum int
	for result := range results {
		sum += result
	}
	return sum
}

func parallelProcess(sum *atomic.Int64, num int) {
	res := num * num
	sum.Add(int64(res))
}

func process(num int) int {
	time.Sleep(time.Millisecond) // simulate processing time
	return num * num
}
