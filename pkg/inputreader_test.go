package inputreader

import (
	"reflect"
	"testing"
)

func TestGetLinesInIntArray(t *testing.T) {
	want := [4]int{20, 55, 87, 94}
	ans := GetLinesInIntArray("TestData/TestIntData.txt")
	if reflect.DeepEqual(want, ans) {
		t.Errorf("Got: %v, Want: %v", ans, want)
	}
}

func TestGetLinesInStringArray(t *testing.T) {
	want := [3]string{"Hello", "Golang", "World !"}
	ans := GetLinesInStringArray("TestData/TestStringData.txt")
	if reflect.DeepEqual(want, ans) {
		t.Errorf("Got: %v, Want: %v", ans, want)
	}
}
