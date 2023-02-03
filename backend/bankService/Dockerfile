FROM python:3.6

WORKDIR /

RUN pip install django djangorestframework
COPY . /

RUN chmod 0777 /run.sh

ENTRYPOINT ["/run.sh"]
