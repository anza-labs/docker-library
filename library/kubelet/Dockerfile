# platforms=linux/amd64,linux/arm64,linux/ricv64
ARG VERSION=v1.33.0
FROM --platform=$BUILDPLATFORM ghcr.io/anza-labs/library/kubernetes-sources:${VERSION} AS pre-build
FROM --platform=$BUILDPLATFORM tonistiigi/xx AS xx
FROM pre-build AS build
COPY --from=xx / /
ARG TARGETARCH
ARG TARGETOS
ARG TARGETPLATFORM
ARG VERSION

ENV CGO_ENABLED=0

RUN xx-go build \
        -ldflags="$(/opt/ldflags.sh ${VERSION})" \
        -trimpath \
        -o ./bin/kubelet \
        ./cmd/kubelet

# Use distroless as minimal base image to package the builder binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:latest AS image

WORKDIR /

ENV GODEBUG=netdns=go,fips140=on
ENV GOFIPS140=v1.0.0

COPY --from=build /src/bin/kubelet /kubelet

ENTRYPOINT [ "/kubelet" ]
