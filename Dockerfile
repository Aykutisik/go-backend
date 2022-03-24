FROM golang:1.17-alpine
WORKDIR /app
COPY . .
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait
RUN go mod download
RUN go build -o app
EXPOSE 8086
CMD /wait && ./app
