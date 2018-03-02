package test

import (
    "testing"
)

func TestCoreApplication(t *testing.T) {

}

func BenchmarkCoreApplication(b *testing.B) {
    for index := 0; index < b.N; index++ {
        TestCoreApplication(new(testing.T))
    }
}
