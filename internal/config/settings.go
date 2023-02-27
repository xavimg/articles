package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var Settings *SettingsRoot

type SettingsRoot struct {
	Api          Api          `yml:"api"`
	Mongo        Mongo        `yml:"mongo"`
	FeedProvider FeedProvider `yml:"feedprovider"`
}

type Api struct {
	Port int `yml:"port"`
}

type Mongo struct {
	URL      string `yml:"url"`
	Database string `yml:"database"`
}

type FeedProvider struct {
	HullCity string `yml:"hullcity"`
}

func LoadSettings() error {
	f, err := os.Open(os.Getenv("CONFIG_FILE"))
	if err != nil {
		logrus.Error(err)
		return err
	}

	return yaml.NewDecoder(f).Decode(&Settings)
}
