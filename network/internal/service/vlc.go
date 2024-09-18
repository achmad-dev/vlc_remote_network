package service

import (
	"network/pkg/keytap"
	"time"

	"github.com/sirupsen/logrus"
)

type VlcService interface {
	StartOrStop() error
	VolumeUp() error
	VolumeDown() error
	Next() error
	Previous() error
}

type baseVlcService struct {
	logger     *logrus.Logger
	keyBonding *keytap.KeyBonding
}

func NewVlcService(logger *logrus.Logger, keyBonding *keytap.KeyBonding) VlcService {
	return &baseVlcService{
		keyBonding: keyBonding,
		logger:     logger,
	}
}

func (b *baseVlcService) StartOrStop() error {
	time.Sleep(time.Millisecond * 500)
	b.keyBonding.SetKeys(keytap.VK_SPACE)
	err := b.keyBonding.Launching()
	if err != nil {
		return err
	}
	return nil
}

func (b *baseVlcService) VolumeUp() error {
	time.Sleep(time.Millisecond * 500)
	b.keyBonding.SetKeys(keytap.VK_UP)
	err := b.keyBonding.Launching()
	if err != nil {
		return err
	}
	return nil
}

func (b *baseVlcService) VolumeDown() error {
	time.Sleep(time.Millisecond * 500)
	b.keyBonding.SetKeys(keytap.VK_DOWN)
	err := b.keyBonding.Launching()
	if err != nil {
		return err
	}
	return nil
}

func (b *baseVlcService) Next() error {
	time.Sleep(time.Millisecond * 500)
	b.keyBonding.SetKeys(keytap.VK_N)
	err := b.keyBonding.Launching()
	if err != nil {
		return err
	}
	return nil
}

func (b *baseVlcService) Previous() error {
	time.Sleep(time.Millisecond * 500)
	b.keyBonding.SetKeys(keytap.VK_P)
	err := b.keyBonding.Launching()
	if err != nil {
		return err
	}
	return nil
}
