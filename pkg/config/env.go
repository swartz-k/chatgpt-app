package config

import (
	"encoding/json"
	"fmt"
	"github.com/swartz-k/chatgpt-app/pkg/log"
	"os"
	"sync"
)

var (
	cfgOnce = sync.Once{}
	cfg     *Config
)

func GetConfig(path *string) *Config {
	cfgOnce.Do(func() {
		if cfg == nil {
			cfg = newDefaultConfig()
		}
		if path != nil && *path != "" {
			log.V(10).Info("load config from %s", *path)
			content, err := os.ReadFile(*path)
			if err != nil {
				log.V(3).Info("cannot open config %s, err %+v", *path, err)
				return
			}
			err = json.Unmarshal(content, cfg)
			if err != nil {
				log.V(3).Info("cannot unmarshal config %s, err %+v", *path, err)
				return
			}
			log.V(10).Info("success use config file at %s", *path)
		}
	})
	if cfg.Session == "" {
		panic(fmt.Errorf("no session config"))
	}
	return cfg
}
