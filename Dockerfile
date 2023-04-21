FROM ubuntu:22.04

RUN apt-get update
RUN apt-get install ca-certificates -y
COPY ./bin/botlevi .

CMD ["./botlevi"]   