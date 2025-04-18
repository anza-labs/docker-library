# platforms=linux/amd64,linux/arm64
FROM ghcr.io/grpc-ecosystem/grpc-health-probe:v0.4.37 AS probe
FROM docker.io/library/alpine:3.21.3 AS alpine

# select arch
RUN echo $(case "$(uname -m)" in \
        x86_64|amd64) echo amd64 ;;\
        aarch64|arm64) echo arm64 ;;\
        *) echo "unsupported" ;;\
    esac) > /tmp/arch

# renovate: datasource=github-tags depName=k3s-io/kine
ARG VERSION=v0.13.14

WORKDIR /workspace
RUN wget "https://github.com/k3s-io/kine/releases/download/${VERSION}/kine-$(cat /tmp/arch)" && \
    mv kine-* kine && \
    chmod +x kine

# Use distroless as minimal base image to package the builder binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:latest AS image
WORKDIR /
COPY --from=alpine /workspace/kine /bin/kine
COPY --from=probe /ko-app/grpc-health-probe /bin/grpc_health_probe
ENTRYPOINT ["/bin/kine"]
