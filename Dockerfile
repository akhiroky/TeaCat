FROM alpine:3.10.1

ARG version=1.0.0

LABEL maintainer="Hiroki Yamashita" \
      description="The app is that implements of wc command and cat command."

RUN    adduser -D TeaCat \
    && apk --no-cache add --update --virtual .builddeps curl tar \
#    && curl -s -L -O https://github.com/akhiroky/TeaCat/realeases/download/v${version}/TeaCat-${version}_linux_amd64.tar.gz \
    && curl -s -L -o TeaCat-${version}_linux_amd64.tar.gz https://www.dropbox.com/s/ch2434soyyukn3m/TeaCat-1.0.0_linux_amd64.tar.gz?dl=0 \
    && tar xfz TeaCat-${version}_linux_amd64.tar.gz        \
    && mv TeaCat-${version} /opt                           \
    && ln -s /opt/TeaCat-${version} /opt/TeaCat               \
    && ln -s /opt/TeaCat-${version}/nml /usr/local/bin/TeaCat \
    && rm TeaCat-${version}_linux_amd64.tar.gz             \
    && apk del --purge .builddeps

ENV HOME="/home/TeaCat"

WORKDIR ${HOME}
USER    TeaCat

ENTRYPOINT [ "/opt/TeaCat/TeaCat" ]