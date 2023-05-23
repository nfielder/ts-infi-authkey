# syntax=docker/dockerfile:1

FROM golang:1.20-alpine as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /ts-infi-authkey cmd/ts-infi-authkey/main.go

# Lean image
FROM gcr.io/distroless/static-debian11

WORKDIR /

COPY --from=builder /ts-infi-authkey /ts-infi-authkey

USER nonroot:nonroot

ENTRYPOINT ["/ts-infi-authkey"]
