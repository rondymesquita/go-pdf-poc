FROM golang
RUN apt update && apt install libmagickwand-dev -y
RUN go get -u github.com/gen2brain/go-fitz
RUN mkdir -p /go/src/go-pdf
WORKDIR /go/src/go-pdf
