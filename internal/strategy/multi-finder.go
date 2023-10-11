package strategy

import (
	"context"
	"runtime"
)

// multiFinder struct
type multiFinder struct {
	cpus int
}

func (f *multiFinder) subSearch(ctx context.Context, outChan chan<- bool, array []string, value string) {
	for _, elem := range array {
		select {
		// stop iterating
		case <-ctx.Done():
			return
		// check element
		default:
			if elem == value {
				outChan <- true

				return
			}
		}
	}

	outChan <- false

	return
}

// search method
func (f *multiFinder) search(array []string, value string) bool {
	// chan for search status, use in goroutines
	outChan := make(chan bool, f.cpus)
	defer close(outChan)

	// stop goroutines after first coincidence
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	// array length for goroutines
	arrSubLen := len(array) / f.cpus

	for i := 0; i < (f.cpus); i++ {
		if i == f.cpus-1 { // if length of array isn't completely divisible, for last step
			go f.subSearch(ctx, outChan, array[i*arrSubLen:], value)
		} else {
			go f.subSearch(ctx, outChan, array[i*arrSubLen:(i+1)*arrSubLen], value)
		}
	}

	// read chan while waiting "true"
	for i := 0; i < f.cpus; i++ {
		select {
		case res := <-outChan:
			if res == true {
				return true
			}
		}
	}

	return false
}

// NewMultiFinder fabric for multiFinder with Strategy interface
func NewMultiFinder() Strategy {
	return &multiFinder{cpus: runtime.NumCPU()}
}
