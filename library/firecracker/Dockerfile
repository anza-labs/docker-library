ARG BASE_IMAGE
FROM docker.io/library/alpine:3.22.1 AS alpine

ARG VERSION

WORKDIR /workspace
RUN wget "https://github.com/firecracker-microvm/firecracker/releases/download/${VERSION}/firecracker-${VERSION}-$(uname -m).tgz" && \
    mv firecracker-*.tgz firecracker.tgz && \
    tar -xvf firecracker.tgz && \
    mv "release-${VERSION}-$(uname -m)/firecracker-${VERSION}-$(uname -m)" firecracker


FROM ${BASE_IMAGE} AS image
WORKDIR /
COPY --from=alpine /workspace/firecracker /usr/local/bin/firecracker
ENTRYPOINT ["/usr/local/bin/firecracker"]
