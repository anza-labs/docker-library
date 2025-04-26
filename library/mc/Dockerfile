# platforms=linux/amd64,linux/arm64,linux/ppc64le
# renovate: datasource=github-tags depName=minio/mc versioning=regex:^RELEASE\.?(?<major>\d+)-(?<minor>\d+)-(?<patch>\d+)(.+)?$
ARG VERSION=RELEASE.2025-04-16T18-13-26Z
FROM docker.io/minio/mc:${VERSION} AS minio

FROM docker.io/library/alpine:3.21.3 AS image

# install jq
RUN apk add --no-cache jq

# install mc
COPY --from=minio /usr/bin/mc /usr/bin/mc

# initialize it
ENV MC_DISABLE_PAGER=1
RUN mc || true
