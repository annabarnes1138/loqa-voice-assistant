package llm

import (
	"log"
)

// WhisperTranscriber handles speech-to-text using Whisper (stub version)
type WhisperTranscriber struct {
	modelPath string
}

// NewWhisperTranscriber creates a new Whisper transcriber
func NewWhisperTranscriber(modelPath string) (*WhisperTranscriber, error) {
	log.Printf("✅ Whisper stub loaded: %s", modelPath)
	return &WhisperTranscriber{
		modelPath: modelPath,
	}, nil
}

// Transcribe converts audio samples to text (stub version)
func (wt *WhisperTranscriber) Transcribe(audioData []float32, sampleRate int) (string, error) {
	// For now, return a stub transcription
	result := "Hello, this is a test transcription"
	log.Printf("🧠 Whisper stub transcription: \"%s\"", result)
	return result, nil
}

// Close cleans up the Whisper model
func (wt *WhisperTranscriber) Close() {
	log.Println("🧠 Whisper stub closed")
}