package test

import (
    "testing"
    "github.com/CodeMommy/SoftwareManager/source/core"
)

func TestCoreHttpGet(t *testing.T) {
    content, err := core.HttpGet("https://github.com")
    if err != nil {
        t.Errorf("Error: %v", err)
    }
    if content == "" {
        t.Error("Empty Content")
    }
}

func BenchmarkCoreHttpGet(b *testing.B) {
    for index := 0; index < b.N; index++ {
        TestCoreHttpGet(new(testing.T))
    }
}
