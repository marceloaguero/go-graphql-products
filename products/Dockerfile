FROM golang:alpine AS builder
ENV GO111MODULE=on
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o server /build/server.go

FROM scratch
COPY --from=builder /build/server /app/
WORKDIR /app
CMD ["./server"]
