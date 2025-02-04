package audio

import (
	"log"
)

func StartVAD(pipeline *Pipeline) {
	
	// Starting Message for User
	log.Println("Starting Voice Activity Detection")

	// This is just a placeholder for the VAD logic. You'd need to interact with GStreamer
	// and listen for the microphone input stream to detect voice activity.
	go func() {
		// Mock VAD behavior
		for {
			// Your actual VAD code would go here to detect if the user is speaking.
			// If speaking, mute the test tone. If not speaking, play the test tone.
			// For now, we simulate that the user is not speaking.

			log.Println("Checking for speech...")
			// Example logic to stop the test tone when the user is speaking
			// p.gstPipeline.SetState(gstreamer.PLAYING) // this would enable the tone
		}
	}()
}
