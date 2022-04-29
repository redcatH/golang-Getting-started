package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello,world"," ","!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t* testing.T,orig string){
		rev,error :=Reverse(orig)
		if error != nil {
			return
		}
		doubleRev,doubleError :=Reverse(rev)
		if doubleError != nil {
			return
		}
		t.Logf("Number of runes:%d,rev=%d,doubleRev=%d",utf8.RuneCountInString(orig),utf8.RuneCountInString(rev),utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before:%q,After:%q",orig,doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q",rev)
		}
	})
}
// func TestReverse(t *testing.T){
// 	testcases := []struct{
// 		in, want string
// 	}{
// 		{"hello,world", "dlrow,olleh"},
// 		{" "," "},
// 		{"!123456789", "987654321!"},
// 	}

// 	for _,tc := range testcases{
// 		rev := Reverse(tc.in)
// 		if rev!= tc.want{
// 			t.Errorf("Reverse(%q)==%q, want %q", tc.in, rev, tc.want)
// 		}
// 	}
// }