FROM golang:1.10.1 as pcdeper
WORKDIR /go/src/git.amabanana.com/plancks-cloud/pc-go-daemon
COPY Gopkg.toml Gopkg.lock vendor ./
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -v -vendor-only

FROM pcdeper as builder
WORKDIR /go/src/git.amabanana.com/plancks-cloud/pc-go-daemon
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pc-go-daemon .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/git.amabanana.com/plancks-cloud/pc-go-daemon/pc-go-daemon .
CMD ["./pc-go-daemon"]