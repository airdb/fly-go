FROM airdb/base:latest

WORKDIR /srv

ADD ./output /srv/output
ADD ./.env.* /srv/output

ENTRYPOINT ["./output/main"]
CMD ["serve"]
#CMD ["config", "./config/config.yml", "--additional-files", "./config/config_yummy.yml"]