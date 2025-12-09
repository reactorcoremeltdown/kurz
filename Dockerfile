FROM golang:alpine AS builder
RUN apk add git make build-base
COPY src/kurz /srv/kurz
WORKDIR /srv/kurz
ENV GOBIN=/usr/local/bin
ENV CGO_ENABLED=1
RUN go get && go build -o kurz

FROM alpine:3.21.3
COPY --from=builder /srv/kurz /srv/kurz
CMD [ "/srv/kurz/kurz" ]
