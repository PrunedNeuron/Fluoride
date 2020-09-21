FROM postgres:9.6-alpine
# RUN localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG en_US.utf8

# Custom initialization scripts
COPY scripts/init_db.sh /docker-entrypoint-initdb.d/init_db.sh
COPY scripts/db/schema.sql /schema.sql
COPY scripts/db/tables.sql /tables.sql

RUN chmod +x /docker-entrypoint-initdb.d/init_db.sh