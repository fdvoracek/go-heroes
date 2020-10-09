FROM golang:1.15.2-alpine
WORKDIR /go/src/app
COPY . .
RUN go build ./main.go
EXPOSE 8080
CMD main