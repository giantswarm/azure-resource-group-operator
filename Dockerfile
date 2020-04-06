FROM alpine:3.10

RUN apk add --no-cache ca-certificates

ADD ./azure-resource-group-operator /azure-resource-group-operator

ENTRYPOINT ["/azure-resource-group-operator"]
