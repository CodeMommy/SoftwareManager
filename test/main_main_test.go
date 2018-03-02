package test

import (
    "testing"
)

func TestMainMain(t *testing.T) {
}

func BenchmarkMainMain(b *testing.B) {
    for index := 0; index < b.N; index++ {
        TestMainMain(new(testing.T))
    }
}
