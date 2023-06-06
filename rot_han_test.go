package rothan

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRot(t *testing.T) {
	source := "2023-06-06 苹果公司(Apple, AAPL)发布了Vision Pro头戴式设备，有可能成为世界第一。"
	assert.Equal(t, source, ROT(ROT(source)))
}

func TestSymbolRegex(t *testing.T) {
	matched, _ := regexp.MatchString("^\\s*[\\p{P}|\\p{Z}}|\\p{S}]+\\s*$", "!@#$%^&*()_+{}|:\"<>?`-=[]\\;',./~")
	assert.Equal(t, true, matched)
}
