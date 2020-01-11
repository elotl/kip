package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapError(t *testing.T) {
	e := fmt.Errorf("1")
	we1 := WrapError(e, "we1")
	assert.Equal(t, "we1: 1", we1.Error())
	innerError := we1.(WrappedError).Cause()
	assert.Equal(t, "1", innerError.Error())

	// make sure that having no message doesn't make strange formatting
	we2 := WrapError(e, "")
	assert.Equal(t, "1", we2.Error())
}
