package api

import "struct/config"

type Api struct {
	Key string
}

func NewApi() *Api {
	config := config.NewConfig()
	return &Api{
		Key: config.Key,
	}
}
