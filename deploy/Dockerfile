FROM golang:alpine AS build

WORKDIR $GOPATH/src/item-search
ADD .. ./
RUN go build -ldflags "-s -w"  -o item-search .

FROM scratch As prod
COPY --chown=65534:0  --from=build  /go/src/item-search .
USER 65534
EXPOSE 8080
CMD ["./item-search"]