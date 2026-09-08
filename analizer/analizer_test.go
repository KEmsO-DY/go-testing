package textstat

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
	type Case struct {
		s    string
		want map[string]int
	}
	testCases := []Case{
		{"HELLO HELLo HELlo HEllo,Hello", map[string]int{"hello": 5}},
		{" ...,. ,, ,..,  . ,", map[string]int{}},
		{" .HELLO..,. War,, ,..,  . ,", map[string]int{"hello": 1, "war": 1}},
	}
	for _, c := range testCases {
		c := c
		name := fmt.Sprintf("Testing:%v", c.s)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := WordCount(c.s)
			assert.Equal(t, got, c.want)
		})
	}
}
