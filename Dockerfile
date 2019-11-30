FROM ubuntu:18.04

# install PostgreSQL
ENV PGSQLVER 10
RUN apt-get update &&\
    apt-get install -y postgresql-$PGSQLVER postgresql-contrib &&\
    apt-get install -y git &&\
    apt-get install -y wget

# create PostgreSQL database
USER postgres
RUN    /etc/init.d/postgresql start &&\
    psql --command "CREATE USER forum WITH SUPERUSER PASSWORD 'forum';" &&\
    createdb -O forum forum &&\
    /etc/init.d/postgresql stop
RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/$PGSQLVER/main/pg_hba.conf &&\
    echo "listen_addresses='*'" >> /etc/postgresql/$PGSQLVER/main/postgresql.conf &&\
    echo "default_text_search_config = 'pg_catalog.english'" >> /etc/postgresql/$PGSQLVER/main/postgresql.conf

EXPOSE 5432

# Golang installing
ENV GOVERSION 1.12
USER root
RUN wget https://storage.googleapis.com/golang/go$GOVERSION.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go$GOVERSION.linux-amd64.tar.gz
ENV GOROOT /usr/local/go
ENV GOPATH /opt/go
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH
EXPOSE 5000

WORKDIR /DB_TP
COPY . .

CMD service postgresql start && go run simple-server.go
