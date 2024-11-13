FROM golang:1.22
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go .
COPY internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -o /webclinicalrecords
EXPOSE 8080
CMD [ "/webclinicalrecords" ]
