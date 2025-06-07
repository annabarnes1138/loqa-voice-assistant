# Loqa – A Local-First Voice Assistant

**Loqa** is a privacy-first, fully offline-capable voice assistant built around modular smart pucks. It enables responsive, voice-driven interaction **without the cloud**, designed from the ground up for **local-only processing**, **room-level awareness**, and **command chaining**.

## 🎯 Goals

- Fully local processing (no cloud dependency)
- Wake word and command chaining supported per room
- Expandable with low-cost voice puck modules
- Optional LLM-powered assistant (offline via ESP32)

---

## 🧱 System Architecture

### 🧠 Loqa Prime
> The brains of the operation. You can deploy one or multiple in your home.

- **Device:** M5Stack CoreS3 SE  
- **Modules:** Built-in mic/speaker, M5Stack LLM Module (UART)
- **Role:**  
  - Acts as a command processor
  - Parses requests from Loqa Lite units
  - Generates natural language responses locally
  - Handles high-level command logic
  - Plays output via onboard speaker

### 🌿 Loqa Lite
> Lightweight room modules that trigger wake word detection and relay commands.

- **Device:** M5Stack AtomS3R  
- **Base:** Atomic Echo Base (adds mic + speaker)
- **Role:**  
  - Performs **local wake word detection** via Edge Impulse  
  - On trigger, sends command audio/text to a Loqa Prime  
  - Optionally plays back Loqa Prime’s audio response (via speaker)
  - Identifies which room/user issued command

---

## 🔄 Communication Flow

```text
[ Loqa Lite ]
 └─> Wake word detected
 └─> Capture voice / record request
 └─> Send to Loqa Prime (Wi-Fi / UART / TBD)

[ Loqa Prime ]
 └─> Receive request
 └─> Run LLM parsing / command chaining
 └─> Respond with natural language text or audio

[ Loqa Lite ]
 └─> (Optional) Play response via speaker
```

---

## 📦 Project Structure

```bash
loqa/
├── prime/              # CoreS3 SE firmware + LLM interface
├── lite/               # AtomS3R firmware (wake word, comms)
├── models/             # Edge Impulse wake word model(s)
├── docs/               # Architecture, planning notes
└── README.md
```

---

## 🔧 Tech Stack

- **ESP-IDF + PlatformIO** (both Loqa Prime and Loqa Lite)
- **Edge Impulse** (wake word model on Lite)
- **UART + Wi-Fi** (inter-device communication)
- **LLM Module** (text generation offline, UART to Prime)
- **Audio I/O** via M5Stack onboard components

---

## 🚧 Current Status

- ✅ Hardware selected: CoreS3 SE, AtomS3R, LLM module, Echo Base
- ⚙️ Wake word model training in progress (Edge Impulse)
- 🔌 Inter-puck communication prototype in development
- 🧠 LLM module integration and response playback pending

---

## 📜 License

TBD — likely MIT or Apache 2.0

---

*Created by [Anna Barnes](https://www.linkedin.com/in/annabethbarnes) to prove that powerful, respectful AI doesn’t need the cloud—or a billion-dollar company behind it.*
