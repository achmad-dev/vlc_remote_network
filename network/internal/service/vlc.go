package service

import (
	"network/pkg/keytap"
	"time"

	"github.com/sirupsen/logrus"
)

// VlcService is a interface for vlcservice for controlling VLC media player
type VlcService interface {
	StartOrStop() error
	VolumeUp() error
	VolumeDown() error
	Next() error
	Previous() error
}

// baseVlcService is a base implementation of vlc service
type baseVlcService struct {
	logger     *logrus.Logger
	keyBonding *keytap.KeyBonding
}

// NewVlcService creates a new vlc service
func NewVlcService(logger *logrus.Logger, keyBonding *keytap.KeyBonding) VlcService {
	return &baseVlcService{
		keyBonding: keyBonding,
		logger:     logger,
	}
}

// StartOrStop starts or stops VLC media player
func (b *baseVlcService) StartOrStop() error {
	time.Sleep(time.Millisecond * 500)
	b.keyBonding.SetKeys(keytap.VK_SPACE)
	err := b.keyBonding.Launching()
	if err != nil {
		return err
	}
	return nil
}

// VolumeUp increases the volume of VLC media player
func (b *baseVlcService) VolumeUp() error {
	time.Sleep(time.Millisecond * 500)
	b.keyBonding.SetKeys(keytap.VK_UP)
	err := b.keyBonding.Launching()
	if err != nil {
		return err
	}
	return nil
}

// VolumeDown decreases the volume of VLC media player
func (b *baseVlcService) VolumeDown() error {
	time.Sleep(time.Millisecond * 500)
	b.keyBonding.SetKeys(keytap.VK_DOWN)
	err := b.keyBonding.Launching()
	if err != nil {
		return err
	}
	return nil
}

// Next plays the next track in VLC media player
func (b *baseVlcService) Next() error {
	time.Sleep(time.Millisecond * 500)
	b.keyBonding.SetKeys(keytap.VK_N)
	err := b.keyBonding.Launching()
	if err != nil {
		return err
	}
	return nil
}

// Previous plays the previous track in VLC media player
func (b *baseVlcService) Previous() error {
	time.Sleep(time.Millisecond * 500)
	b.keyBonding.SetKeys(keytap.VK_P)
	err := b.keyBonding.Launching()
	if err != nil {
		return err
	}
	return nil
}
