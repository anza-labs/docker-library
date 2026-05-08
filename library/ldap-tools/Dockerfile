ARG BASE_IMAGE
FROM ${BASE_IMAGE} AS image

# renovate: datasource=repology depName=alpine_edge/openldap-clients versioning=loose
ARG OPENLDAP_CLIENTS_VERSION=2.6.10-r0

# install packages
RUN apk add --no-cache \
    openldap-clients="${OPENLDAP_CLIENTS_VERSION}"
