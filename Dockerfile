FROM golang:1.25 AS build
WORKDIR /src
COPY go.mod main.go index.html ./
RUN CGO_ENABLED=0 go build -trimpath -o /go-node-info .

FROM scratch
COPY --from=build /go-node-info /go-node-info
USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/go-node-info"]
