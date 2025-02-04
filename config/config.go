package config

import (
	"log"
	"os"
)

type Config struct {
	TestTonePath string
	// MicInputPath is not needed anymore, as microphone input is captured directly by GStreamer.
}

func LoadConfig() *Config {
	cfg := &Config{
		// local file path for the test tone
		TestTonePath: "test-tone.wav",  
	}

	// Load configurations from environment variables or files here
	if testTonePath, exists := os.LookupEnv("TEST_TONE_PATH"); exists {
		cfg.TestTonePath = testTonePath
	}

	log.Printf("Using test tone from %s", cfg.TestTonePath)

	// No need to log or configure MicInputPath directly, as GStreamer will handle microphone input

	return cfg
}
