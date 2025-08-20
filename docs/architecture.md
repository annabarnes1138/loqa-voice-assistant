# Loqa Voice Assistant Architecture

## Overview

Loqa is a voice-controlled smart home system with a distributed architecture featuring "pucks" (edge devices) that capture audio and a central "hub" that processes speech and executes commands.

## Components

### Hub Service (Go)
- **Location**: `hub/loqa-hub/`
- **Purpose**: Central processing service for voice commands
- **Features**:
  - gRPC audio streaming server
  - Whisper.cpp integration for speech recognition
  - Command parsing and execution
  - Response generation
  - NATS messaging for event-driven architecture

### Device Service (Go)
- **Location**: `hub/loqa-hub/cmd/device-service/`
- **Purpose**: Manages connected devices and executes commands
- **Features**:
  - Device discovery and management
  - Command execution interface
  - Status monitoring

### Test Puck (Go)
- **Location**: `puck/test-go/`
- **Purpose**: Test implementation of puck functionality
- **Features**:
  - Real-time audio capture with PortAudio
  - Voice Activity Detection (VAD)
  - Wake word detection for "Hey Loqa"
  - gRPC streaming to hub
  - Audio playback for responses

### Protocol Definition
- **Location**: `proto/`
- **Format**: Protocol Buffers (gRPC)
- **Features**:
  - Bidirectional audio streaming
  - Command/response messaging
  - Language-agnostic (Go bindings in `proto/go/`)

## Architecture Diagram

```
┌─────────────────┐    gRPC Stream     ┌─────────────────┐
│     Puck        │ ──────────────────► │      Hub        │
│                 │    Audio Chunks    │                 │
│ • Audio Capture │                    │ • Whisper STT   │
│ • VAD           │ ◄────────────────── │ • LLM Parser    │
│ • Wake Word     │   Commands/TTS     │ • Command Exec  │
│ • TTS Playback  │                    │ • NATS Msg      │
└─────────────────┘                    └─────────────────┘
                                                │
                                                │ NATS
                                                ▼
                                       ┌─────────────────┐
                                       │ Device Service  │
                                       │                 │
                                       │ • Device Mgmt   │
                                       │ • Command Exec  │
                                       │ • Status Mon    │
                                       └─────────────────┘
```

## Data Flow

1. **Wake Word**: Puck detects "Hey Loqa" wake word locally
2. **Audio Capture**: Puck captures audio with voice activity detection
3. **Streaming**: Audio chunks streamed to hub via gRPC
4. **Processing**: Hub transcribes audio using Whisper.cpp
5. **Parsing**: Commands extracted from transcription
6. **Messaging**: Commands sent via NATS to device service
7. **Execution**: Device service executes commands on smart home devices
8. **Response**: Hub sends command acknowledgment and TTS audio back to puck

## Configuration

The system supports environment-based configuration:

- **Server**: Host, ports, timeouts
- **Whisper**: Model path, language, processing parameters
- **NATS**: Connection URL, subjects, reconnection settings
- **Logging**: Level and format configuration

## Deployment

### Docker Compose
- Complete containerized deployment
- Services: hub, device-service, NATS
- Automated networking and volume management

### Local Development
- Build script with Whisper.cpp CGO configuration
- Environment variable configuration
- Hot-reload capabilities

## Future Implementation

- **Real Pucks**: Will be implemented in C++/Rust for production hardware
- **Current Go Puck**: Test implementation for validation and development
- **Cloud Integration**: Optional cloud processing for advanced NLP
- **Multi-tenant**: Support for multiple homes/users