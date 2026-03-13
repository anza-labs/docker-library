ARG BASE_IMAGE
FROM ${BASE_IMAGE} AS build-base

# renovate: datasource=repology depName=alpine_edge/ca-certificates versioning=loose
ARG CA_CERTIFICATES_VERSION="20251003-r0"
# renovate: datasource=repology depName=alpine_edge/cni-plugins versioning=loose
ARG CNI_PLUGINS_VERSION="1.9.0-r3"
# renovate: datasource=repology depName=alpine_edge/cri-o versioning=loose
ARG CRIO_VERSION="1.34.3-r3"
# renovate: datasource=repology depName=alpine_edge/crun versioning=loose
ARG CRUN_VERSION="1.26-r0"
# renovate: datasource=repology depName=alpine_edge/kubelet versioning=loose
ARG KUBELET_VERSION="1.35.1-r1"
# renovate: datasource=repology depName=alpine_edge/libcap versioning=loose
ARG LIBCAP_VERSION="2.77-r0"
# renovate: datasource=repology depName=alpine_edge/libseccomp versioning=loose
ARG LIBSECCOMP_VERSION="2.6.0-r1"
# renovate: datasource=repology depName=alpine_edge/nftables versioning=loose
ARG NFTABLES_VERSION="1.1.6-r0"
# renovate: datasource=repology depName=alpine_edge/runit versioning=loose
ARG RUNIT_VERSION="2.3.1-r0"

RUN apk add --no-cache \
    "ca-certificates=${CA_CERTIFICATES_VERSION}" \
    "cni-plugins=${CNI_PLUGINS_VERSION}" \
    "cri-o=${CRIO_VERSION}" \
    "crun=${CRUN_VERSION}" \
    "kubelet=${KUBELET_VERSION}" \
    "libcap=${LIBCAP_VERSION}" \
    "libseccomp=${LIBSECCOMP_VERSION}" \
    "nftables=${NFTABLES_VERSION}" \
    "runit=${RUNIT_VERSION}"

RUN apk del --rdepends \
    alpine-keys \
    alpine-baselayout \
    alpine-baselayout-data \
    alpine-release \
    apk-tools

RUN mkdir -p /etc/service /var/log

COPY --chown=0:0 ./etc/sv/crio/run /etc/sv/crio/run
COPY --chown=0:0 ./etc/crio/crio.conf /etc/crio/crio.conf
COPY --chown=0:0 ./etc/containers/storage.conf /etc/containers/storage.conf
RUN ln -s /etc/sv/crio /etc/service/ \
    && chmod 0755 /etc/sv/crio/run

COPY --chown=0:0 ./etc/sv/kubelet/run /etc/sv/kubelet/run
COPY --chown=0:0 ./etc/kubernetes/kubelet.conf /etc/kubernetes/kubelet.conf
RUN ln -s /etc/sv/kubelet /etc/service/ \
    && chmod 0755 /etc/sv/kubelet/run

FROM scratch AS image
COPY --from=build-base / /
ENV PATH="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
ENTRYPOINT [ "/usr/bin/runsvdir", "-P", "/etc/service" ]
