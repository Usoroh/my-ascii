# Base image to pull from
FROM golang:latest

# Set the working directory
WORKDIR /

COPY . /

#build a go file
RUN go build -o main .

EXPOSE 9090

CMD ["./main"]

