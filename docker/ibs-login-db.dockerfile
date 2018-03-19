FROM mariadb:latest

ADD sql/init.sql /docker-entrypoint-initdb.d/
