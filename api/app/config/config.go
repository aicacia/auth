package config

import (
	"encoding/json"
	"log/slog"
	"strings"

	"github.com/aicacia/auth/api/app/repository"
	atomic_value "github.com/aicacia/go-atomic-value"
	"github.com/lib/pq"
	"github.com/mitchellh/mapstructure"
)

var config atomic_value.AtomicValue[*ConfigST]

func Get() *ConfigST {
	return config.Load()
}

type ConfigST struct {
	Host      string `json:"host"`
	Port      int16  `json:"port"`
	URL       string `json:"url"`
	Dashboard struct {
		Enabled bool `json:"enabled"`
	} `json:"dashboard"`
	OpenAPI struct {
		Enabled bool `json:"enabled"`
	} `json:"openapi"`
}

func InitConfig() error {
	configs, err := repository.GetConfigs()
	if err != nil {
		return err
	}
	configJSON := make(map[string]interface{})
	for _, config := range configs {
		var value interface{}
		err := json.Unmarshal([]byte(config.Value), &value)
		if err != nil {
			slog.Error("invalid json", "key", config.Key, "error", err)
			continue
		}
		setKeyValue(configJSON, config.Key, value)
	}
	var c ConfigST
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &c,
		TagName:  "json",
	})
	if err != nil {
		return err
	}
	err = decoder.Decode(configJSON)
	if err != nil {
		return err
	}
	config.Store(&c)

	listener, err := repository.CreateListener("configs_channel")
	if err != nil {
		return err
	}
	go configListener(listener)
	return nil
}

func CloseConfigListener() error {
	configListenerSignal <- true
	return nil
}

var configListenerSignal = make(chan bool, 1)

func configListener(listener *pq.Listener) {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("Recovered in configListener", "error", err)
		}
	}()
	for {
		select {
		case <-configListenerSignal:
			err := listener.Close()
			if err != nil {
				slog.Error("error closing config listener", "error", err)
				return
			} else {
				slog.Error("config listener closed\n")
			}
			return
		case notification := <-listener.Notify:
			var extra ExtraST
			err := json.Unmarshal([]byte(notification.Extra), &extra)
			if err != nil {
				slog.Error("invalid json", "extra", notification.Extra, "error", err)
			} else {
				updateConfig(extra.Key, extra.Value)
			}
		}
	}
}

type ExtraST struct {
	Table      string      `json:"table"`
	Key        string      `json:"key"`
	Value      interface{} `json:"value"`
	ActionType string      `json:"action_type"`
}

func updateConfig(key string, value interface{}) error {
	configJSON := make(map[string]interface{})
	setKeyValue(configJSON, key, value)

	c := *Get()
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &c,
		TagName:  "json",
	})
	if err != nil {
		return err
	}
	err = decoder.Decode(configJSON)
	if err != nil {
		return err
	}
	config.Store(&c)

	return nil
}

func setKeyValue(parent map[string]interface{}, key string, value interface{}) {
	entry := parent
	path := strings.Split(key, ".")
	for _, key := range path[:len(path)-1] {
		subEntry, ok := entry[key]
		if !ok {
			subEntry = make(map[string]interface{})
			entry[key] = subEntry
		}
		entry = subEntry.(map[string]interface{})
	}
	k := path[len(path)-1]
	entry[k] = value
	slog.Debug("set config", "key", key, "value", value)
}
