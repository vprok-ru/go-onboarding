package strategy

import (
	"context"
	"runtime"
)

type multiFinder struct {
	out  chan bool
	cpus int
}

func (f *multiFinder) subSearch(ctx context.Context, array []string, value string) {
	for _, elem := range array {
		select {
		// stop iterating
		case <-ctx.Done():
			return
		// check element
		default:
			if elem == value {
				f.out <- true

				return
			}
		}
	}

	f.out <- false

	return
}

// search
func (f *multiFinder) search(array []string, value string) bool {
	// chan for search status, use in goroutines
	f.out = make(chan bool, f.cpus)
	defer close(f.out)

	// stop goroutines after first coincidence
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	// array length for goroutines
	arrSubLen := len(array) / f.cpus

	for i := 0; i < (f.cpus); i++ {
		if i == f.cpus-1 { // if length of array isn't completely divisible, for last step
			go f.subSearch(ctx, array[i*arrSubLen:], value)
		} else {
			go f.subSearch(ctx, array[i*arrSubLen:(i+1)*arrSubLen], value)
		}
	}

	// read chan while waiting "true"
	for i := 0; i < f.cpus; i++ {
		select {
		case res := <-f.out:
			if res == true {
				return true
			}
		}
	}

	return false
}

func NewMultiFinder() Strategy {
	return &multiFinder{cpus: runtime.NumCPU()}
}
