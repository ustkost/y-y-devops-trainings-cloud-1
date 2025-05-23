FROM golang:1.21 as build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY tpl ./tpl
COPY *.png ./
RUN CGO_ENABLED=0 go build -o /main
FROM gcr.io/distroless/static-debian12:latest-amd64 as build-release-stage
WORKDIR /
COPY --from=build-stage /main /main
EXPOSE 8080
USER nonroot:nonroot
CMD ["./main"]
