import random
import string

from locust import HttpUser, task, between


def get_random_string(length):
    letters = string.ascii_lowercase
    return ''.join(random.choice(letters) for _ in range(length))


def get_correct_time():
    correct_time_start = random.randint(946684800000, 2524508000000)
    correct_time_end = random.randint(2523508000000, 2524508000000)
    return correct_time_start, correct_time_end


def get_wrong_time():
    time_start = random.randint(2524508010000, 2524508020000)
    return time_start, random.randint(10000, 300000)


def select_airport():
    if random.randint(1, 10) < 7:
        return random.choice(
            ["ALG", "AAE", "BLJ", "BJA", "BSK", "CFK", "CZL", "HME", "GJL", "ORN", "QSF", "TMR", "TLM", "HBE", "ALY",
             "ATZ", "ASW", "CAI", "AAC", "DBB", "HRG", "LXR", "RMF", "MUH", "SKV", "SSH", "HMB", "TCP", "BEN", "SEB",
             "TIP", "MJI", "AGA", "CMN", "FEZ", "RAK", "NDR", "OUD", "RBA", "TNG", "TTU", "VIL", "EUN", "KRT", "PZU",
             "DJE", "NBE", "MIR", "SFA", "TBJ", "TOE", "TUN", "COO", "BOY", "OUA", "BVC", "SID", "RAI", "VXE", "ABJ",
             "BJL", "ACC", "KMS", "TKD", "NYI", "HZO", "WZA", "TML", "CKY", "OXB", "BQE", "ROB", "BKO", "NKC", "NDB",
             "ATR", "NIM"])
    return 'ALL'


class UserWithoutSearchFlight(HttpUser):
    wait_time = between(1, 2)

    def on_start(self):
        gmail = "user" + get_random_string(12) + "@gmail.com"
        password = get_random_string(16)
        phone = "09" + get_random_string(9)
        self.client.post("/auth/signup/", json={
            "email": gmail,
            "first_name": get_random_string(8),
            "last_name": get_random_string(8),
            "gender": "female",
            "phone_number": phone,
            "password": password
        })

        self.client.post("/auth/signin", json={
            "email": gmail,
            "phone_number": phone,
            "password": password
        })

    @task
    def signin(self):
        self.client.post("/auth/signin", json={
            "email": "user@gmail.com",
            "phone_number": "09919700535",
            "password": "password"
        })

    @task(3)
    def suggest_origin_destination(self):
        self.client.get(f"/suggest_origin_destination", params={'name': get_random_string(2)})

    @task(5)
    def get_news(self):
        self.client.get(f"/news")

    @task(1)
    def but_ticket(self):
        self.client.get(f"/ticket/buy", json={
            "flight_id": "AD3E48",
            "class_name": random.choice(['First Class', 'Economy', 'Business']),
            "passengers": [{
                "name": get_random_string(8),
                "family": get_random_string(8),
                "passport": get_random_string(8)
            }]
        })

    @task(2)
    def get_ticket_list(self):
        self.client.get(f"/ticket/list")


class NormalUser(UserWithoutSearchFlight):
    @task(2)
    def get_flights(self):
        if random.randint(1, 10) < 11:
            start, end = get_correct_time()
        else:
            start, end = get_wrong_time()
        origin = select_airport()
        destination = select_airport()
        self.client.get(f"/flights", params={'origin': origin, 'destination': destination,
                                             'number_of_passengers': random.randint(1, 10),
                                             'departure_time': start, 'return_time': end})
