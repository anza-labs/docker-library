# platforms=linux/amd64,linux/arm64,linux/ppc64le,linux/riscv64
FROM docker.io/library/golang:1.24-alpine AS build

# renovate: datasource=github-tags depName=kubernetes-sigs/cloud-provider-kind
ARG VERSION=v0.6.0
RUN go install sigs.k8s.io/cloud-provider-kind@${VERSION}

# build real cloud-provider-kind image
FROM docker:25.0-cli
COPY --from=0 --chown=root:root /go/bin/cloud-provider-kind /bin/cloud-provider-kind
ENTRYPOINT ["/bin/cloud-provider-kind"]
