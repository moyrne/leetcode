package no_704_binary_search

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums=[-1,0,3,5,9,12],target=9,output:4",
			args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9},
			want: 4,
		},
		{
			name: "nums=[-1,0,3,5,9,12],target=2,output:-1",
			args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2},
			want: -1,
		},
		{
			name: "nums=[5],target=5,output:0",
			args: args{nums: []int{5}, target: 5},
			want: 0,
		},
		{
			name: "nums=[2,5],target=5,output:1",
			args: args{nums: []int{2, 5}, target: 5},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
