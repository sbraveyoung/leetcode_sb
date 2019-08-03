package mapSum

import (
	"testing"
)

type testST struct {
	op          string
	keyOrPrefix string
	val         int
}

func test(ts []testST) bool {
	ms := Constructor()
	for _, t := range ts {
		switch t.op {
		case "i":
			ms.Insert(t.keyOrPrefix, t.val)
		case "s":
			val := ms.Sum(t.keyOrPrefix)
			if val != t.val {
				return false
			}
		default:
		}
	}
	return true
}

func TestConstructor(t *testing.T) {
	type args []testST
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				testST{"i", "apple", 3},
				testST{"s", "ap", 3},
				testST{"i", "app", 2},
				testST{"s", "ap", 5},
			},
		},
		{
			name: "second",
			args: args{
				testST{"i", "apple", 3},
				testST{"s", "ap", 3},
				testST{"i", "app", 2},
				testST{"s", "app", 5},
				testST{"s", "appl", 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test(tt.args); got != true {
				t.Errorf("Constructor() = %v", got)
			}
		})
	}
}
