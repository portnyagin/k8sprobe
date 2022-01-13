package app

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"io"
	"net/http"
	"os"
)

type Config struct {
	Port   string `env:"PORT" envDefault:"8080"`
	Name   string `env:"NAME"`
	Friend string `env:"FRIEND"`
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Service %s. Say hello called\n", config.Name)
	w.Header().Set("Content-Type", "plain/text")
	io.WriteString(w, fmt.Sprintf("Hello, i am %s!\n", config.Name))
	w.WriteHeader(http.StatusOK)
}

func HandlerCallFriend(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Service %s. Call friend called\n", config.Name)
	resp, err := http.Get("http://" + config.Friend + ":8080/sayHello")
	if err != nil {
		fmt.Println(err)
	}
	b, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("Service %s. Call friend result: \"%s\"", config.Name, b)
}

var config Config

func Start() {
	fmt.Println(os.Args)
	if err := env.Parse(&config); err != nil {
		fmt.Println("can't load service config", err)
	}
	//config.Port = "8050"
	if config.Friend == "" {
		config.Friend = "Http://localhost"
	}
	fmt.Printf("Service %s started\n", config.Name)
	fmt.Println(config)

	http.HandleFunc("/sayHello", HelloHandler)
	http.HandleFunc("/callFriend", HandlerCallFriend)
	http.ListenAndServe(":"+config.Port, nil)
}
