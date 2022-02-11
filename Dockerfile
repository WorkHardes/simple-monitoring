FROM golang:1.17-alpine

COPY . /simple-monitoring

RUN go mod download
RUN GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

WORKDIR /simple-monitoring

CMD [ "./app" ]
