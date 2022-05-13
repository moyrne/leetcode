package no_944_delete_columns_to_make_sorted

import "testing"

func Test_minDeletionSize(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "实例1",
			args: args{
				strs: []string{"cba", "daf", "ghi"},
			},
			wantAns: 1,
		},
		{
			name: "实例2",
			args: args{
				strs: []string{"a", "b"},
			},
			wantAns: 0,
		},
		{
			name: "实例3",
			args: args{
				strs: []string{"zyx", "wvu", "tsr"},
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minDeletionSize(tt.args.strs); gotAns != tt.wantAns {
				t.Errorf("minDeletionSize() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
