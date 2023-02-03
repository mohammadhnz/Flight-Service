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

router.get("/tickets/", (req, res) => {
  console.log(req.query, req.params)
  ticket_grpc.searchTickets(req.query, (error, response) => {
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


router.post("/ticket/buy/{id}", authHandler, (req, res) => {
  console.log(req.query, req.params)
  ticket_grpc.searchTickets(req.query, (error, response) => {
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

module.exports = router;
