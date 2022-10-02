package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type config struct {
	Version string `yaml:"version"`
	IP      string `yaml:"ip"`
	Port    string `yaml:"port"`
}

var configFile *string
var c config

func getVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"), "getVersion()")
	fmt.Fprintf(w, c.Version)
}

func getConfig() error {
	if configFile == nil || *configFile == "" {
		panic("invalid config file")
	}
	file, err := os.Open(*configFile)
	if err != nil {
		return err
	}

	yamlDec := yaml.NewDecoder(file)
	return yamlDec.Decode(&c)
}

func main() {
	configFile = flag.String("conf", "", "")
	flag.Parse()

	if err := getConfig(); err != nil {
		panic(err)
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getVersion)

	fmt.Printf("listening on %s:%s\n", c.IP, c.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", c.IP, c.Port), router))
}
