#!/bin/bash

echo "Applying php settings"
echo "* memory_limit = $PHP_MEMORY_LIMIT"
echo "* max_execution_time = $PHP_MAX_EXECUTION_TIME"
echo "* max_input_time = $PHP_MAX_INPUT_TIME"
echo "* display_errors = $PHP_DISPLAY_ERRORS"

mv /usr/local/etc/php/php.ini-development /usr/local/etc/php/php.ini

sed -i -e "s|memory_limit.*|memory_limit = $PHP_MEMORY_LIMIT|" \
  -e "s|max_execution_time.*|max_execution_time = $PHP_MAX_EXECUTION_TIME|" \
  -e "s|session.cache_limiter.*|session.cache_limiter = public|" \
  -e "s|display_errors.*|display_errors = $PHP_DISPLAY_ERRORS|" \
  -e "s|max_input_time.*|max_input_time = $PHP_MAX_INPUT_TIME|" /usr/local/etc/php/php.ini

# if [ -f "/app/config/settings.inc.php" ]; then
#   echo "Prestashop is already installed."
#   ls -l
#   php-fpm
#   exit $?
# fi

echo "Installing Prestashop..."

if [ $PS_DIR_INSTALL != "install" ]; then
  mv /app/install /app/$PS_DIR_INSTALL
  echo "* install dir is now named '$PS_DIR_INSTALL'"
  ls -l /app/$PS_DIR_INSTALL
  rm -fR /app/$PS_DIR_INSTALL/fixtures/fashion/*
  mkdir -p /app/$PS_DIR_INSTALL/fixtures/fashion
  tar xvf /tarballs/generated_1M.tar.gz -C /app/$PS_DIR_INSTALL/fixtures/fashion/
  ls -l /app/$PS_DIR_INSTALL/fixtures/fashion
  echo "* Applying patches..."
  sed -i 's/dni=\"\"/dni=\"0000\"/g' /app/$PS_DIR_INSTALL/fixtures/fashion/data/address.xml
fi
if [ $PS_DIR_ADMIN != "admin" ]; then
  mv /app/admin /app/$PS_DIR_ADMIN
  echo "* admin dir is now named '$PS_DIR_ADMIN'"
fi

echo "* waiting for sql server to ready up..."
SQL_S=1
while [ $SQL_S -ne 0 ]; do
  mysql -h $DB_SERVER -P $DB_PORT -u $DB_USER -p$DB_PASSWORD -e "show status" > /dev/null 2>&1
  SQL_S=$?
  echo -n ". "
  sleep 1
done
echo ""

echo "* installing..."

su php -c "php /app/$PS_DIR_INSTALL/index_cli.php \
  --domain=\"$PS_DOMAIN\" \
  --db_server=\"$DB_SERVER\" \
  --db_user=\"$DB_USER\" \
  --db_password=\"$DB_PASSWORD\" \
  --db_name=\"$DB_NAME\" \
  --db_clear=$DB_CLEAR \
  --db_create=$DB_CREATE \
  --prefix=\"$DB_PREFIX\" \
  --name=\"$PS_SHOP_NAME\" \
  --language=\"$PS_LANGUAGE\" \
  --country=\"$PS_SHOP_COUNTRY\" \
  --timezone=\"$PS_TIMEZONE\" \
  --firstname=\"$PS_FIRSTNAME\" \
  --lastname=\"$PS_LASTNAME\" \
  --password=\"$PS_PASSWORD\" \
  --email=\"$PS_EMAIL\" \
  --newsletter=\"$PS_NEWSLETTER\" \
  --ssl=$PS_SSL \
  --rewrite_engine=$PS_REWRITE "

INSTALL_STATUS=$?

if [ $INSTALL_STATUS -ne 0 ]; then
  echo "INSTALLATION FAILED!"
  exit $INSTALL_STATUS
else

  # Required when SSL MODE is activated in Prestashop
  sed -e 's/<IfModule mod_env.c>/<IfModule mod_env.c>\nSetEnv HTTPS On\n/' -i /app/.htaccess

  # install prestarocket theme for v1.7.7.x
  # cd /app/console
  # php psc.php theme:get https://github.com/lucmichalski/classic-rocket/archive/3.0.5.zip
  cd /tmp
  wget https://github.com/lucmichalski/classic-rocket/archive/3.0.5.zip
  unzip -q 3.0.5.zip -d /app/themes/ && \
  mv /app/themes/classic-rocket-3.0.5 /app/themes/classic-rocket && \
  rm -f 3.0.5.zip

  # install prestarocket modules
  cd /tmp
  wget https://github.com/PrestaShop/ps_searchbarjqauto/archive/master.zip 
  unzip -q master.zip -d /app/modules/ && \
  rm -f master.zip && \
  mv /app/modules/ps_searchbarjqauto-master /app/modules/ps_searchbarjqauto && \
  cd /app/console
  php psc.php module:install ps_searchbarjqauto

  echo "Installation complete"
  chown -Rf php:prestashop /app

  echo "List themes"
  ls -l /app/themes

  echo "List modules"
  ls -l /app/modules

  # disable for import purpose of the fake dataset
  #echo "Delete install folder"
  #rm -fR /app/install*

  find /app -type f -exec chmod 644 {} \; && find /app -type d -exec chmod 755 {} \;
  php-fpm
  exit $?
fi