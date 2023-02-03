const express = require('express');
const router = express.Router();

const path = require('path');
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
  ticket_grpc.createTicket({
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

router.get("/payment/callback/:code/:status", (req, res) => {
  ticket_grpc.parchase({
    user_id: req.user.id,
    tracking_code: req.params.code,
    status: req.params.status,
  }, (error, response) => {
    if (!error) {
      const staticFile = req.params.status==1 ? "paymentSuccess.html" : "paymentFailed.html";
      return res.sendFile(path.join(__dirname, "./static/"+staticFile));
    } else {
      console.error(error);
      return res.status(500).json({ ok: false, error});
    }
  });
});

router.get("/ticket/list", authHandler, (req, res) => {
  ticket_grpc.getUsersTickets({
    user_id: req.user.id,
  }, (error, response) => {
    if (!error) {
      const {list} = response;
      return res.json({ ok: true, data: list});
    } else {
      console.error(error);
      return res.status(500).json({ ok: false, error});
    }
  });
});

module.exports = router;
