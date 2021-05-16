from alpine

WORKDIR /opt

COPY app .

CMD [ "./app" ]
