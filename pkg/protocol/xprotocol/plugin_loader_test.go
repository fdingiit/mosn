package xprotocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	noneExistDir = "/abcdefg/hijklmn/opq/rst"
	notPluginDir = "./"
)

func TestInitCodec_NoneExistDir(t *testing.T) {
	if !assert.NotNil(t, InitCodec(noneExistDir)) {
		t.FailNow()
	}
}

func TestInitCodec_NoPlugin(t *testing.T) {
	if !assert.NotNil(t, InitCodec(notPluginDir)) {
		t.FailNow()
	}
}

func TestInitCodec_Load(t *testing.T) {
	codecPath := "/Users/dingfei/Go/src/github.com/fdingiit/xprotocol_thrift"
	if !assert.Nil(t, InitCodec(codecPath)) {
		t.FailNow()
	}
}
