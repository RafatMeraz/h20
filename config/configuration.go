package config

import "time"

type config struct {
	TokenSecret         []byte `default:"key-of-jwt"`
	TokenValidationTime time.Duration
	Port                int `default:"8040"`
}

var AppConfiguration = config{
	TokenSecret:         []byte("key-of-jwt"),
	TokenValidationTime: time.Hour * 12,
	Port:                8040,
}
