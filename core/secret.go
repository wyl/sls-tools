package core

import (
	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
)

type AliSecret struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
}

func NewAliSecret() *AliSecret {
	auth := AliSecret{}
	path, _ := homedir.Expand("~/.config/sls-tools/config.toml")

	if _, err := toml.DecodeFile(path, &auth); err != nil {
		panic(err)
	}
	return &auth
}
