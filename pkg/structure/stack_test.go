package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack{}
	assert.True(t, s.IsEmpty())

	s.Push("A")
	assert.False(t, s.IsEmpty())

	v, containsItems := s.Pop()
	assert.True(t, s.IsEmpty())
	assert.True(t, containsItems)
	assert.Equal(t, "A", v)

	v, containsItems = s.Pop()
	assert.True(t, s.IsEmpty())
	assert.False(t, containsItems)
	assert.Equal(t, "", v)
}

func TestStackWithInitialState(t *testing.T) {
	s := Stack([]string{"A", "B", "C"})
	assert.False(t, s.IsEmpty())

	s.Push("D")
	assert.False(t, s.IsEmpty())

	v, containsItems := s.Pop()
	assert.False(t, s.IsEmpty())
	assert.True(t, containsItems)
	assert.Equal(t, "D", v)

	v, containsItems = s.Pop()
	assert.False(t, s.IsEmpty())
	assert.True(t, containsItems)
	assert.Equal(t, "C", v)

	v, containsItems = s.Pop()
	assert.False(t, s.IsEmpty())
	assert.True(t, containsItems)
	assert.Equal(t, "B", v)

	v, containsItems = s.Pop()
	assert.True(t, s.IsEmpty())
	assert.True(t, containsItems)
	assert.Equal(t, "A", v)

	v, containsItems = s.Pop()
	assert.True(t, s.IsEmpty())
	assert.False(t, containsItems)
	assert.Equal(t, "", v)
}
