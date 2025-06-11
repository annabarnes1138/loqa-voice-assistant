![loqa_social_preview_padded_1280x640](https://github.com/user-attachments/assets/99016e57-ace5-4140-a4f3-c49262f83253)
# Loqa – A Local-First Voice Assistant

**Loqa** (formerly *Rosey*) is a privacy-first, local-only voice assistant that operates entirely offline. It enables natural language interaction without relying on cloud infrastructure, commercial APIs, or internet connectivity—designed from the ground up to be private, modular, and extensible.

---

## 🧱 System Architecture

### 🖥️ Loqa Prime (Server)
A single backend node responsible for all heavy processing:

- **Hardware:** Mini PC (e.g. Beelink SER5)
- **Responsibilities:**
  - Wake word registration and routing
  - Speech-to-text (STT)
  - Intent parsing and command chaining
  - Text-to-speech (TTS)
  - Audio response playback

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

## 🌱 Future Plans

- Support for **NSL (Neuro-Symbolic Learning)** to allow Loqa to learn new skills from voice interactions
- Multi-room context awareness
- Embedded user identification (voice fingerprinting)
- Offline skill scripting from natural language
- Optional local app for configuration and debugging

---

## 📦 Project Structure

```bash
loqa/
├── prime/              # Server software (Python-based)
├── lite/               # ESP32 puck firmware (ESP-IDF)
├── models/             # Edge Impulse wake word models
├── docs/               # Diagrams, planning notes
└── README.md
```

---

## 🛠️ Tech Stack

- **ESP32-S3** (PlatformIO + ESP-IDF)
- **Edge Impulse** (wake word inference)
- **Python 3** (server logic)
- **ALSA / PulseAudio** (audio output)
- **MQTT / HTTP** (communication)
- **Optional LLM module** (future experimentation)

---

## 📜 License

TBD — likely MIT or Apache 2.0

---

*Created by [Anna Barnes](https://www.linkedin.com/in/annabethbarnes) to bring voice assistance back to the edge—where it belongs.*

