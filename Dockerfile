FROM ubuntu:14.04

RUN apt-get update

RUN mkdir /var/www/
COPY server /var/www/

