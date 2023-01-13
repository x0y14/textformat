package textformat

import (
	"fmt"
	"testing"
)

func TestTokenize(t *testing.T) {
	in := "hello, #s(#d+#f)!"
	result, _ := format(in, []any{"john", 32, 1.1})
	fmt.Println(result)
}
