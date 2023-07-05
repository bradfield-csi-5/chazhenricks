package main

import "testing"

// func BenchmarkFirstArg(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		firstArg()
// 	}
// }



func BenchmarkStringJoinArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringJoinArgs()
	}
}
