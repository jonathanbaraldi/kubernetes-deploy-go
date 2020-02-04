FROM golang:1.11
EXPOSE 80
COPY ./bin/hello-server /usr/local/bin/
ENV GOKUBE v5
CMD ["hello-server"]
