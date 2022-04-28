package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type settings struct {
	Timestamp string   `yaml:"timestamp"`
	LogFormat string   `yaml:"logFormat"`
	Param     []string `yaml:"param"`
}

var s settings

func main() {
	load()
	doMain()
}

func doMain() {
	timestamp := time.Now().Format("2006-01-02 03:04:05")
	msgs := []interface{}{timestamp, "INFO", "start", "messages."}
	fmt.Fprintf(os.Stdout, logTemplate(), msgs...)
}

func load() {
	f, err := os.Open("settings.yml")
	if err != nil {
		log.Fatalln(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	s = settings{}
	err = yaml.Unmarshal(data, &s)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(s)
}

func logTemplate() string {
	return "%s %s 【%s】 %s\n"
}
