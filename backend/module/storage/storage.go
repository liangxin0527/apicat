package storage

import (
	"apicat-cloud/backend/module/storage/cloudflare"
	"apicat-cloud/backend/module/storage/common"
	"apicat-cloud/backend/module/storage/local"
	"apicat-cloud/backend/module/storage/qiniu"
	"errors"
	"log/slog"
)

const (
	LOCAL      = "disk"
	CLOUDFLARE = "cloudflare"
	QINIU      = "qiniu"
)

func NewStorage(cfg map[string]interface{}) (common.Storage, error) {
	slog.Debug("storage.NewStorage", "cfg", cfg)
	if cfg == nil {
		return nil, errors.New("storage config is nil")
	}

	switch cfg["Driver"] {
	case CLOUDFLARE:
		return cloudflare.NewR2(cfg["Cloudflare"].(map[string]interface{}))
	case QINIU:
		return qiniu.NewQiniu(cfg["Qiniu"].(map[string]interface{}))
	case LOCAL:
		return local.NewDisk(cfg["LocalDisk"].(map[string]interface{}))
	default:
		return nil, errors.New("storage driver not found")
	}
}

func Init(cfg map[string]interface{}) error {
	if cfg == nil {
		return errors.New("storage config is nil")
	}

	switch cfg["Driver"] {
	case CLOUDFLARE:
		if s, err := cloudflare.NewR2(cfg["Cloudflare"].(map[string]interface{})); err != nil {
			return err
		} else {
			return s.Check()
		}
	case QINIU:
		if s, err := qiniu.NewQiniu(cfg["Qiniu"].(map[string]interface{})); err != nil {
			return err
		} else {
			return s.Check()
		}
	case LOCAL:
		if s, err := local.NewDisk(cfg["LocalDisk"].(map[string]interface{})); err != nil {
			return err
		} else {
			return s.Check()
		}
	default:
		return errors.New("storage driver not found")
	}
}
