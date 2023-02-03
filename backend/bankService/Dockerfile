FROM python:3.6

RUN apt-get update && apt-get install -y libpq-dev
WORKDIR /

RUN pip install django djangorestframework
COPY . /

RUN chmod 0777 /run.sh

ENTRYPOINT ["/run.sh"]
