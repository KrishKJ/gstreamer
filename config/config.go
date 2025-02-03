package config

import (
	"flag"
	"os"
	"strconv"
)

// Config holds configurable settings
type Config struct {
	TestToneFreq   int
	SilenceThreshold float64
	MicSource      string
}

// LoadConfig loads configuration from flags or environment variables
func LoadConfig() Config {
	freq := flag.Int("freq", 440, "Test tone frequency")
	threshold := flag.Float64("threshold", -50.0, "Silence detection threshold (dB)")
	mic := flag.String("mic", "pulsesrc", "Microphone source")

	flag.Parse()

	return Config{
		TestToneFreq:   getEnvInt("TEST_TONE_FREQ", *freq),
		SilenceThreshold: getEnvFloat("SILENCE_THRESHOLD", *threshold),
		MicSource:      getEnv("MIC_SOURCE", *mic),
	}
}

// Helper functions to read environment variables
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		if parsed, err := strconv.Atoi(value); err == nil {
			return parsed
		}
	}
	return fallback
}

func getEnvFloat(key string, fallback float64) float64 {
	if value, exists := os.LookupEnv(key); exists {
		if parsed, err := strconv.ParseFloat(value, 64); err == nil {
			return parsed
		}
	}
	return fallback
}