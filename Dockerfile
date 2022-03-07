FROM golang:1.17-alpine

COPY . /simple-monitoring

WORKDIR /simple-monitoring

RUN go mod download
RUN GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

CMD [ "./.bin/app" ]
