package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	ImageUrl              string   `json:"image_url"`
	DefaultAccessToForums []uint16 `json:"default_access_to_forums"`
	ForumTopicsInPage     uint32   `json:"forum_topics_in_page"`
	ForumMessagesInPage   uint32   `json:"forum_messages_in_page"`
	BlogsInPage           uint16   `json:"blogs_in_page"`
	BlogTopicsInPage      uint16   `json:"blog_topics_in_page"`
}

func ParseConfig(pathToConfig string) Config {
	file, err := os.Open(pathToConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var config Config

	if err := json.Unmarshal(bytes, &config); err != nil {
		log.Fatal(err)
	}

	return config
}
