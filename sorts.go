package sorts

// Partition implementation of Lomuto partition https://en.wikipedia.org/wiki/Quicksort
func Partition(input []int, lo, hi int) int {
	pivot := input[hi]
	i := lo
	for j := i; j < hi; j++ {
		if input[j] <= pivot {
			input[i], input[j] = input[j], input[i]
			i = i + 1
		}
	}
	input[i], input[hi] = input[hi], input[i]
	return i
}

// QuickSort implementation of quick sort
func QuickSort(input []int, lo, hi int) {
	if lo < hi {
		p := Partition(input, lo, hi)
		QuickSort(input, lo, p-1)
		QuickSort(input, p+1, hi)
	}
}

// ParallelThreshold on which sequential implementation should be chosen instead of parallel one
const ParallelThreshold = 10000

// ParallelQuickSort concurrent quick sort implementation that uses goroutines
func ParallelQuickSort(input []int, lo, hi int, done chan bool) {
	if lo < hi {
		if hi-lo > ParallelThreshold {
			p := Partition(input, lo, hi)
			jobsFinished := make(chan bool)
			go ParallelQuickSort(input, lo, p-1, jobsFinished)
			go ParallelQuickSort(input, p+1, hi, jobsFinished)
			<-jobsFinished
			<-jobsFinished
		} else {
			QuickSort(input, lo, hi)
		}
	}
	done <- true
}
