ARG BASE_IMAGE
FROM ${BASE_IMAGE} AS image

ARG VERSION

# renovate: datasource=repology depName=alpine_edge/font-freefont versioning=loose
ARG FONT_FREEFONT_VERSION=20120503-r4

# renovate: datasource=repology depName=alpine_edge/font-noto-emoji versioning=loose
ARG FONT_NOTO_EMOJI_VERSION=2.051-r0

# renovate: datasource=repology depName=alpine_edge/font-wqy-zenhei versioning=loose
ARG FONT_WQY_ZENHEI_VERSION=0.9.45-r5

# Installs latest Chromium package.
RUN apk upgrade --no-cache --available \
    && apk add --no-cache \
        chromium="${VERSION}" \
        chromium-swiftshader="${VERSION}" \
        font-freefont="${FONT_FREEFONT_VERSION}" \
        font-noto-emoji="${FONT_NOTO_EMOJI_VERSION}" \
        font-wqy-zenhei="${FONT_WQY_ZENHEI_VERSION}"

COPY ./etc/fonts/local.conf /etc/fonts/local.conf

RUN mkdir -p /usr/src/app \
    && adduser -D chrome \
    && chown -R chrome:chrome /usr/src/app

USER chrome
WORKDIR /usr/src/app

ENTRYPOINT [ "chromium-browser", \
    "--headless", \
    "--no-sandbox", \
    "--no-first-run", \
    "--no-default-browser-check" \
]
