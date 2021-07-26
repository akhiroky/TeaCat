FROM alpine:3.10.1

ARG version=2.0.0
ARG os = linux

LABEL maintainer="Hiroki Yamashita" \
      description="The app is that implements of wc command and cat command."

RUN    adduser -D TeaCat \
    && apk --no-cache add --update --virtual .builddeps curl tar \
    && curl -s -L -O https://github.com/akhiroky/TeaCat/realeases/download/v${version}/TeaCat-${version}__amd64.tar.gz \
    && tar xfz TeaCat-${version}_${os}_amd64.tar.gz        \
    && mv TeaCat-${version} /opt                           \
    && ln -s /opt/TeaCat-${version} /opt/TeaCat               \
    && ln -s /opt/TeaCat-${version}/TeaCat /usr/local/bin/TeaCat \
    && rm TeaCat-${version}_${os}_amd64.tar.gz             \
    && apk del --purge .builddeps

ENV HOME="/home/TeaCat"

WORKDIR ${HOME}
USER    TeaCat

ENTRYPOINT [ "/opt/TeaCat/TeaCat" ]
