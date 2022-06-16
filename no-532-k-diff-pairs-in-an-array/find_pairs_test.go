package no_532_k_diff_pairs_in_an_array

import "testing"

func Test_findPairs(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例 1",
			args: args{
				nums: []int{3, 1, 4, 1, 5},
				k:    2,
			},
			want: 2,
		},
		{
			name: "示例 2",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    1,
			},
			want: 4,
		},
		{
			name: "示例 3",
			args: args{
				nums: []int{1, 3, 1, 5, 4},
				k:    0,
			},
			want: 1,
		},
		{
			name: "示例 55",
			args: args{
				nums: []int{2, 9, 0, 8, 9, 6, 5, 9, 8, 1, 9, 6, 9, 2, 8, 8, 7, 5, 7, 8, 8, 3, 7, 4, 1, 1, 6, 2, 9, 9, 3, 9, 2, 4, 6, 5, 6, 5, 1, 5, 9, 9, 8, 1, 4, 3, 2, 8, 5, 3, 5, 4, 5, 7, 0, 0, 7, 6, 4, 7, 2, 4, 9, 3, 6, 6, 5, 0, 0, 0, 8, 9, 9, 6, 5, 4, 6, 2, 1, 3, 2, 5, 0, 1, 4, 2, 6, 9, 5, 4, 9, 6, 0, 8, 3, 8, 0, 0, 2, 1},
				k:    1,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPairs(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
