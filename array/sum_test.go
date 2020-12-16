package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	
	checkSums := func (t *testing.T, got, expected int, numbers []int)  {
		if got != expected {
			t.Errorf("given '%v', want '%d' but got '%v'", expected, got, numbers)
		}
	}

	t.Run("collection of 5 numbers", func (t *testing.T)  {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expected := 15
		checkSums(t, got, expected, numbers)
		// if got != expected {
		// 	t.Errorf("given '%v', want '%d' but got '%v'", expected, got, numbers)
		// }
	})
	t.Run("collection of any sizes", func (t *testing.T)  {
		numbers := []int{1,2,3}
		got := Sum(numbers)
		expected := 6
		checkSums(t, got, expected, numbers)
		// if got != expected {
		// 	t.Errorf("given '%v', want '%d' but got '%v'", expected, got, numbers)
		// }
	})
}

func TestSumAll(t *testing.T)  {
	t.Run("sum all array", func (t *testing.T)  {
		got := SumAll([]int{1,2}, []int{0,9})
		want := []int{3, 9}
		//不能对切片使用等号运算符
		// if got != want {
		// 	t.Errorf("want '%v' but got '%v'", want, got)
		// }
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("sum all array", func (t *testing.T)  {
		got := SumAll([]int{1,2,3}, []int{0,9,3})
		want := []int{6, 12}
		//不能对切片使用等号运算符
		// if got != want {
		// 	t.Errorf("want '%v' but got '%v'", want, got)
		// }
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	
	checkSumTail := func (t *testing.T, got, want []int)  {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum all numbers except first one", func (t *testing.T)  {
		got := SumAllTails([]int{1,2}, []int{0,9})
		want := []int{2, 9}
		checkSumTail(t, got, want)
		// if !reflect.DeepEqual(got, want) {
		// 		t.Errorf("got %v want %v", got, want)
		// }
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want :=[]int{0, 9}
		checkSumTail(t, got, want)
		// if !reflect.DeepEqual(got, want) {
		// 		t.Errorf("got %v want %v", got, want)
		// }
	})
}
