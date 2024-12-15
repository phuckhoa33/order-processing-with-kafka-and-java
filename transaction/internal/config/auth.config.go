package config

import "os"

type AuthConfig struct {
	AccessTokenSecret     string
	AccessTokenExpiredIn  string
	RefreshTokenSecret    string
	RefreshTokenExpiredIn string
}

func LoadAuthConfig() AuthConfig {
	return AuthConfig{
		AccessTokenSecret:     os.Getenv("ACCESS_TOKEN_SECRET"),
		AccessTokenExpiredIn:  os.Getenv("ACCESS_TOKEN_EXPIRES_IN"),
		RefreshTokenSecret:    os.Getenv("REFRESH_TOKEN_SECRET"),
		RefreshTokenExpiredIn: os.Getenv("REFRESH_TOKEN_EXPIRES_IN"),
	}
}
