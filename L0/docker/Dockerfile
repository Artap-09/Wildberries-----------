# syntax=docker/dockerfile:1
FROM alpine:3
MAINTAINER Peter Artemiev
ENV PSQL_NAME="postgres" PSQL_PASS="qwerty" PSQL_HOST="PSQL" PSQL_PORT="5432" PSQL_DB="postgres" PSQL_SLL="disable"
ADD L0 .
RUN chmod u+x L0
EXPOSE 4969
CMD /L0