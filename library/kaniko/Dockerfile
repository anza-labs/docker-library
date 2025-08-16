ARG BASE_IMAGE
FROM docker.io/library/golang:1.25-alpine AS build

# install some auxiliary packages
RUN apk add --no-cache \
    bash git make

ARG VERSION

WORKDIR /src
RUN wget "https://github.com/chainguard-dev/kaniko/archive/${VERSION}.tar.gz" && \
    tar -xf ${VERSION}.tar.gz && \
    mv kaniko-${VERSION#v} /src/kaniko && \
    rm -rf ${VERSION}.tar.gz
WORKDIR /src/kaniko

RUN make \
    out/executor \
    out/warmer

FROM ghcr.io/anza-labs/library/go-busybox:c2b9e48e151e8efeab13029a9bcf9236eae176ca AS go-busybox
FROM ${BASE_IMAGE} AS executor-debug

ENV PATH=/bin:/usr/bin:/usr/local/bin

COPY --from=build /src/kaniko/out/executor /usr/local/bin/executor
COPY --from=go-busybox / /

FROM ${BASE_IMAGE} AS executor

COPY --from=build /src/kaniko/out/executor /usr/local/bin/executor

ENTRYPOINT [ "/usr/local/bin/executor" ]

FROM ${BASE_IMAGE} AS warmer

COPY --from=build /src/kaniko/out/warmer /usr/local/bin/warmer

ENTRYPOINT [ "/usr/local/bin/warmer" ]
