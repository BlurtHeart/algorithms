package matrix

import (
	"reflect"
	"testing"
)

type rotateTester struct {
	input, output [][]int
}

func TestRotate(t *testing.T) {
	tests := []rotateTester{
		rotateTester{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[][]int{
				[]int{7, 4, 1},
				[]int{8, 5, 2},
				[]int{9, 6, 3},
			},
		},
	}
	for _, v := range tests {
		Rotate(v.input)
		for i := 0; i < len(v.input); i++ {
			if !reflect.DeepEqual(v.input[i], v.output[i]) {
				t.Errorf("want:%v, got:%v", v.output, v.input)
			}
		}
	}
}

func TestAntiRotate(t *testing.T) {
	tests := []rotateTester{
		rotateTester{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[][]int{
				[]int{3, 6, 9},
				[]int{2, 5, 8},
				[]int{1, 4, 7},
			},
		},
	}
	for _, v := range tests {
		AntiRotate(v.input)
		for i := 0; i < len(v.input); i++ {
			if !reflect.DeepEqual(v.input[i], v.output[i]) {
				t.Errorf("want:%v, got:%v", v.output, v.input)
			}
		}
	}
}

func TestSpiralOrder(t *testing.T) {
	type tester struct {
		input  [][]int
		output []int
	}
	tests := []tester{
		tester{
			input: [][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			output: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		tester{
			input: [][]int{
				[]int{1, 2, 3},
			},
			output: []int{1, 2, 3},
		},
		tester{
			input: [][]int{
				[]int{1},
				[]int{2},
				[]int{3},
			},
			output: []int{1, 2, 3},
		},
	}
	for _, v := range tests {
		got := SpiralOrder(v.input)
		if !reflect.DeepEqual(v.output, got) {
			t.Errorf("want:%v, got:%v", v.output, got)
		}
	}
}

func TestGenerateSpiralMatrix(t *testing.T) {
	type tester struct {
		input  int
		output [][]int
	}
	tests := []tester{
		tester{
			input: 3,
			output: [][]int{
				[]int{1, 2, 3},
				[]int{8, 9, 4},
				[]int{7, 6, 5},
			},
		},
	}
	for _, v := range tests {
		got := GenerateSpiralMatrix(v.input)
		if !reflect.DeepEqual(v.output, got) {
			t.Errorf("want:%v, got:%v", v.output, got)
		}
	}
}

func TestSearchMatrix(t *testing.T) {
	type tester struct {
		input  [][]int
		target int
		output bool
	}
	tests := []tester{
		tester{
			input: [][]int{
				[]int{1, 3, 5, 7},
				[]int{10, 11, 16, 20},
				[]int{23, 30, 34, 50},
			},
			target: 3,
			output: true,
		},
		tester{input: [][]int{[]int{}}, target: 1, output: false},
		tester{input: [][]int{}, target: 1, output: false},
	}
	for _, v := range tests {
		got := SearchMatrix(v.input, v.target)
		if got != v.output {
			t.Errorf("want:%v, got:%v", v.output, got)
		}
	}
}
