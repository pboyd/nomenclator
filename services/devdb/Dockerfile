# This is a minor extension to the standard postgres Docker image, which
# creates a second database for testing.

FROM postgres:15-alpine

COPY create_databases.sql /docker-entrypoint-initdb.d/create_databases.sql
