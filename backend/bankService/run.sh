#!/bin/bash

sleep 3

echo "Running app..."

python manage.py migrate --noinput

exec python manage.py runserver 0.0.0.0:8080
