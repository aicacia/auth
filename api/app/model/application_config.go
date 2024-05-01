package model

import (
	"github.com/aicacia/auth/api/app/repository"
	"github.com/mitchellh/mapstructure"
)

type ApplicationConfigST struct {
	Website string `json:"website"`
	SignUp  struct {
		Enabled  bool `json:"enabled"`
		Password struct {
			Enabled bool `json:"enabled"`
		} `json:"password"`
	} `json:"signup"`
} // @name ApplicationConfig

func ApplicationConfigFromApplicationConfigRows(rows []repository.ApplicationConfigRowST) (*ApplicationConfigST, error) {
	result := repository.ApplicationConfigRowsToMap(rows)
	var c ApplicationConfigST
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &c,
		TagName:  "json",
	})
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(result)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
