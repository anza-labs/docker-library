# platforms=linux/amd64,linux/arm64
FROM docker.io/library/alpine:3.22.1 AS alpine

ARG VERSION

WORKDIR /workspace
RUN wget "https://github.com/firecracker-microvm/firecracker/releases/download/${VERSION}/firecracker-${VERSION}-$(uname -m).tgz" && \
    mv firecracker-*.tgz firecracker.tgz && \
    tar -xvf firecracker.tgz && \
    mv "release-${VERSION}-$(uname -m)/firecracker-${VERSION}-$(uname -m)" firecracker

# Use distroless as minimal base image to package the builder binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:latest AS image
WORKDIR /
COPY --from=alpine /workspace/firecracker /bin/firecracker
ENTRYPOINT ["/bin/firecracker"]
