FROM golang:1.15.1-alpine3.12 AS GO_BUILD
RUN apk update && apk add build-base
COPY server /server
WORKDIR /server
RUN go build -o /go/bin/server

FROM alpine:3.12.0
COPY --from=GO_BUILD /go/bin/server ./
CMD ./server
