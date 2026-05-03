package hello_test

import (
	"fmt"
	"testing"

	"github.com/qiukuip/go-hello"
)

func TestHello(t *testing.T) {
	data := "jack"
	expected := fmt.Sprintf("hello %s!\n", data)
	result := hello.Hello(data)
	if expected != result {
		t.Fatalf("expected result %s, but got %s\n", expected, result)
	}
}
