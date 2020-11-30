package reorganizeString

import "testing"

func Test_reorganizeString(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		//{
		//	name: "first",
		//	args: args{
		//		S: "aab",
		//	},
		//	want: "aba",
		//},
		//{
		//	name: "second",
		//	args: args{
		//		S: "aaab",
		//	},
		//	want: "",
		//},
		//{
		//	name: "third",
		//	args: args{
		//		S: "aabbcc",
		//	},
		//	want: "abcabc",
		//},
		//{
		//	name: "fourth",
		//	args: args{
		//		S: "abbbccd",
		//	},
		//	want: "bcbabcd",
		//},
		//{
		//	name: "fifth",
		//	args: args{
		//		S: "abbbbcd",
		//	},
		//	want: "babcbdb",
		//},
		//{
		//	name: "sixth",
		//	args: args{
		//		S: "abbbbbcd",
		//	},
		//	want: "",
		//},
		//{
		//	name: "seventh",
		//	args: args{
		//		S: "kkkkzrkatkwpkkkktrq",
		//	},
		//	want: "krkrktktkakpkqkwkzk",
		//},
		{
			name: "eights",
			args: args{
				S: "gpneqthatplqrofqgwwfmhzxjddhyupnluzkkysofgqawjyrwhfgdpkhiqgkpupgdeonipvptkfqluytogoljiaexrnxckeofqojltdjuujcnjdjohqbrzzzznymyrbbcjjmacdqyhpwtcmmlpjbqictcvjgswqyqcjcribfmyajsodsqicwallszoqkxjsoskxxstdeavavnqnrjelsxxlermaxmlgqaaeuvneovumneazaegtlztlxhihpqbajjwjujyorhldxxbdocklrklgvnoubegjrfrscigsemporrjkiyncugkksedfpuiqzbmwdaagqlxivxawccavcrtelscbewrqaxvhknxpyzdzjuhvoizxkcxuxllbkyyygtqdngpffvdvtivnbnlsurzroxyxcevsojbhjhujqxenhlvlgzcsibcxwomfpyevumljanfpjpyhsqxxnaewknpnuhpeffdvtyjqvvyzjeoctivqwann",
			},
			want: "jpjpjpjpjpjyjyjyjyjyjyjyjyjyjyjyjyjyjyjyjyjyjyjyjyjgjgjgjgjgqgqgqgqgqgqgqgqgqgqgqgqgqgqrqrqrqrqrqrqrqrqrqrqrqrxrxrxrxrxrxrxuxuxuxuxuxuxuxuxuxuxuxuxuxuxuxuxuxuxdldldldldldldldldldldldldldldldldlhlhlhlhlhlhlhlhnhnhnhnhnhnhnhnhnhnknknknknknknknknknknknknknkakakakasasasasasasasasasasasasasasasasasazazczczczczczczczczczczczczczczcfcfcfcfcfcfcfcfefefefefefefefeieieieieieieieieieieieieieieiotototototototototototototototobobobobobobvbvbvbvbvbvbvbvbvwvwvwvwvwvwvwvwvwvwvwvwvwpwpmpmpmpmpmpmpmpmpmpmpmpmpmp",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorganizeString(tt.args.S); !right(got, tt.want) {
				t.Errorf("reorganizeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func right(got, want string) bool {
	if len(got) != len(want) {
		return false
	}
	var array Slice
	for i := 0; i < 26; i++ {
		array = append(array, Dict{
			char:  'a' + i,
			count: 0,
		})
	}

	for _, c := range got {
		array[c-'a'].count++
	}
	for _, c := range want {
		array[c-'a'].count--
	}
	for _, a := range array {
		if a.count != 0 {
			return false
		}
	}

	for i := 0; i < len(got)-1; i++ {
		if got[i] == got[i+1] {
			return false
		}
	}
	return true
}
