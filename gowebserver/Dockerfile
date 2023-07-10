FROM golang:1.20 as build

WORKDIR /go/src/gowebserver
COPY . .

ENV CGO_ENABLED=0
RUN go build -o /go/bin/gowebserver

FROM gcr.io/distroless/static

COPY --from=build /go/bin/gowebserver /
CMD ["/gowebserver"]
