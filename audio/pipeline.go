package audio

import (
	"fmt"
	"log"
	"os"

	"github.com/KrishKJ/gstreamer/config"
	gst "github.com/go-gst/go-gst/gst"
)

type Pipeline struct {
	gstPipeline *gst.Pipeline
}

// Init initializes GStreamer
func Init() error {
	// Initialize GStreamer with os.Args (it doesn't return an error, it panics on failure)
	gst.Init(&os.Args)
	return nil // No need to check for error since gst.Init() doesn't return one
}

func NewPipeline(cfg *config.Config) (*Pipeline, error) {
	// Initialize GStreamer
	if err := Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize GStreamer: %v", err)
	}

	// Build pipeline based on configuration
	// pipelineStr := fmt.Sprintf("playbin uri=file://%s", cfg.TestTonePath)

	// Pipeline for test tone and microphone (use your actual microphone input logic)
	// Example for macOS, replace this with the correct source
	pipelineStr := fmt.Sprintf("avfsrc ! audioconvert ! audioresample ! autoaudiosink")  // Replace with correct microphone capture

	// Build pipeline based on configuration
	if cfg.TestTonePath != "" {
		pipelineStr = fmt.Sprintf("filesrc location=%s ! decodebin ! audioconvert ! audioresample ! autoaudiosink", cfg.TestTonePath)
	}

	// Create the pipeline
	pipeline, err := gst.NewPipeline(pipelineStr)
	if err != nil {
		return nil, fmt.Errorf("failed to create pipeline: %v", err)
	}

	return &Pipeline{
		gstPipeline: pipeline,
	}, nil
}

// Start plays the pipeline
func (p *Pipeline) Start() error {
	return p.gstPipeline.SetState(gst.StatePlaying)
}

func (p *Pipeline) Stop() {
	// Set pipeline state to Null to stop it
	p.gstPipeline.SetState(gst.StateNull)
	log.Println("Pipeline stopped")
}
