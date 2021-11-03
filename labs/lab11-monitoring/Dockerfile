FROM golang:1.17 as build
ADD . /go/src/github.com/replicatedhq/kots-field-labs/labs/lab11-monitoring
RUN go install /go/src/github.com/replicatedhq/kots-field-labs/labs/lab11-monitoring/cmd/flaky-app

FROM debian:stretch-slim
COPY --from=build /go/bin/flaky-app /go/bin/flaky-app
EXPOSE 3000
CMD ["/go/bin/flaky-app"]
