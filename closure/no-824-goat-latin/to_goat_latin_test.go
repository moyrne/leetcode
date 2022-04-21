package no_824_goat_latin

import "testing"

func Test_toGoatLatin(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "I speak Goat Latin",
			args: args{
				sentence: "I speak Goat Latin",
			},
			want: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		{
			name: "The quick brown fox jumped over the lazy dog",
			args: args{
				sentence: "The quick brown fox jumped over the lazy dog",
			},
			want: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
		{
			name: "Each word consists of lowercase and uppercase letters only",
			args: args{
				sentence: "Each word consists of lowercase and uppercase letters only",
			},
			want: "Eachmaa ordwmaaa onsistscmaaaa ofmaaaaa owercaselmaaaaaa andmaaaaaaa uppercasemaaaaaaaa etterslmaaaaaaaaa onlymaaaaaaaaaa",
		},
		{
			name: "HZ sg L",
			args: args{
				sentence: "HZ sg L",
			},
			want: "ZHmaa gsmaaa Lmaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toGoatLatin(tt.args.sentence); got != tt.want {
				t.Errorf("toGoatLatin() = %v, want %v", got, tt.want)
			}
		})
	}
}
