FROM golang:1.21-alpine


RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o avitotech-testtask ./cmd/app/main.go

CMD [ "./avitotech-testtask" ]