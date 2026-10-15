package creational

import (
	"fmt"
	"sync"
)

// Singleton: use sync.Once for lazy, safe initialization.
// Prefer dependency injection over global singletons in application code.

type Config struct {
	DSN string
}

var (
	configInstance *Config
	configOnce    sync.Once
)

func GetConfig() *Config {
	configOnce.Do(func() {
		configInstance = &Config{DSN: "postgres://localhost/mydb"}
	})
	return configInstance
}

func ExampleSingleton() {
	a := GetConfig()
	b := GetConfig()
	fmt.Println(a == b)
	fmt.Println(a.DSN)
	// Output:
	// true
	// postgres://localhost/mydb
}
