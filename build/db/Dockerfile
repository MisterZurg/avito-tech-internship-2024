FROM postgres:16.2-alpine

ENV POSTGRES_USER postgres

ENV POSTGRES_PASSWORD password

ENV POSTGRES_DB db

EXPOSE 5432

ADD ./build/db/postgres_init.sql /docker-entrypoint-initdb.d/