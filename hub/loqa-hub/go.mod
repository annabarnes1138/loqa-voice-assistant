module github.com/annabarnes1138/loqa-voice-assistant/hub/loqa-hub

go 1.24

require (
	google.golang.org/grpc v1.65.0
	loqa-voice-assistant/proto/go v0.0.0-00010101000000-000000000000
)

replace loqa-voice-assistant/proto/go => ../../proto/go

require (
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240604185151-ef581f913117 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
