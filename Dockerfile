FROM golang:1.25 AS build

WORKDIR /src

COPY go.mod go.sum secbutler.go ./
COPY cmd ./cmd
COPY pkg ./pkg
RUN go mod download

# Build a static Linux binary
RUN CGO_ENABLED=0 GOOS=linux  \
    go build -trimpath -ldflags="-s -w" -o /out/secbutler ./

FROM scratch

# Since scratch is empty, copy CA certs from the builder image.
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary in
COPY --from=build /out/secbutler /secbutler

ENTRYPOINT ["/secbutler"]