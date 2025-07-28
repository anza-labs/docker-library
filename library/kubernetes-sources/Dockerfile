ARG BASE_IMAGE
FROM ${BASE_IMAGE} AS image

ARG VERSION

WORKDIR /
RUN wget "https://github.com/kubernetes/kubernetes/archive/refs/tags/${VERSION}.tar.gz" && \
    tar -xf ${VERSION}.tar.gz && \
    mv kubernetes-${VERSION#v} /src && \
    rm -rf ${VERSION}.tar.gz
WORKDIR /src
