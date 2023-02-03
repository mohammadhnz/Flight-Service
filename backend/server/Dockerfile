FROM node:14
WORKDIR /

RUN npm install
COPY . /

RUN chmod 0777 /run.sh

ENTRYPOINT ["/run.sh"]
