FROM golang:1.20-alpine
WORKDIR /go/src/github.com/vspcsi/inventory/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o binary cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/vspcsi/inventory/binary ./
ENV GIN_MODE release
CMD ["./binary"]
