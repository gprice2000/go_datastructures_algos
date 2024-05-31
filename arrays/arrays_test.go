package main

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	t.Run("testing insert in front", func(t *testing.T) {
		got := Insert(5, 0, [5]int{6, 2, 1, 9})
		want := [5]int{5, 6, 2, 1, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %d. Want: %d", got, want)
		}
	})
}
