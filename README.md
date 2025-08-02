![loqa_social_preview_padded_1280x640](https://github.com/user-attachments/assets/99016e57-ace5-4140-a4f3-c49262f83253)
# Loqa – A Local-First Voice Assistant

**Loqa** (formerly *Rosey*) is a privacy-first, local-only voice assistant that operates entirely offline. It enables natural language interaction without relying on cloud infrastructure, commercial APIs, or internet connectivity—designed from the ground up to be private, modular, and extensible.

---

## 🧱 System Architecture

### 🖥️ Loqa Hub (Server)
A single backend node responsible for all heavy processing:

- **Hardware:** Mini PC (e.g. Beelink SER5) running a Go-based orchestrator and Python microservices
- **Responsibilities:**
  - Accepts wake word events from pucks
  - Routes audio to Python ASR service (Whisper)
  - Sends transcript to intent parser (LLM)
  - Executes chained commands
  - Sends response text to Python TTS service
  - Streams audio response back to puck

### 🎙️ Loqa Lite (Puck)
Multiple embedded clients placed in rooms throughout the home:

- **Hardware:** ESP32-S3-based puck + microphone array
- **Responsibilities:**
  - Local wake word detection (Edge Impulse)
  - Record audio on trigger and forward to Loqa Prime
  - Playback audio response from the server
  - Designed for near-room-scale voice capture

---

## 🔄 Communication Flow

```text
[ Loqa Lite ]
 └─> Wake word detected locally
 └─> Record request audio
 └─> Transmit to Loqa Prime via Wi-Fi

[ Loqa Prime ]
 └─> Convert speech to text
 └─> Parse intent and execute command chain
 └─> Generate audio response
 └─> Send audio back to Loqa Lite for playback
```

---

## 🧭 System Diagram

![Loqa System Diagram](docs/loqa-system-diagram.png)

---

## 🌱 Future Plans

- Support for **NSL (Neuro-Symbolic Learning)** to allow Loqa to learn new skills from voice interactions
- Multi-room context awareness
- Embedded user identification (voice fingerprinting)
- Offline skill scripting from natural language
- Optional local app for configuration and debugging

---

## 📦 Project Structure

```bash
loqa-voice-assistant/
├── docker-compose.yml     # Orchestrates Go + Python services
├── README.md
├── .env                   # Optional shared config
 
├── hub/                   # Backend server
│   ├── loqa-hub/          # Go orchestrator (wake, chaining)
│   │   ├── cmd/
│   │   └── internal/
│   ├── services/          # Python ASR, TTS, Intent services
│   ├── tests/
│   └── scripts/
│
├── puck/                  # Embedded client firmware
│   ├── firmware/          # ESP-IDF / PlatformIO
│   ├── hardware/          # Schematics, BOM
│   └── tests/
│
├── shared/                # Prompts, transcripts, test audio
└── tools/                 # CLI tools or puck simulators
```

---

## 🛠️ Tech Stack

- **Go** (Orchestration, HTTP routing)
- **Python** (ASR, intent parsing, TTS)
- **ESP32-S3** (PlatformIO + ESP-IDF)
- **Edge Impulse** (wake word inference)
- **Whisper.cpp / Coqui TTS** (offline STT/TTS)
- **Rodio / ALSA** (audio output)
- **HTTP / TCP** (communication)
- **Optional LLM module** (future experimentation)

---

## 📜 License

TBD — likely MIT or Apache 2.0

---

*Created by [Anna Barnes](https://www.linkedin.com/in/annabethbarnes) to bring voice assistance back to the edge—where it belongs.*

---

You can replace `docs/loqa-system-diagram.png` with your actual file path or diagram export later.
