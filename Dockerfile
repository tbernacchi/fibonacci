##STEP1
FROM golang:1.22 as builder

RUN mkdir -p /app

WORKDIR /app/

RUN go mod init fibonacci && go mod tidy

ADD . /app

#Be attention to your archicteture
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./server/main.go

EXPOSE 8000

CMD ["./main" ]

#STEP2
FROM alpine:latest

#Set the working directory
WORKDIR /app

#Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Add a user for running the application
RUN adduser -S -D -H -h /app appuser && chown -R appuser: /app

# Switch to the new user
USER appuser

# Command to run the executable
CMD ["./main"]
