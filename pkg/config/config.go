// Package config provides functions for working with a centralized
// configuration.
package config

import (
	"errors"
	"fmt"

	"github.com/joho/godotenv"
)

var config map[string]string
var isInitialized bool = false

// IsInitialized returns a boolean indicating whether the config has been
// initialized or not.
func IsInitialized() bool {
	return isInitialized
}

// Get retrieves the value associated with the given key from the config.
// Returns an error if the config has not been initialized or if the key is not
// found.
func Get(key string) (string, error) {
	if !isInitialized {
		return "", errors.New("config is not initialized")
	}
	value, ok := config[key]
	if !ok {
		return "", fmt.Errorf("config key %s not found", key)
	}
	return value, nil
}

// GetOrDefault retrieves the value associated with the given key from the
// config, or returns the default value if the key is not found or an error
// occurs.
func GetOrDefault(key, dflt string) string {
	value, err := Get(key)
	if err != nil {
		return dflt
	}
	return value
}

// Init loads environment variables from a file located at envPath. Returns an
// error if there was an issue reading the file.
func Init(envPath string) error {
	fmt.Println(config)
	var err error
	config, err = godotenv.Read(envPath)
	if err == nil {
		isInitialized = true
	}
	return err
}
