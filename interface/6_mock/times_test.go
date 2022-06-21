package times

import (
	"testing"

	"github.com/GreenStoneWei/golang-for-heros/interface/6_mock/number"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTimeTen(t *testing.T) {
	t.Run("with mock", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)

		mg := NewMockgenerator(ctrl)
		mg.EXPECT().RandomNumber(number.ODD).Return(9)

		result := TimesTen(number.ODD, mg)

		assert.Equal(t, result, 90)

	})
}
