package fse

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestHistoryList(t *testing.T) {
	list := newHistoryList()
	list.append("3")
	list.append("2")
	list.append("1")
	assert.Equal(t, 3, list.size)

	current := list.root
	assert.Equal(t, "", current.data)
	assert.NotNil(t, current.prev)
	assert.Nil(t, current.next)

	for i := 1; i <= 2; i++ {
		current = current.prev
		assert.Equal(t, strconv.Itoa(i), current.data)
		assert.NotNil(t, current.prev)
		assert.NotNil(t, current.next)
	}

	current = current.prev
	assert.Equal(t, "3", current.data)
	assert.Nil(t, current.prev)
	assert.NotNil(t, current.next)

	assert.True(t, list.tailNode == current)

	list.append("3")
	assert.Equal(t, 3, list.size)

	assert.True(t, list.remove("2"))
	assert.Equal(t, 2, list.size)

	assert.True(t, list.remove("1"))
	assert.Equal(t, 1, list.size)

	assert.True(t, list.removeTail())
	assert.Equal(t, 0, list.size)

	assert.True(t, list.root == list.tailNode)
	assert.Nil(t, list.root.prev)
	assert.Nil(t, list.root.next)
}
