// The utils package provides utility functions.
package utils

import (
	"github.com/NooFreeNames/MultiSocket/pkg/config"
	"github.com/NooFreeNames/MultiSocket/pkg/logger"
)

// GetOrLog retrieves the value of a configuration parameter specified by the
// given key. If an error occurs while retrieving the value, it logs a fatal
// error message using the logger package and exits the program.
func GetOrLog(key string) string {
	value, err := config.Get(key)
	if err != nil {
		logger.Fatal("Failed to get config value: %v", err)
	}
	return value
}
