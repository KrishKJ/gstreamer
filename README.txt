# GStreamer Audio Stream Application

This application uses Golang and GStreamer bindings to create a simple pipeline that processes two audio streams: a test tone and a microphone input. The test tone is played only when the user is not speaking, using Voice Activity Detection (VAD).

## Prerequisites

- Go (version 1.18 or higher)
- GStreamer (with development libraries)

### Installation

For MacOS, you can install GStreamer using Homebrew:

For Dependencies
go mod init
go mod tidy

```bash
brew install gstreamer
brew install gstreamer-devel

Installing gstreamer Plugins:
brew install --cask gstreamer
brew install --cask gstreamer-plugins-good
brew install --cask gstreamer-plugins-bad
brew install --cask gstreamer-plugins-ugly
brew install --cask gstreamer-plugins-base


Install Go bindings for GStreamer with `github.com/go-gst/go-gst/gst`

Running Instructions:

1. Install GStreamer and its development packages (for Mac, follow the `brew` commands in the README).
2. Clone your repository and navigate into the project folder.
3. Input your microphone path in the pipeline string
4. Run the application with `go run main.go`.