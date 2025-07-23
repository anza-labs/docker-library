# platforms=linux/amd64,linux/arm64
FROM docker.io/library/golang:1.24 AS sources

# renovate: datasource=github-tags depName=kubernetes/kubernetes
ARG VERSION=v1.33.0

WORKDIR /
RUN wget "https://github.com/kubernetes/kubernetes/archive/refs/tags/${VERSION}.tar.gz" && \
    tar -xf ${VERSION}.tar.gz && \
    mv kubernetes-${VERSION#v} /src && \
    rm -rf ${VERSION}.tar.gz
WORKDIR /src
