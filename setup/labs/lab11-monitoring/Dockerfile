FROM golang:1.17 as build
WORKDIR /go/src/app
COPY . .
RUN go install ./cmd/flaky-app

FROM debian:stretch-slim
COPY --from=build /go/bin/flaky-app /go/bin/flaky-app
EXPOSE 3000
CMD ["/go/bin/flaky-app"]
