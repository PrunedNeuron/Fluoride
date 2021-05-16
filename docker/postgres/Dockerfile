FROM postgres:13.3-alpine
ENV LANG en_US.utf8

# Custom initialization scripts
COPY scripts/init_db.sh /docker-entrypoint-initdb.d/init_db.sh
COPY scripts/db/schema.sql /schema.sql
COPY scripts/db/tables.sql /tables.sql

RUN chmod +x /docker-entrypoint-initdb.d/init_db.sh