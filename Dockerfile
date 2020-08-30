FROM registry.svc.ci.openshift.org/openshift/release:golang-1.10 as builder
RUN go get github.com/kedark3/pacman
WORKDIR /go/src/github.com/kedark3/pacman/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/pacman-exec

FROM scratch
LABEL author="Kedar Kulkarni"

WORKDIR /pacman
COPY --from=builder /go/bin/pacman-exec /pacman/pacman
VOLUME /pacman/static
EXPOSE 8080

CMD ["./pacman"]