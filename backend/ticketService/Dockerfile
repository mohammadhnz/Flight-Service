FROM node:14
WORKDIR /

COPY . /
RUN npm install
RUN chmod 0777 /run.sh

ENTRYPOINT ["/run.sh"]
