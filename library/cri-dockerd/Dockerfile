ARG BASE_IMAGE
FROM ghcr.io/anza-labs/library/cni:v1.7.1 AS cni
FROM docker.io/library/alpine:3.22.1 AS build

# select arch
RUN echo $(case "$(uname -m)" in \
        x86_64|amd64) echo amd64 ;;\
        aarch64|arm64) echo arm64 ;;\
        *) echo "unsupported" ;;\
    esac) > /tmp/arch

ARG VERSION

WORKDIR /src

RUN wget "https://github.com/Mirantis/cri-dockerd/releases/download/${VERSION}/cri-dockerd-${VERSION#v}.$(cat /tmp/arch).tgz" && \
    tar -xf "cri-dockerd-${VERSION#v}.$(cat /tmp/arch).tgz"

FROM ${BASE_IMAGE} AS image
COPY --from=cni / /
COPY --from=build /src/cri-dockerd/cri-dockerd /usr/local/bin/cri-dockerd

ENTRYPOINT [ "/usr/local/bin/cri-dockerd" ]
