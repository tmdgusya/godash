package arrayhelper_test

import (
	arrayhelper "godash/src/array_helper"
	"testing"
)

func TestEqualForList(t *testing.T) {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{1, 2, 3, 4, 5}

	if arrayhelper.Equal(arr1[:], arr2[:]) == false {
		t.Errorf("Not Equals Array!")
	}
}
