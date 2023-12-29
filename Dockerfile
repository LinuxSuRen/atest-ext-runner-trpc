FROM golang:1.20 AS builder

ARG VERSION
ARG GOPROXY
WORKDIR /workspace
COPY cmd/ cmd/
COPY pkg/ pkg/
COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go
COPY README.md README.md

RUN GOPROXY=${GOPROXY} go mod download
RUN GOPROXY=${GOPROXY} CGO_ENABLED=0 go build -ldflags "-w -s" -o atest-runner-trpc .

FROM alpine:3.12

LABEL org.opencontainers.image.source=https://github.com/LinuxSuRen/atest-ext-runner-trpc
LABEL org.opencontainers.image.description="tRPC Runner Extension of the API Testing."

COPY --from=builder /workspace/atest-runner-trpc /usr/local/bin/atest-runner-trpc

CMD [ "atest-runner-trpc" ]
