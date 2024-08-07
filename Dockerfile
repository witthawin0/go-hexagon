FROM golang:latest

WORKDIR /app

COPY go.mod so.sum ./

RUN go mod download

COPY . . 

RUN go build -o main ./cmd .

EXPOSE 1234

CMD [ "./main" ]