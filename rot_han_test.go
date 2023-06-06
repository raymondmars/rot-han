package rothan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRot(t *testing.T) {
	assert.Equal(t, "苹果公司(Apple, AAPL)发布了Vision Pro头戴式设备", ROT(ROT("苹果公司(Apple, AAPL)发布了Vision Pro头戴式设备")))
}
