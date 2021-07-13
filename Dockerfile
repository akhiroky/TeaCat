FROM alpine:3.10.1

ARG version=1.0.0-alpha
ARG os = linux

LABEL maintainer="Hiroki Yamashita" \
      description="The app is that implements of wc command and cat command."

RUN    adduser -D TeaCat \
    && apk --no-cache add --update --virtual .builddeps curl tar \
#    && curl -s -L -O https://github.com/akhiroky/TeaCat/realeases/download/v${version}/TeaCat-${version}_linux_amd64.tar.gz \
    && curl -s -L -o TeaCat-${version}_${os}_amd64.tar.gz https://github.com/akhiroky/TeaCat/releases/download/v#{version}/TeaCat-${version}_${os}_amd64.tar.gz \
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
