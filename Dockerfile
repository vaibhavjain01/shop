FROM golang:alpine3.10

ARG URL
ENV ENV_URL $URL

USER root
RUN apk update && apk add --no-cache git
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN echo "DOCKER RUNNING"
#RUN update-ca-certificates

WORKDIR /go/src/server/
COPY . ./
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix -o metrics

#ENTRYPOINT /go/src/server/server --jia-url "${ENV_URL}"
ENTRYPOINT /go/src/server/server