# renovate: datasource=github-tags depName=minio/mc
ARG VERSION=RELEASE.2025-02-21T16-00-46Z-cpuv1
FROM minio/mc:${VERSION} AS minio

FROM alpine:3.21 AS image

# install jq
RUN apk add --no-cache jq

# install mc
COPY --from=minio /usr/bin/mc /usr/bin/mc

# initialize it
ENV MC_DISABLE_PAGER=1
RUN mc || true
