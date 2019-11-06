FROM postgres:10
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=${FANTASY_DB_PASS}

COPY ./setup.sql /docker-entrypoint-initdb.d/