FROM golang:1.26.5 AS build

WORKDIR /src

COPY go.mod go.sum secbutler.go ./
COPY cmd ./cmd
COPY pkg ./pkg
RUN go mod download

# Build a static Linux binary
RUN CGO_ENABLED=0 GOOS=linux  \
    go build -trimpath -ldflags="-s -w" -o /out/secbutler ./

FROM alpine:3.23

RUN apk add --no-cache ca-certificates git

# Copy the binary in
COPY --from=build /out/secbutler /secbutler

ENTRYPOINT ["/secbutler"]
