package main

import (
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/fatih/color"
	yaml "gopkg.in/yaml.v2"
)

type APIConf struct {
	ConsumerKey       string `yaml:"Consumer-Key"`
	ConsumerSecret    string `yaml:"Consumer-Secret"`
	AccessToken       string `yaml:"Access-Token"`
	AccessTokenSecret string `yaml:"Access-Token-Secret"`
}

var apiConf APIConf

func main() {
	err := SetConfig(&apiConf)
	if err != nil {
		panic(err)
	}
	api := anaconda.NewTwitterApiWithCredentials(
		apiConf.AccessToken,
		apiConf.AccessTokenSecret,
		apiConf.ConsumerKey,
		apiConf.ConsumerSecret)

	v := url.Values{}
	s := api.UserStream(v)

	// Tweetを出力
	for t := range s.C {
		switch v := t.(type) {
		case anaconda.Tweet:
			fmt.Printf("%-15s: %s\n", color.GreenString(v.User.ScreenName), v.Text)
		case anaconda.EventTweet:
			switch v.Event.Event {
			case "favorite":
				sn := color.RedString(v.Source.ScreenName)
				tw := v.TargetObject.Text
				fmt.Printf("Favorited by %-15s: %s\n", sn, tw)
			case "unfavorite":
				sn := color.CyanString(v.Source.ScreenName)
				tw := v.TargetObject.Text
				fmt.Printf("UnFavorited by %-15s: %s\n", sn, tw)
			}
		}
	}
}

func SetConfig(conf interface{}) error {
	// yamlを読み込む
	buf, err := ioutil.ReadFile("./conf.yml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		return err
	}
	return nil
}
