FROM scratch

ADD earnings-scheduler /

CMD mkdir config

COPY config/earnings-scheduler.yaml config/

ENTRYPOINT ["/earnings-scheduler"]