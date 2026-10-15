package creational

import (
	"bytes"
	"fmt"
	"sync"
)

// Object Pool: sync.Pool is the standard library answer.
// Objects may be GC'd between uses — do not store state across Get/Put.

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func ProcessData(data []byte) string {
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	buf.Write(data)
	return buf.String()
}

func ExampleObjectPool() {
	result := ProcessData([]byte("hello pool"))
	fmt.Println(result)
	// Output:
	// hello pool
}
