package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type configuration struct {
	Users  []string `json:"users"`
	Groups []string `json:"groups"`
}

type safeConfiguratin struct {
	configuration
	sync.Mutex
}

var (
	conf safeConfiguratin
	once sync.Once
)

func Getconfig() configuration {
	once.Do(func() {
		readConfig()
		go watchConfig()
	})

	conf.Lock()
	defer conf.Unlock()

	return conf.configuration
}

func watchConfig() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for {
		<-ticker.C
		readConfig()
	}
}

func readConfig() {

	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("err", err)
	}
	defer file.Close()

	conf.Lock()
	defer conf.Unlock()

	decoder := json.NewDecoder(file)
	// conf := configuration{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("Прочитан конфиг. Users: %v, Groups: %v", conf.Users, conf.Groups)
}
