import csv
import pytz
import pandas as pd
from uuid import uuid4
import numpy as np
from faker import Faker

FLIGHT_COUNT = 500000


def random_generator_f():
    pervs = {}

    def wrapper():
        res = uuid4().hex[:6].upper()
        if not pervs.get(res):
            pervs[res] = True
            return res
        return wrapper()

    return wrapper


flight_file = open("flight.csv", "w+", encoding="utf-8", newline="")
flight_writer = csv.writer(flight_file)

flight_writer.writerow(
    [
        "flight_serial",
        "flight_id",
        "origin",
        "destination",
        "aircraft",
        "departure_utc",
        "duration",
        "y_price",
        "j_price",
        "f_price",
    ]
)

airport_df = pd.read_csv("airport.csv", encoding="utf-8")
aircraft_df = pd.read_csv("aircraft.csv", encoding="utf-8")

iata_codes = airport_df["iata_code"]
registrations = aircraft_df["registration"]

random_generator = random_generator_f()

faker = Faker()


for flight_serial in range(1, FLIGHT_COUNT + 1):
    if flight_serial % 10000 == 0:
        print(flight_serial)
    flight_id = random_generator()
    origin, destination = np.random.choice(iata_codes, 2, replace=False)
    aircraft = np.random.choice(registrations)
    departure_utc = faker.future_datetime(tzinfo=pytz.UTC)
    duration = faker.time_delta("+1d")
    y_price = np.random.randint(50, 150)
    j_price = np.random.randint(150, 300)
    f_price = np.random.randint(300, 600)

    flight_writer.writerow(
        [
            flight_serial,
            flight_id,
            origin,
            destination,
            aircraft,
            departure_utc,
            duration,
            y_price,
            j_price,
            f_price,
        ]
    )


flight_file.close()
