FROM golang:alpine
WORKDIR /app
RUN apk update && apk add fio
COPY . .
RUN go build -o fio-tests .
CMD ["tail", "-f", "/dev/null"]
