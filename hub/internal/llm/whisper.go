package llm

import (
	"log"
)

// WhisperTranscriber handles speech-to-text using Whisper
type WhisperTranscriber struct {
	modelPath string
}

// NewWhisperTranscriber creates a new Whisper transcriber
func NewWhisperTranscriber(modelPath string) (*WhisperTranscriber, error) {
	log.Printf("✅ Whisper transcriber initialized: %s", modelPath)
	return &WhisperTranscriber{
		modelPath: modelPath,
	}, nil
}

// Transcribe converts audio samples to text
func (wt *WhisperTranscriber) Transcribe(audioData []float32, sampleRate int) (string, error) {
	// TODO: Implement actual Whisper.cpp integration
	result := "Hello, this is a test transcription"
	log.Printf("🧠 Whisper transcription: \"%s\"", result)
	return result, nil
}

// Close cleans up the Whisper model
func (wt *WhisperTranscriber) Close() {
	log.Println("🧠 Whisper transcriber closed")
}