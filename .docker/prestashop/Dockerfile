FROM alpine:latest as get-prestashop

RUN apk --no-cache add wget
RUN mkdir /tmp/app /tmp/download && \
    cd /tmp/download && \
    wget https://github.com/PrestaShop/PrestaShop/releases/download/1.7.6.2/prestashop_1.7.6.2.zip && \
    unzip -q prestashop_1.7.6.2.zip -d . && \
    unzip -q prestashop.zip -d /tmp/app

FROM alpine:latest

RUN apk --no-cache add \
    bash \
    php7 \
    mariadb-client \
    php7-fpm \
    php7-gd \
    php7-intl \
    php7-opcache \
    php7-pdo \
    php7-mbstring \
    php7-mysqli \
    php7-curl \
    php7-dom \
    php7-fileinfo \
    php7-zip \
    php7-simplexml \
    php7-json \
    php7-session \
    php7-tokenizer \
    php7-pdo \
    php7-pdo_mysql \
    php7-iconv \
    php7-posix \
    php7-pecl-mcrypt \
    php7-pecl-imagick \
    icu

RUN addgroup prestashop && \
  adduser -HD -G prestashop php

COPY --from=get-prestashop --chown=php:prestashop /tmp/app /app
COPY config/ /etc/php7/
COPY start.sh /var/opt/

RUN chmod +x /var/opt/start.sh && \
  rm -f /etc/php7/php-fpm.d/www.conf

EXPOSE 9000
WORKDIR /app

CMD ["bash", "/var/opt/start.sh"]
