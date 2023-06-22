package helper

import (
	// "benchmark"
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("world")

	if result != "hello world" {
		// panic("wrong expected")
		// t.Fail()
		t.Error("input tidak sesuai")
	}
	fmt.Println("Success")
}

func TestHelloName(t *testing.T) {
	result := HelloWorld("mca")
	if result != "hello mca" {
		// t.FailNow()
		// t.Fatal()
		t.Fail()
	}
	fmt.Println("Success")
}

func TestHelloAssert(t *testing.T) {
	result := HelloWorld("yes")
	assert.Equal(t, "hello yes", result)
	fmt.Println("Success")
}

func TestHelloRequire(t *testing.T) {
	result := HelloWorld("yes")
	require.Equal(t, "hello yes", result, "result harus sama")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("can't run on MAC")
	}
}

func TestMain(m *testing.M) {
	fmt.Println("before")
	m.Run()
	fmt.Println("after")
}

func TestSubTest(t *testing.T) {
	t.Run("M", func(t *testing.T) {
		result := HelloWorld("M")
		assert.Equal(t, "hello M", result)
	})

	t.Run("C", func(t *testing.T) {
		result := HelloWorld("C")
		require.Equal(t, "hello C", result)
	})
}

func TestTableTest(t *testing.T) {
	tests := []struct {
		name    string
		request string
		expect  string
	}{
		{
			name:    "HelloWorld(mca)",
			request: "mca",
			expect:  "hello mca",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expect, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("count")
	}
}

// go test
// go test -v
// go test -v -run=funcName
// go test -v -run=subTest/funcName
// go test -v ./...
