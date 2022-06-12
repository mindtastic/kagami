FROM golang:1.18-alpine AS builder

WORKDIR /build

COPY main.go .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -ldflags="-s -w" -o kagami main.go

FROM scratch

COPY --from=builder ["/build/kagami", "/"]

# Command to run when starting the container.
ENTRYPOINT ["/kagami"]
