package test

import (
    "testing"
    "github.com/CodeMommy/SoftwareManager/source/core"
)

func TestCoreErrorCheck(t *testing.T) {
    core.ErrorCheck(nil)
}

func BenchmarkCoreErrorCheck(b *testing.B) {
    for index := 0; index < b.N; index++ {
        TestCoreErrorCheck(new(testing.T))
    }
}
