package grpc

import (
	"io"
	"log"
	"os"

	pb "loqa-voice-assistant/proto/go"
	"github.com/annabarnes1138/loqa-voice-assistant/hub/loqa-hub/internal/llm"
)

// AudioService implements the gRPC AudioService
type AudioService struct {
	pb.UnimplementedAudioServiceServer
	transcriber    *llm.WhisperTranscriber
	commandParser  *llm.CommandParser
}

// NewAudioService creates a new audio service
func NewAudioService(modelPath string) (*AudioService, error) {
	transcriber, err := llm.NewWhisperTranscriber(modelPath)
	if err != nil {
		return nil, err
	}

	// Initialize command parser with Ollama
	ollamaURL := os.Getenv("OLLAMA_URL")
	if ollamaURL == "" {
		ollamaURL = "http://localhost:11434"
	}
	
	ollamaModel := os.Getenv("OLLAMA_MODEL")
	if ollamaModel == "" {
		ollamaModel = "llama3.2:3b"
	}

	commandParser := llm.NewCommandParser(ollamaURL, ollamaModel)
	
	// Test connection to Ollama (non-blocking)
	go func() {
		if err := commandParser.TestConnection(); err != nil {
			log.Printf("⚠️  Warning: Cannot connect to Ollama: %v", err)
			log.Println("🔄 Command parsing will use fallback logic")
		}
	}()

	return &AudioService{
		transcriber:   transcriber,
		commandParser: commandParser,
	}, nil
}

// StreamAudio handles bidirectional audio streaming from pucks
func (as *AudioService) StreamAudio(stream pb.AudioService_StreamAudioServer) error {
	log.Println("🎙️  Hub: New audio stream connected")

	for {
		// Receive audio chunk from puck
		chunk, err := stream.Recv()
		if err == io.EOF {
			log.Println("🎙️  Hub: Audio stream ended")
			return nil
		}
		if err != nil {
			log.Printf("❌ Error receiving audio chunk: %v", err)
			return err
		}

		log.Printf("📥 Hub: Received audio chunk from puck %s (%d bytes, wake_word: %v)", 
			chunk.PuckId, len(chunk.AudioData), chunk.IsWakeWord)

		// Process audio if it's end of speech
		if chunk.IsEndOfSpeech {
			log.Printf("🎯 Hub: Processing complete utterance from puck %s", chunk.PuckId)

			// Convert audio bytes back to float32
			audioData := bytesToFloat32Array(chunk.AudioData)
			
			// Transcribe audio using Whisper
			transcription, err := as.transcriber.Transcribe(audioData, int(chunk.SampleRate))
			if err != nil {
				log.Printf("❌ Error transcribing audio: %v", err)
				continue
			}
			
			wakeWordStatus := ""
			if chunk.IsWakeWord {
				wakeWordStatus = " [wake word detected]"
			}
			log.Printf("📝 Processing audio (%d samples)%s -> \"%s\"", len(audioData), wakeWordStatus, transcription)

			if transcription == "" {
				log.Printf("🔇 No speech detected in audio from puck %s", chunk.PuckId)
				continue
			}

			log.Printf("📝 Transcribed: \"%s\"", transcription)

			// Parse command using LLM
			command, err := as.commandParser.ParseCommand(transcription)
			if err != nil {
				log.Printf("❌ Error parsing command: %v", err)
				command = &llm.Command{
					Intent:     "unknown",
					Entities:   make(map[string]string),
					Confidence: 0.0,
					Response:   "I'm having trouble understanding you right now.",
				}
			}

			commandStr := command.Intent
			responseText := command.Response
			
			log.Printf("🧠 Parsed command - Intent: %s, Entities: %v, Confidence: %.2f", 
				command.Intent, command.Entities, command.Confidence)

			// Send response back to puck
			response := &pb.AudioResponse{
				RequestId:     chunk.PuckId, // Use puck ID as request ID
				Transcription: transcription,
				Command:       commandStr,
				ResponseText:  responseText,
				Success:       true,
			}

			if err := stream.Send(response); err != nil {
				log.Printf("❌ Error sending response: %v", err)
				return err
			}

			log.Printf("📤 Hub: Sent response to puck %s - Command: %s", 
				chunk.PuckId, commandStr)
		}
	}
}

// Helper function to convert bytes back to float32 array
func bytesToFloat32Array(data []byte) []float32 {
	// Convert 16-bit PCM bytes to float32 samples
	samples := make([]float32, len(data)/2)
	for i := 0; i < len(samples); i++ {
		// Reconstruct int16 from bytes (little-endian)
		val := int16(data[i*2]) | int16(data[i*2+1])<<8
		// Convert to float32 [-1,1]
		samples[i] = float32(val) / 32767.0
	}
	return samples
}