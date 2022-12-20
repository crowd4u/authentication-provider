FROM php:7.2-apache
COPY --from=composer /usr/bin/composer /usr/bin/composer
RUN apt-get update
RUN apt-get install -y vim
RUN docker-php-ext-install pdo_mysql
# RUN apt update && apt install git
RUN echo '\
log_errors = On\n\
error_log = /dev/stderr\n\
error_reporting = E_ALL\n\
' >> /usr/local/etc/php/php.ini \
EXPOSE 80
