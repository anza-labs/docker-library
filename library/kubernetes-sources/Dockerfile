FROM docker.io/library/golang:1.24 AS sources

ARG VERSION

WORKDIR /
RUN wget "https://github.com/kubernetes/kubernetes/archive/refs/tags/${VERSION}.tar.gz" && \
    tar -xf ${VERSION}.tar.gz && \
    mv kubernetes-${VERSION#v} /src && \
    rm -rf ${VERSION}.tar.gz
WORKDIR /src
