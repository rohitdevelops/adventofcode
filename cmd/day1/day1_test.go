// Objective: Find set of n (2 or 3) entries from the input entries in the file
// that sum to 2020, and print their product.
// Puzzle Link : https://adventofcode.com/2020/day/1

package main

import (
	"testing"
)

func Test_findtwoentries(t *testing.T) {
	type args struct {
		numbers   []int
		sumverify int
	}
	type want struct {
		ent1, ent2, prod int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{"test1", args{[]int{1721, 979, 366, 299, 675, 1456}, 2020}, want{1721, 299, 514579}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, b, c := findtwoentries(tt.args.numbers, tt.args.sumverify)
			if (a != tt.want.ent1) && (b != tt.want.ent2) && (c != tt.want.prod) {
				t.Errorf("got %d, %d and %d, want %d, %d and %d", a, b, c, tt.want.ent1, tt.want.ent2, tt.want.prod)
			}
		})
	}
}

func Test_findthreeentries(t *testing.T) {
	type args struct {
		numbers   []int
		sumverify int
	}
	type want struct {
		ent1, ent2, ent3, prod int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{"test1", args{[]int{1721, 979, 366, 299, 675, 1456}, 2020}, want{979, 366, 675, 241861950}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, b, c, d := findthreeentries(tt.args.numbers, tt.args.sumverify)
			if (a != tt.want.ent1) && (b != tt.want.ent2) && (c != tt.want.ent3) && (d != tt.want.prod) {
				t.Errorf("got %d, %d, %d and %d, want %d, %d, %d and %d", a, b, c, d, tt.want.ent1, tt.want.ent2, tt.want.ent3, tt.want.prod)
			}
		})
	}
}
