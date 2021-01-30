FROM golang:buster as builder

RUN update-ca-certificates

// Simple wait app to keep container running if no app exist
RUN echo "package main;import(\"math\";\"time\");func main(){<-time.After(time.Duration(math.MaxInt64))}" > wait.go
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -o /wait wait.go

# Build app
RUN CGO_ENABLED=0 GOOS=linux go build .

# Run-time image
FROM gcr.io/distroless/base

COPY --from=builder /wait /wait
COPY --from=builder /go/bin/ /usr/local/bin/

ENTRYPOINT ["/wait"]