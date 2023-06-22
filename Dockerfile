FROM golang:1.18-alpine as builder

RUN apk add --no-cache g++ git
# override git so go can access private repos
# ONLY use this if there is 2nd stage container, otherwise ARG is visible in image history
ARG GH_TOKEN
RUN git config --global url."https://$GH_TOKEN:@github.com/".insteadOf "https://github.com/"

WORKDIR /im-executor
ADD . /im-executor
RUN go mod download
RUN go mod tidy

RUN go build -o /im-executor/bin/executor ./cmd/main

FROM alpine:latest
VOLUME /executor/env
WORKDIR /executor/env
COPY --from=builder /im-executor/bin/executor /usr/local/bin
CMD ["/bin/sh", "-c", "executor start --test --loglevel debug --home /executor/env 2> /executor/env/app.log"]