package app

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"io"
	"net/http"
	"os"
)

type Config struct {
	Port   string `env:"PORT" envDefault:":8080"`
	Name   string `env:"NAME"`
	Friend string `env:"FRIEND"`
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, fmt.Sprintf("Hello, i am %s!\n", config.Name))
}

func HandlerCallFriend(w http.ResponseWriter, req *http.Request) {
	http.Get(config.Friend + "/sayHello")
}

var config Config

func Start() {
	fmt.Println(os.Args)
	if err := env.Parse(&config); err != nil {
		fmt.Println("can't load service config", err)
	}
	fmt.Printf("Service %s started\n", config.Name)
	fmt.Println(config)

	http.HandleFunc("/sayHello", HelloHandler)
	http.HandleFunc("/callFriend", HandlerCallFriend)
	http.ListenAndServe(":"+config.Port, nil)
}
