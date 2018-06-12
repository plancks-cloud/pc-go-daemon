FROM golang:1.10.1 as pcdeper
WORKDIR /go/src/git.amabanana.com/plancks-cloud/pc-go-daemon
COPY Gopkg.toml Gopkg.lock vendor ./
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -v -vendor-only

FROM pcdeper as builder
WORKDIR /go/src/git.amabanana.com/plancks-cloud/pc-go-daemon
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o plancks-cloud .

FROM scratch
WORKDIR /
COPY --from=builder /go/src/git.amabanana.com/plancks-cloud/pc-go-daemon/plancks-cloud .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/plancks-cloud"]