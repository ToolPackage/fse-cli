package fse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPath(t *testing.T) {
	path := newPath("C:\\a\\b\\c")
	assert.Equal(t, "C:\\a\\b\\c", path.String())
	path.StepBack()
	assert.Equal(t, "C:\\a\\b", path.String())
	path.StepForward("xxx")
	assert.Equal(t, "C:\\a\\b\\xxx", path.String())
	path.StepBack()
	path.StepBack()
	path.StepBack()
	assert.Equal(t, "C:", path.String())
	path.StepBack()
	assert.Equal(t, "C:", path.String())
}
