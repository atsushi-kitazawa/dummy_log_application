package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"
)

type settings struct {
	Timestamp     string   `yaml:"timestamp"`
	LogFormat     string   `yaml:"logFormat"`
	StartLogParam []string `yaml:"startLogParam"`
	LogParam      []string `yaml:"logParam"`
	EndLogParam   []string `yaml:"endLogParam"`
	File          string   `yaml:"file"`
}

var (
	s    settings
	cnt  int
	mode string
)

func main() {
	flag.IntVar(&cnt, "cnt", 10, "log count")
	flag.StringVar(&mode, "mode", "append", "output mode")
	flag.Parse()

	load()
	doMain()
}

func doMain() {
	if !exists(exePath() + "/" + s.File) {
		create(exePath() + "/" + s.File)
	} else {
		if "clear" == mode {
			remove(exePath() + "/" + s.File)
			create(exePath() + "/" + s.File)
		}
	}

	f, err := os.OpenFile(exePath()+"/"+s.File, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	msgs := []interface{}{time.Now().Format(s.Timestamp), s.StartLogParam[0], s.StartLogParam[1], s.StartLogParam[2]}
	fmt.Fprintf(f, s.LogFormat, msgs...)

	for i := 1; i <= cnt; i++ {
		msgs = []interface{}{time.Now().Format(s.Timestamp), s.LogParam[0], s.LogParam[1], s.LogParam[2] + strconv.Itoa(i)}
		fmt.Fprintf(f, s.LogFormat, msgs...)
	}

	msgs = []interface{}{time.Now().Format(s.Timestamp), s.EndLogParam[0], s.EndLogParam[1], s.EndLogParam[2]}
	fmt.Fprintf(f, s.LogFormat, msgs...)
}

func load() {
	f, err := os.Open(exePath() + "/settings.yml")
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
}

func exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func create(name string) {
	_, err := os.Create(s.File)
	if err != nil {
		log.Fatalln(err)
	}
}

func remove(name string) {
	err := os.Remove(s.File)
	if err != nil {
		log.Fatalln(err)
	}
}

func exePath() string {
	exe, _ := os.Executable()
	path := filepath.Dir(exe)
	return path
}
