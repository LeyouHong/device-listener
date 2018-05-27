FROM golang:1.8 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux

ADD . /go/src/github.com/device-listener
RUN cd /go/src/github.com/device-listener && make

COPY --from=builder /go/src/github.com/device-listener/build/device-listener /usr/local/bin/

ENV PORT=32770
ENV NAME=DEVICE_LISTENER

EXPOSE 32770

CMD ["device-listener"]
