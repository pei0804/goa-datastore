package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertIdIntoInt64(t *testing.T) {
	assert := assert.New(t)
	int64ID, err := ConvertIdIntoInt64("5649050225344512")
	assert.NoError(err)
	assert.Equal(int64(5649050225344512), int64ID)
}
