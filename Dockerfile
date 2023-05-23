# syntax=docker/dockerfile:1

FROM gcr.io/distroless/static-debian11

WORKDIR /

COPY ts-infi-authkey /ts-infi-authkey

USER nonroot:nonroot

ENTRYPOINT ["/ts-infi-authkey"]
