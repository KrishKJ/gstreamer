package audio

import (
	"fmt"
	// "github.com/notedit/gst-go"
	// "github.com/tolexo/gst/config"
	// "/gst-go"
	"github.com/ziutek/gst"
	// "gst/config"
	"github.com/KrishKJ/gstreamer/config"

)

// AudioPipeline wraps the gst pipeline
type AudioPipeline struct {
	pipeline *gst.Pipeline
	testTone *gst.Element
	level    *gst.Element
}

// NewPipeline initializes the gst pipeline
func NewPipeline(cfg config.Config) (*AudioPipeline, error) {
	// Define gst pipeline
	pipelineStr := fmt.Sprintf(`
		%s name=micsrc ! audioconvert ! audioresample ! level interval=50000000 name=level
		audiotestsrc wave=sine freq=%d ! audioconvert ! audioresample name=testtone
		adder name=mixer ! audioconvert ! autoaudiosink
	`, cfg.MicSource, cfg.TestToneFreq)

	pipeline, err := gst.NewPipeline(pipelineStr)
	if err != nil {
		return nil, fmt.Errorf("failed to create pipeline: %w", err)
	}

	return &AudioPipeline{
		pipeline: pipeline,
		testTone: pipeline.GetElement("testtone"),
		level:    pipeline.GetElement("level"),
	}, nil
}

// Start plays the pipeline
func (p *AudioPipeline) Start() error {
	return p.pipeline.SetState(gst.StatePlaying)
}

// Stop stops the pipeline
func (p *AudioPipeline) Stop() error {
	return p.pipeline.SetState(gst.StateNull)
}
