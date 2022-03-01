FROM golang:1.17.6
WORKDIR /code
ADD . /code/
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
