FROM scratch

ADD main /

CMD mkdir config

COPY config/approved-scheduler.yaml config/

ENTRYPOINT ["/main"]