package audio

import (
	"github.com/notedit/gstreamer-go"
)

// StartVAD monitors the microphone volume and adjusts the test tone
func StartVAD(pipeline *AudioPipeline) {
	for {
		msg := pipeline.pipeline.GetMessage()
		if msg != nil && msg.GetType() == gstreamer.MessageElement {
			structure := msg.GetStructure()
			if structure != nil && structure.GetName() == "level" {
				rms, _ := structure.GetValueDouble("rms")
				if rms < -50.0 { // Silence threshold
					pipeline.testTone.Set("volume", 1.0)
				} else {
					pipeline.testTone.Set("volume", 0.0)
				}
			}
		}
	}
}