FROM golang:1.22
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY main.go  ./
COPY internal ./internal
COPY templates ./templates
COPY records ./records

RUN go build -o ./webclinicalrecords
EXPOSE 8080
CMD [ "./webclinicalrecords" ]
