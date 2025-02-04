package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/KrishKJ/gstreamer/audio"
	"github.com/KrishKJ/gstreamer/config"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize and start GStreamer pipeline
	pipeline, err := audio.NewPipeline(cfg)
	if err != nil {
		fmt.Println("Failed to initialize pipeline:", err)
		return
	}

	// Start voice activity detection (VAD) in a separate goroutine
	go audio.StartVAD(pipeline)

	// Handle graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	// Cleanup
	pipeline.Stop()
}
