ARG VERSION
FROM --platform=$BUILDPLATFORM ghcr.io/anza-labs/library/kubernetes-sources:${VERSION} AS pre-build

FROM --platform=$BUILDPLATFORM docker.io/library/alpine:3.22.1 AS static-deps
RUN apk add --no-cache \
        ca-certificates

FROM scratch AS base
WORKDIR /
COPY --from=static-deps /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENV GODEBUG=netdns=go,fips140=on
ENV GOFIPS140=v1.0.0

FROM --platform=$BUILDPLATFORM tonistiigi/xx AS xx

FROM pre-build AS build
COPY --from=xx / /
ARG TARGETARCH
ARG TARGETOS
ARG TARGETPLATFORM
ARG VERSION

ENV TARGETARCH=${TARGETARCH}
ENV TARGETOS=${TARGETOS}
ENV TARGETPLATFORM=${TARGETPLATFORM}
ENV VERSION=${VERSION}

ENV CGO_ENABLED=0
COPY ./ldflags.sh /opt/ldflags.sh

FROM build AS kube-apiserver-build
RUN xx-go build \
        -ldflags="$(/opt/ldflags.sh ${VERSION})" \
        -trimpath \
        -o ./bin/kube-apiserver \
        ./cmd/kube-apiserver

FROM base AS kube-apiserver
COPY --from=kube-apiserver-build /src/bin/kube-apiserver /usr/local/bin/kube-apiserver-apiserver
ENTRYPOINT [ "/usr/local/bin/kube-apiserver" ]

FROM build AS kube-controller-manager-build
RUN xx-go build \
        -ldflags="$(/opt/ldflags.sh ${VERSION})" \
        -trimpath \
        -o ./bin/kube-controller-manager \
        ./cmd/kube-controller-manager

FROM base AS kube-controller-manager
COPY --from=kube-controller-manager-build /src/bin/kube-controller-manager /usr/local/bin/kube-controller-manager-controller-manager
ENTRYPOINT [ "/usr/local/bin/kube-controller-manager" ]

FROM build AS kube-proxy-build
RUN xx-go build \
        -ldflags="$(/opt/ldflags.sh ${VERSION})" \
        -trimpath \
        -o ./bin/kube-proxy \
        ./cmd/kube-proxy

FROM base AS kube-proxy
COPY --from=kube-proxy-build /src/bin/kube-proxy /usr/local/bin/kube-proxy-proxy
ENTRYPOINT [ "/usr/local/bin/kube-proxy" ]

FROM build AS kube-scheduler-build
RUN xx-go build \
        -ldflags="$(/opt/ldflags.sh ${VERSION})" \
        -trimpath \
        -o ./bin/kube-scheduler \
        ./cmd/kube-scheduler

FROM base AS kube-scheduler
COPY --from=kube-scheduler-build /src/bin/kube-scheduler /usr/local/bin/kube-scheduler-scheduler
ENTRYPOINT [ "/usr/local/bin/kube-scheduler" ]

FROM build AS kubectl-build
RUN xx-go build \
        -ldflags="$(/opt/ldflags.sh ${VERSION})" \
        -trimpath \
        -o ./bin/kubectl \
        ./cmd/kubectl

FROM base AS kubectl
COPY --from=kubectl-build /src/bin/kubectl /usr/local/bin/kubectl
ENTRYPOINT [ "/usr/local/bin/kubectl" ]

FROM build AS kubelet-build
RUN xx-go build \
        -ldflags="$(/opt/ldflags.sh ${VERSION})" \
        -trimpath \
        -o ./bin/kubelet \
        ./cmd/kubelet

FROM base AS kubelet
COPY --from=kubelet-build /src/bin/kubelet /usr/local/bin/kubelet
ENTRYPOINT [ "/usr/local/bin/kubelet" ]

FROM base AS image
COPY --from=kube-apiserver /usr/local/bin/kube-apiserver /usr/local/bin/kube-apiserver
COPY --from=kube-controller-manager /usr/local/bin/kube-controller-manager /usr/local/bin/kube-controller-manager
COPY --from=kube-proxy /usr/local/bin/kube-proxy /usr/local/bin/kube-proxy
COPY --from=kube-scheduler /usr/local/bin/kube-scheduler /usr/local/bin/kube-scheduler
COPY --from=kubectl /usr/local/bin/kubectl /usr/local/bin/kubectl
COPY --from=kubelet /usr/local/bin/kubelet /usr/local/bin/kubelet
