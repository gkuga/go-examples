package main

import "log"

import "github.com/kelseyhightower/envconfig"

type Env struct {
	HogeId int `required:"true" split_words:"true"`
	FooID  int `required:"true" split_words:"true"`
}

func main() {
	var goenv Env
	if err := envconfig.Process("", &goenv); err != nil {
		log.Printf("[ERROR] Failed to process env: %s", err)
	}
}
