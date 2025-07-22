# platforms=linux/amd64,linux/arm64
FROM docker.io/library/debian:12.11 AS image

# select arch
RUN echo $(case "$(uname -m)" in \
        x86_64|amd64) echo x86_64 ;;\
        aarch64|arm64) echo aarch64 ;;\
        *) echo "unsupported" ;;\
    esac) > /tmp/arch

# install some auxiliary packages
RUN apt-get update && \
    apt-get install --yes --no-install-recommends \
        bash bc ca-certificates curl git make pkgconf sudo tar xz-utils && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# renovate: datasource=github-tags depName=amber-lang/amber
ARG VERSION="0.4.0-alpha"

# download and install amber
RUN curl -sL "https://github.com/amber-lang/amber/releases/download/${VERSION}/amber-$(cat /tmp/arch)-unknown-linux-gnu.tar.xz" > amber.tar.gz && \
    tar -xf amber.tar.gz && \
    mv amber /usr/bin/amber && \
    rm -rf /tmp/arch /amber.tar.gz

# set entrypoint
ENTRYPOINT [ "/usr/bin/amber" ]
