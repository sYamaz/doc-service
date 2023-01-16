FROM postgres:14-alpine3.17

COPY .docker/db/init/ docker-entrypoint-initdb.d
#RUN chmod 755 docker-entrypoint-initdb.d/init.sh

CMD docker-entrypoint.sh postgres