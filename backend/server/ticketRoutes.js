const express = require('express')
const router = express.Router()

const { ticket_grpc } = require("./grpc_clients");
const authHandler = require("./authHandler");

router.get("/news/", (req, res) => {
  ticket_grpc.getNews({}, (error, response) => {
    if (!error) {
      const { list } = response;
      console.log('[GET] news', response.length);
      return res.json({ ok: true, data: list});
    } else {
      console.error(error);
      return res.status(500).json({ ok: false, error});
    }
  });
});

router.get("/flights/", (req, res) => {
  ticket_grpc.searchFlights(req.query, (error, response) => {
    if (!error) {
      const { list } = response;
      console.log('search tickets', list.length);
      return res.json({ ok: true, data: list});
    } else {
      console.error(error);
      return res.status(500).json({ ok: false, error});
    }
  });
});


router.get("/suggest_origin_destination/", (req, res) => {
  ticket_grpc.suggestOriginDestination(req.query, (error, response) => {
    if (!error) {
      const { list } = response;
      console.log('suggest_origin_destination: ', list.length);
      return res.json({ ok: true, data: list});
    } else {
      console.error(error);
      return res.status(500).json({ ok: false, error});
    }
  });
});

router.post("/ticket/buy", authHandler, (req, res) => {
  ticket_grpc.buyTicket({
    user_id: req.user.id,
    flight_id: req.body.flight_id,
    class_name: req.body.class_name,
    passengers: req.body.passengers,
  }, (error, response) => {
    if (!error) {
      const ticket = response;
      return res.json({ ok: true, data: ticket});
    } else {
      console.error(error);
      return res.status(500).json({ ok: false, error});
    }
  });
});

module.exports = router;
