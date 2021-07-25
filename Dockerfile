FROM golang:alpine

WORKDIR $GOPATH/src/github.com/filipeandrade6/vigia-go

COPY . .
RUN /bin/bash ./scripts/build.sh

WORKDIR $GOPATH
ENTRYPOINT ["vigia-go"]