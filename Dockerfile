# FROM node:16 as build-ui

# ARG dist=0.0

# COPY ["./ui/package.json", "./ui/tsconfig.json", "/ui/"]

# WORKDIR /ui
# RUN npm install

# COPY ui ./ui
# RUN npm run build

# ##

FROM golang:1.16-alpine as build-server

RUN apk add --no-cache git gcc musl-dev

ARG GITHUB_USER=bot
ARG GITHUB_TOKEN=token

RUN git config --global url."https://${GITHUB_USER}:${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"
ENV GOPRIVATE=github.com/bitmark-inc

WORKDIR $GOPATH/github.com/bitmark-inc/logistic-server

ADD go.mod .
RUN go mod download

ADD . .

RUN go build -o /go/bin/logistic-server ./services/logistic-server

##

FROM alpine:3.14
ARG dist=0.0

RUN apk add --no-cache curl

COPY --from=build-server /go/bin/logistic-server /logistic-server
COPY /ui/dist /usr/share/ui

ADD services/logistic-server/config.yaml.sample /.config/config.yaml
