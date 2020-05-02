FROM debian

RUN apt-get update && apt-get install -y ca-certificates

COPY build/serverman /usr/bin

CMD [ "/usr/bin/serverman" ]
