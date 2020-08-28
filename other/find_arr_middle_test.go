package other

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindArrMiddle(t *testing.T) {
	arr := []int64{1, 2, 3, 1, 4, 1, 1, 1, 5, 1, 5, 6}
	result, err := FindMoreThanHalf(arr)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), result)
}

func TestFindKSmall(t *testing.T) {
	arr := []int64{1, 2, 3, 1, 4, 1, 1, 1, 5, 1, 5, 6}
	result := []int64{1, 1, 1, 1, 1, 1, 2, 3, 4, 5, 5, 6}
	for idx := range result {
		k, err := FindKSmall(arr, idx)
		assert.Equal(t, nil, err)
		assert.Equal(t, result[idx], int64(k))
	}

}
