package first_bad_version_278

import "testing"

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n          int
		badVersion []bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "18",
			args: args{
				n:          1,
				badVersion: []bool{true},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			badVersion = tt.args.badVersion
			if got := firstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
