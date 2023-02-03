import csv
from bs4 import BeautifulSoup
import requests
from geopy.geocoders import Nominatim
from timezonefinder import TimezoneFinder

url = "https://en.wikipedia.org/wiki/List_of_international_airports_by_country"

response = requests.get(url)

soup = BeautifulSoup(response.content, "html.parser")


def normalize_text(text: str):
    s = text
    if text[-1] in ["\n", "\t"]:
        s = text[:-1]

    return s


country_file = open("country.csv", "w+", encoding="utf-8", newline="")
city_file = open("city.csv", "w+", encoding="utf-8", newline="")
airport_file = open("airport.csv", "w+", encoding="utf-8", newline="")

country_list = []
city_list = []
airport_list = []
flight_list = []

country_writer = csv.writer(country_file)
city_writer = csv.writer(city_file)
airport_writer = csv.writer(airport_file)

country_writer.writerow(["country_name"])
city_writer.writerow(
    [
        "country_name",
        "city_name",
        "timezone_name",
    ]
)
airport_writer.writerow(
    [
        "country_name",
        "city_name",
        "airport_name",
        "iata_code",
    ]
)

geolocator = Nominatim(user_agent="geoapiExercises")
timezone_finder = TimezoneFinder()


def get_timezone(country_name: str):
    location = geolocator.geocode(country_name)
    return timezone_finder.timezone_at(lng=location.longitude, lat=location.latitude)


for h4 in soup.find_all("h4"):
    country = normalize_text(h4.find("span", {"class": "mw-headline"}).text)
    print(country)

    country_writer.writerow([country])
    country_list.append([country])

    timezone = get_timezone(country)

    table = h4.find_next_sibling("table")
    trs = table.find_all("tr")[1:]
    tds_len = len(trs[0].find_all("td"))
    prev_location = ""
    for tr in trs:
        tds = tr.find_all("td")
        location = ""
        airport = ""
        iata = ""
        if len(tds) >= tds_len:
            location = normalize_text(tds[0].text)
            prev_location = location
            airport = normalize_text(tds[1].text)
            iata = normalize_text(tds[2].text)
        elif len(tds) == tds_len - 1:
            location = prev_location
            airport = normalize_text(tds[0].text)
            iata = normalize_text(tds[1].text)
        try:
            tz = get_timezone(location)
        except Exception as e:
            tz = timezone
        finally:
            if tz == None:
                print("ridi")
            if not (country, location, tz) in city_list:
                city_writer.writerow([country, location, tz])
                city_list.append((country, location, tz))

            airport_writer.writerow([country, location, airport, iata])
            airport_list.append([country, location, airport, iata])


country_file.close()
city_file.close()
airport_file.close()
