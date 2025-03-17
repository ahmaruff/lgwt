package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

    t.Run("Colection of 5 int", func(t *testing.T) {
        numbers := []int{1,2,3,4,5}

        got := Sum(numbers)
        want := 15

        if got != want {
            t.Errorf("got %d want %d, given %v", got, want, numbers)
        }
    })

    t.Run("Colection of ANY int", func(t *testing.T) {
        numbers := []int{1,2,3}

        got := Sum(numbers)
        want := 6

        if got != want {
            t.Errorf("got %d want %d, given %v", got, want, numbers)
        }
    })
}


func BenchmarkSum(t *testing.B) {
     
    numbers := []int{1,2,3,4,5}

    for i :=0; i<t.N; i++ {
        Sum(numbers)
    }
}


func TestSumAll(t *testing.T) {
    got := SumAll([]int{1,2,3}, []int{4}, []int{5,5})
    want := []int{6,4,10}

    if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}


func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}
