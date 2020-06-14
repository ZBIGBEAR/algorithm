package other

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestRotate(t *testing.T) {
	suite.Run(t, new(OddEvenNumberSuite))
}

type OddEvenNumberSuite struct {
	suite.Suite
	ctx context.Context
}

func (r *OddEvenNumberSuite) SetupTest() {
	r.ctx = context.Background()
}

func (r *OddEvenNumberSuite) TestRotateArrayFindSmall() {
	for i := 0; i < 10; i++ {
		r.Equal(IsEvenNumber(i), IsEvenNumber1(i))
	}
}
