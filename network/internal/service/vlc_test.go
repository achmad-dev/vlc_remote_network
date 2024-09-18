package service_test

import (
	"network/internal/service"
	"network/pkg/keytap"
	"network/pkg/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVLCService(t *testing.T) {
	keyb, err := keytap.NewKeyBonding()
	assert.NoError(t, err)
	logger := logger.InitLogger()
	vlcService := service.NewVlcService(logger, &keyb)
	err = vlcService.StartOrStop()
	assert.NoError(t, err)
	err = vlcService.VolumeUp()
	assert.NoError(t, err)
	err = vlcService.VolumeDown()
	assert.NoError(t, err)
	err = vlcService.Next()
	assert.NoError(t, err)
	err = vlcService.Previous()
	assert.NoError(t, err)
}
