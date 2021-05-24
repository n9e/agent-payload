.PHONY: build

build:
	docker run --rm -it  \
		-w /go/src/github.com/n9e/agent-payload \
		-v ${PWD}:/go/src/github.com/n9e/agent-payload \
		ybbbbasdf/golang:v0.0.1 \
		bash -c "protoc --proto_path=/go/src:. --gogofast_out=/go/src --java_out=java proto/logs/agent_logs_payload.proto && \
		protoc --proto_path=/go/src:. --gogofast_out=/go/src proto/metrics/agent_payload.proto && \
		protoc --proto_path=/go/src:. --gogofaster_out=/go/src proto/process/agent.proto"

clean:
	rm -rf pb/agent_logs_payload.pb.go gogen/agent_payload.pb.go process/agent.pb.go java/com
