package sort

import (
	"reflect"
	"testing"
)

type sortTester struct {
	input, output []int
}

var sortTests = []sortTester{
	sortTester{[]int{2, 3, 2, 5, 1}, []int{1, 2, 2, 3, 5}},
}

func TestBubbleSort(t *testing.T) {
	tests := []sortTester{
		sortTester{[]int{2, 3, 2, 5, 1}, []int{1, 2, 2, 3, 5}},
	}
	for _, v := range tests {
		BubbleSort(v.input)
		if !reflect.DeepEqual(v.input, v.output) {
			t.Errorf("want:%v, got:%v", v.output, v.input)
		}
	}
}

func TestQuickSort(t *testing.T) {
	tests := []sortTester{
		sortTester{[]int{2, 3, 2, 5, 1}, []int{1, 2, 2, 3, 5}},
	}
	for _, v := range tests {
		QuickSort(v.input)
		if !reflect.DeepEqual(v.input, v.output) {
			t.Errorf("want:%v, got:%v", v.output, v.input)
		}
	}
}

func TestInsertSort(t *testing.T) {
	tests := []sortTester{
		sortTester{[]int{2, 3, 2, 5, 1}, []int{1, 2, 2, 3, 5}},
	}
	for _, v := range tests {
		InsertSort(v.input)
		if !reflect.DeepEqual(v.input, v.output) {
			t.Errorf("want:%v, got:%v", v.output, v.input)
		}
	}
}

func TestShellSort(t *testing.T) {
	tests := []sortTester{
		sortTester{[]int{49, 38, 65, 97, 76, 13, 27, 49, 55, 4}, []int{4, 13, 27, 38, 49, 49, 55, 65, 76, 97}},
	}
	for _, v := range tests {
		ShellSort(v.input)
		if !reflect.DeepEqual(v.input, v.output) {
			t.Errorf("want:%v, got:%v", v.output, v.input)
		}
	}
}