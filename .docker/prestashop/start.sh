#!/bin/bash

echo "Applying php settings"
echo "* memory_limit = $PHP_MEMORY_LIMIT"
echo "* max_execution_time = $PHP_MAX_EXECUTION_TIME"
echo "* max_input_time = $PHP_MAX_INPUT_TIME"
sed -i -e "s|memory_limit.*|memory_limit = $PHP_MEMORY_LIMIT|" \
  -e "s|max_execution_time.*|max_execution_time = $PHP_MAX_EXECUTION_TIME|" \
  -e "s|max_input_time.*|max_input_time = $PHP_MAX_INPUT_TIME|" /etc/php7/php.ini

if [ -f "/app/config/settings.inc.php" ]; then
  echo "Prestashop is already installed."
  php-fpm7
  exit $?
fi

echo "Installing Prestashop..."

if [ $PS_DIR_INSTALL != "install" ]; then
  mv /app/install /app/$PS_DIR_INSTALL
  echo "* install dir is now named '$PS_DIR_INSTALL'"
fi
if [ $PS_DIR_ADMIN != "admin" ]; then
  mv /app/admin /app/$PS_DIR_ADMIN
  echo "* admin dir is now named '$PS_DIR_ADMIN'"
fi

echo "* waiting for sql server to ready up..."
SQL_S=1
while [ $SQL_S -ne 0 ]; do
  mysql -h "$DB_SERVER" -P $DB_PORT -u "$DB_USER" -p"$DB_PASSWORD" -e "show status" > /dev/null 2>&1
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
  --name=\"$PS_LANGUAGE\" \
  --country=\"$PS_SHOP_COUNTRY\" \
  --timezone=\"$PS_TIMEZONE\" \
  --firstname=\"$PS_FIRSTNAME\" \
  --lastname=\"$PS_LASTNAME\" \
  --password=\"$PS_PASSWORD\" \
  --email=\"$PS_EMAIL\" \
  --newsletter=\"$PS_NEWSLETTER\" \
  --ssl=$PS_SSL"

INSTALL_STATUS=$?

if [ $INSTALL_STATUS -ne 0 ]; then
  echo "INSTALLATION FAILED!"
  exit $INSTALL_STATUS
else
  echo "Installation complete"
  php-fpm7
  exit $?
fi
