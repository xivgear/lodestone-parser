FROM golang:1.17 AS BUILDER
WORKDIR /go/src/github.com/xivgear/lodestone-parser/
COPY . .
RUN make build

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=BUILDER /go/src/github.com/xivgear/lodestone-parser/app ./
CMD ["./app"]
