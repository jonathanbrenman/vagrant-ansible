package controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildPingController(t *testing.T) {
	factoryCtrl := NewCtrlFactory()
	pingCtrl := factoryCtrl.BuildPingController()
	assert.NotNil(t, pingCtrl)
}