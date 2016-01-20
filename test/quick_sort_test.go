package test

import (
	"github.com/YaroslavDev/sorts"
	"math/rand"
	"testing"
	"time"
	"sort"
)

const N = 100

func TestQuickSort(t *testing.T) {
	testData := generateInput(N)

	sorts.QuickSort(testData, 0, len(testData)-1)

	successIfSorted(testData, t)
}

func TestParallelQuickSort(t *testing.T) {
	testData := generateInput(N)

	jobDone := make(chan bool)
	go sorts.ParallelQuickSort(testData, 0, len(testData)-1, jobDone)
	<-jobDone

	successIfSorted(testData, t)
}

func BenchmarkQuickSort(b *testing.B) {
	b.StopTimer()
	testData := generateInput(b.N)
	b.StartTimer()

	sorts.QuickSort(testData, 0, len(testData)-1)
}

func BenchmarkParallelQuickSort(b *testing.B) {
	b.StopTimer()
	testData := generateInput(b.N)
	b.StartTimer()

	jobDone := make(chan bool)
	go sorts.ParallelQuickSort(testData, 0, len(testData)-1, jobDone)
	<-jobDone
}

func BenchmarkNativeSort(b *testing.B) {
	b.StopTimer()
	testData := generateInput(b.N)
	b.StartTimer()

	sort.Ints(testData)
}

func generateInput(N int) []int {
	rand.Seed(int64(time.Now().Second()))
	input := make([]int, N)
	for index, _ := range input {
		input[index] = rand.Intn(N)
	}
	return input
}

func successIfSorted(arr []int, t *testing.T) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			t.Error("Array is not sorted")
		}
	}
}
