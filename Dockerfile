FROM public.ecr.aws/docker/library/golang:1.17-alpine3.15 as BUILDER
RUN apk add --no-cache git
WORKDIR ${GOPATH}/src/go-http-sleep
COPY go.mod main.go .
RUN CGO_ENABLED=0 go build -o /go/bin/go-http-sleep

FROM scratch
COPY --from=BUILDER /go/bin/go-http-sleep /opt/
EXPOSE 8080
ENTRYPOINT ["/opt/go-http-sleep"]
