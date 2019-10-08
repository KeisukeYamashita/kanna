# Build Go Server Binary
FROM golang:1.13
LABEL MAINTAINER=KeisukeYamashita

ENV GO111MODULE on

WORKDIR /project

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go install -v \
            ./cmd/kanna

FROM alpine:latest
COPY --from=0 /go/bin/kanna /bin/kanna
ENTRYPOINT ["/bin/kanna"]