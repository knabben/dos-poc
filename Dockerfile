FROM golang:bookworm as build
WORKDIR /go/src/github.com/knabben/dos-poc
COPY . .
RUN go mod download
RUN go vet -v
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11
COPY --from=build /go/bin/app /
COPY ./certs certs/
EXPOSE 6443
CMD ["/app"]
