FROM golang:1.22-bookworm
WORKDIR /app
COPY . ./
RUN go mod download
EXPOSE 8080
CMD ["bash"]