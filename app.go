//go:generate catb --in_dir=backend/route/proto --out_dir=doc
package apicatcloud

import (
	"apicat-cloud/backend/config"
	"apicat-cloud/backend/model"
	"apicat-cloud/backend/model/sysconfig"
	"apicat-cloud/backend/module/cache"
	"apicat-cloud/backend/module/logger"
	"apicat-cloud/backend/module/mock"
	"apicat-cloud/backend/module/storage"
	"apicat-cloud/backend/route"
	"fmt"
	"log"
	"net/url"
)

type App struct{}

func NewApp(conf string) *App {
	if err := config.Load(conf); err != nil {
		log.Printf("load config %s faild, use default config. err: %s", conf, err)
	}
	config.LoadFromEnv()
	return &App{}
}

func (a *App) Run() error {
	if err := model.Init(); err != nil {
		return fmt.Errorf("init %v", err)
	}
	sysconfig.Init()

	if err := runMock(); err != nil {
		return err
	}

	if err := logger.Init(config.Get().App.Debug, config.Get().Log); err != nil {
		return fmt.Errorf("init %v", err)
	}

	if err := cache.Init(config.Get().Cache.ToMapInterface()); err != nil {
		return fmt.Errorf("init %v", err)
	}

	if err := storage.Init(config.Get().Storage.ToMapInterface()); err != nil {
		return fmt.Errorf("init %v", err)
	}

	if err := route.Init(); err != nil {
		return fmt.Errorf("init %v", err)
	}
	return nil
}

func runMock() error {
	cfg := config.Get().App
	if cfg.AppUrl == "" || cfg.MockServerBind == "" {
		return fmt.Errorf("init mock err, cfg: %v", cfg)
	}

	// 尝试解析URL
	u, err := url.Parse(cfg.AppUrl)
	if err != nil {
		return fmt.Errorf("init mock err, cfg: %v", cfg)
	}

	// 检查协议是否是http或https
	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("init mock err, cfg: %v", cfg)
	}

	go mock.Run(cfg.MockServerBind, mock.WithApiUrl(cfg.AppUrl))
	return nil
}
