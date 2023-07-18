package config

import "time"

type config struct {
	TokenSecret         string `default:"key-of-jwt"`
	TokenValidationTime time.Duration
	Port                int `default:"8040"`
}

var AppConfiguration = config{
	TokenSecret:         "key-of-jwt",
	TokenValidationTime: time.Hour * 12,
	Port:                8040,
}
