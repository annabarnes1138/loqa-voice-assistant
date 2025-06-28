flowchart TD
  A[Loqa Lite (Puck): Wake Word Detection] --> B[Audio Stream Sent to Loqa Prime]

  B --> C[STT: Whisper.cpp (Speech-to-Text)]
  C --> D[Intent + Argument Parsing<br/>(e.g., "Start coffee and dim lights")]

  D --> E[Neuro-Symbolic Reasoning Engine<br/>Scallop / µKanren / DeepProbLog]
  E --> F[Logic Processing<br/>- Resolve chaining<br/>- Apply routines<br/>- Handle conditions]

  F --> G[Action Execution Engine<br/>- MQTT / D-Bus / GPIO]
  G --> H[TTS: Local Speech Synthesis<br/>(e.g., Coqui TTS)]

  H --> I[Loqa Lite (Puck): Play Response Audio]

  %% Optional Future Layers
  E --> J[Skill Learning via Voice<br/>"Teach me what to do when I say..."]
  E --> K[User-Specific Logic & Preferences<br/>(Lightweight Knowledge Graph)]