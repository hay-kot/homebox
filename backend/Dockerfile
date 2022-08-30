# Build API
FROM golang:alpine AS builder
RUN apk add --no-cache git build-base
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/api -v ./app/api/*.go
RUN go build -o /go/bin/manage -v ./app/cli/*.go


# Production Stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
COPY ./config.template.yml /app/config.yml
COPY --from=builder /go/bin/api /app
COPY --from=builder /go/bin/manage /bin

RUN chmod +x /app/api
RUN chmod +x /bin/manage

LABEL Name=gowebtemplate Version=0.0.1
EXPOSE 7745
WORKDIR /app
CMD [ "./api" ]
