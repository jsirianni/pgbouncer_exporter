FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o /pgbouncer_exporter

FROM busybox:latest
LABEL maintainer="Joe Sirianni <joe@observiq.com>"

COPY --from=builder /pgbouncer_exporter /bin/pgbouncer_exporter
COPY LICENSE                                /LICENSE

USER       nobody
ENTRYPOINT ["/bin/pgbouncer_exporter"]
EXPOSE     9127
