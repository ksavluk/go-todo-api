# cache modules
FROM golang:1.17 as modules

ADD go.mod go.sum /m/
RUN cd /m && go mod download

# build binary
FROM golang:1.17 as builder
COPY . /src
COPY --from=modules /go/pkg /go/pkg
WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux make build

# package final image
FROM alpine:latest
COPY --from=builder /src/bin/* /app/
EXPOSE 8080

RUN addgroup -S app && adduser -S app -G app

RUN apk update \
    && apk upgrade \
    && apk --no-cache add --update -t ca-certificates \
    && rm -rf /tmp/* /var/cache/apk/*

WORKDIR /app
USER app

CMD /app/todo