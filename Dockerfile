FROM ubuntu:22.04

COPY ./vendor .
COPY ./bin/botlevi .

CMD ["./botlevi"]