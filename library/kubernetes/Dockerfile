ARG VERSION
FROM --platform=$BUILDPLATFORM ghcr.io/anza-labs/library/kubernetes-sources:${VERSION} AS pre-build

FROM scratch AS base
WORKDIR /
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
COPY --from=kube-apiserver-build /src/bin/kube-apiserver /bin/kube-apiserver-apiserver
ENTRYPOINT [ "/bin/kube-apiserver" ]

FROM build AS kube-controller-manager-build
RUN xx-go build \
        -ldflags="$(/opt/ldflags.sh ${VERSION})" \
        -trimpath \
        -o ./bin/kube-controller-manager \
        ./cmd/kube-controller-manager

FROM base AS kube-controller-manager
COPY --from=kube-controller-manager-build /src/bin/kube-controller-manager /bin/kube-controller-manager-controller-manager
ENTRYPOINT [ "/bin/kube-controller-manager" ]

FROM build AS kube-proxy-build
RUN xx-go build \
        -ldflags="$(/opt/ldflags.sh ${VERSION})" \
        -trimpath \
        -o ./bin/kube-proxy \
        ./cmd/kube-proxy

FROM base AS kube-proxy
COPY --from=kube-proxy-build /src/bin/kube-proxy /bin/kube-proxy-proxy
ENTRYPOINT [ "/bin/kube-proxy" ]

FROM build AS kube-scheduler-build
RUN xx-go build \
        -ldflags="$(/opt/ldflags.sh ${VERSION})" \
        -trimpath \
        -o ./bin/kube-scheduler \
        ./cmd/kube-scheduler

FROM base AS kube-scheduler
COPY --from=kube-scheduler-build /src/bin/kube-scheduler /bin/kube-scheduler-scheduler
ENTRYPOINT [ "/bin/kube-scheduler" ]

FROM build AS kubectl-build
RUN xx-go build \
        -ldflags="$(/opt/ldflags.sh ${VERSION})" \
        -trimpath \
        -o ./bin/kubectl \
        ./cmd/kubectl

FROM base AS kubectl
COPY --from=kubectl-build /src/bin/kubectl /bin/kubectl
ENTRYPOINT [ "/bin/kubectl" ]

FROM build AS kubelet-build
RUN xx-go build \
        -ldflags="$(/opt/ldflags.sh ${VERSION})" \
        -trimpath \
        -o ./bin/kubelet \
        ./cmd/kubelet

FROM base AS kubelet
COPY --from=kubelet-build /src/bin/kubelet /bin/kubelet
ENTRYPOINT [ "/bin/kubelet" ]

FROM base AS image
COPY --from=kube-apiserver /bin/kube-apiserver /bin/kube-apiserver
COPY --from=kube-controller-manager /bin/kube-controller-manager /bin/kube-controller-manager
COPY --from=kube-proxy /bin/kube-proxy /bin/kube-proxy
COPY --from=kube-scheduler /bin/kube-scheduler /bin/kube-scheduler
COPY --from=kubectl /bin/kubectl /bin/kubectl
COPY --from=kubelet /bin/kubelet /bin/kubelet
