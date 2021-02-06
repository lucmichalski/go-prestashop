FROM alpine:3.13 as get-prestashop

ARG PRESTASHOP_VERSION=${PRESTASHOP_VERSION:-"1.7.6.2"}

RUN apk --no-cache add wget
RUN mkdir /tmp/app /tmp/download && \
    cd /tmp/download && \
    wget https://github.com/PrestaShop/PrestaShop/releases/download/${PRESTASHOP_VERSION}/prestashop_${PRESTASHOP_VERSION}.zip && \
    unzip -q prestashop_${PRESTASHOP_VERSION}.zip -d . && \
    unzip -q prestashop.zip -d /tmp/app

FROM php:7.2-fpm-alpine

RUN apk --no-cache add \
    bash \
    nano \
    jq \
    mariadb-client

# install extensions
# intl, zip, soap
RUN apk add --update --no-cache libintl gettext-dev icu icu-dev libxml2-dev \
    && docker-php-ext-install intl zip soap gettext mbstring

# mysqli, pdo, pdo_mysql, pdo_pgsql
RUN apk add --update --no-cache postgresql-dev \
    && docker-php-ext-install mysqli pdo pdo_mysql

# mcrypt, gd, iconv
RUN apk add --update --no-cache \
        freetype-dev \
        libjpeg-turbo-dev \
        libmcrypt-dev \
        libpng-dev \
    && docker-php-ext-install -j"$(getconf _NPROCESSORS_ONLN)" iconv \
    && docker-php-ext-configure gd --with-freetype-dir=/usr/include/ --with-jpeg-dir=/usr/include/ \
    && docker-php-ext-install -j"$(getconf _NPROCESSORS_ONLN)" gd

# php-redis
ENV PHPREDIS_VERSION 5.3.2

RUN docker-php-source extract \
    && curl -L -o /tmp/redis.tar.gz https://github.com/phpredis/phpredis/archive/$PHPREDIS_VERSION.tar.gz \
    && tar xfz /tmp/redis.tar.gz \
    && rm -r /tmp/redis.tar.gz \
    && mv phpredis-$PHPREDIS_VERSION /usr/src/php/ext/redis \
    && docker-php-ext-install redis \
    && docker-php-source delete

# apcu
RUN docker-php-source extract \
    && apk add --no-cache --virtual .phpize-deps-configure $PHPIZE_DEPS \
    && pecl install apcu \
    && docker-php-ext-enable apcu \
    && apk del .phpize-deps-configure \
    && docker-php-source delete

RUN apk add --no-cache curl-dev && docker-php-ext-configure curl --with-curl=/usr/lib64 && docker-php-ext-install curl 

RUN docker-php-ext-install phar && \
    docker-php-ext-install simplexml && \
    php -m

RUN addgroup prestashop && \
  adduser -HD -G prestashop php

COPY --from=get-prestashop --chown=php:prestashop /tmp/app /app
# COPY config/php.ini /usr/local/etc/php/php.ini
COPY config/ /usr/local/etc/
COPY docker-entrypoint.sh /var/opt/

RUN chmod +x /var/opt/docker-entrypoint.sh && \
    rm -f /usr/local/etc/php-fpm.d/www.conf && \
    rm -f /usr/local/etc/php-fpm.d/zz-docker.conf && \
    rm -f /usr/local/etc/php-fpm.d/docker.conf && \
    rm -f /usr/local/etc/php-fpm.d/www.conf.default && \
    ls -l /usr/local/etc/php-fpm.d/

EXPOSE 9000
WORKDIR /app

RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/bin --filename=composer

RUN wget https://github.com/sas-adilis/ps_console/archive/master.zip && \
    unzip -q master.zip -d . && \
    rm -f master.zip && \
    mv ps_console-master console && \
    chmod 755 console && \
    cd console && \
    composer install

RUN chown -Rf php:prestashop /app

# VOLUME ["/app"]
# USER php

CMD ["bash", "/var/opt/docker-entrypoint.sh"]