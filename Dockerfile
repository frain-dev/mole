FROM golang as builder

WORKDIR /go/src/app
COPY . .

RUN go build -o smokescreen .

FROM alpine:3.16.2

COPY --from=builder /go/src/app/smokescreen /usr/local/bin/smokescreen
COPY acl.yaml /etc/smokescreen/acl.yaml

RUN apk add --no-cache gcompat

EXPOSE 4750

CMD ["smokescreen", "--egress-acl-file", "/etc/smokescreen/acl.yaml"]