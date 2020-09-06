FROM registry.svc.ci.openshift.org/openshift/release:golang-1.10 as builder
RUN go get github.com/kedark3/pacman
WORKDIR /go/src/github.com/kedark3/pacman/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/pacman-exec pacman.go assets_vfsdata.go

FROM scratch
LABEL author="Kedar Kulkarni"

WORKDIR /pacman
COPY --from=builder /go/bin/pacman-exec /pacman/pacman
EXPOSE 43880

ENTRYPOINT ["./pacman"]