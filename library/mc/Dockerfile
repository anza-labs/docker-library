ARG VERSION
FROM docker.io/minio/mc:${VERSION} AS minio

FROM docker.io/library/alpine:3.22.1 AS image

# install jq
RUN apk add --no-cache jq

# install mc
COPY --from=minio /usr/bin/mc /usr/local/bin/mc

# initialize it
ENV MC_DISABLE_PAGER=1
RUN /usr/local/bin/mc || true
