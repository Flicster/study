package patterns

import "sync"

func Pool() {
	resCh := make(chan int, 100)
	jobsCh := make(chan int, 100)

	for x := 0; x < 3; x++ {
		go worker(func(i int) int {
			return i * i
		}, jobsCh, resCh)
	}

	for x := 0; x < 100; x++ {
		jobsCh <- x
	}
	close(jobsCh)

	for x := 0; x < 100; x++ {
		println(<-resCh)
	}
	close(resCh)
}

func worker(w func(int) int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- w(j)
	}
}

// ---------------

func Pipeline() {
	firstCh := make(chan int, 10)
	secondCh := make(chan int, 10)

	go func() {
		for x := 0; x < 10; x++ {
			firstCh <- x
		}
		close(firstCh)
	}()

	go func() {
		for f := range firstCh {
			secondCh <- f
		}
		close(secondCh)
	}()

	for s := range secondCh {
		print(s)
	}
}

// ---------------

func Merge() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go func() {
		for x := 0; x < 10; x++ {
			first <- 1
		}
		close(first)
	}()
	go func() {
		for x := 0; x < 10; x++ {
			second <- 2
		}
		close(second)
	}()
	go func() {
		for x := 0; x < 10; x++ {
			third <- 3
		}
		close(third)
	}()

	resCh := merge(first, second, third)

	for r := range resCh {
		println(r)
	}
}

func merge(data ...<-chan int) <-chan int {
	resCh := make(chan int)

	go func() {
		defer close(resCh)

		wg := &sync.WaitGroup{}
		wg.Add(len(data))
		for _, d := range data {
			d := d
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for x := range d {
					resCh <- x
				}
			}(d, wg)
		}
		wg.Wait()
	}()
	return resCh
}
