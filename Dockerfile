FROM golang:1.16-alpine as builder

RUN apk add --no-cache make
WORKDIR /bake
COPY . ./
RUN make install-tools build

FROM golang:1.16-alpine

ENV USER=docker
ENV UID=7193

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/etc/bake" \
    --no-create-home \
    --uid "$UID" \
    "$USER"

WORKDIR /etc/bake
COPY --from=builder /bake/build/bake /usr/local/bin/bake

USER docker
ENTRYPOINT ["bake"]
