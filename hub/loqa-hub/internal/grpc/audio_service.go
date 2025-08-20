package grpc

import (
	"io"
	"log"

	pb "loqa-voice-assistant/proto/go"
)

// AudioService implements the gRPC AudioService
type AudioService struct {
	pb.UnimplementedAudioServiceServer
}

// NewAudioService creates a new audio service
func NewAudioService() (*AudioService, error) {
	return &AudioService{}, nil
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

			// For now, send a simple echo response
			transcription := "Hello from Loqa Hub"
			commandStr := "greeting"
			responseText := "Hello! I'm your local voice assistant."

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