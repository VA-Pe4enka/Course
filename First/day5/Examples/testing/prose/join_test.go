package prose

import "testing"

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		testData{list: []string{}, want: ""},
		testData{list: []string{"яблоки"}, want: "яблоки"},
		testData{list: []string{"яблоки", "груши"}, want: "яблоки и груши"},
		testData{list: []string{"яблоки", "груши", "айва"}, want: "яблоки, груши и айва"},
	}
	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", test.list, got, test.want)
		}
	}
}
