package main

import "testing"

func TestOnlyId(t *testing.T) {
	t.Log(onlyID())
}
func BenchmarkOnlyId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		onlyID()
	}
}
func TestOnlyId1(t *testing.T) {
	t.Log(onlyID1())
}
func BenchmarkOnlyId1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		onlyID1()
	}
}
