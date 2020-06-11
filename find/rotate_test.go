package find

import (
	"algorithm/find/utils"
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestRotate(t *testing.T) {
	suite.Run(t, new(RotateSuite))
}

type RotateSuite struct {
	suite.Suite
	ctx context.Context
}

func (r *RotateSuite) SetupTest() {
	r.ctx = context.Background()
}

func (r *RotateSuite) TestRotateArrayFindSmall() {
	expectVal := int64(1)
	for i := 0; i < 10; i++ {
		arr := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		rotateArr := utils.RotateN(arr, i)
		small := RotateArrayFindSmall(rotateArr)
		r.Equal(expectVal, small)
	}
}
