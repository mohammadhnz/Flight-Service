const express = require('express')
const router = express.Router()

const { ticket_grpc, auth_grpc } = require("./grpc_clients");

router.get("/news/", (req, res) => {
  console.log(ticket_grpc.getNews);
  ticket_grpc.getNews({}, (error, response) => {
    if (!error) {
      const { list } = response;
      console.log('[GET] products', response);
      return res.json({ ok: true, list});
    } else {
      console.error(error);
      return res.status(500).json({ ok: false, error});
    }
  });
});

router.get("/tickets/", (req, res) => {
  console.log(ticket_grpc.getNews);
  ticket_grpc.searchTickets({}, (error, response) => {
    if (!error) {
      const { list } = response;
      console.log('search tickets', list.length);
      return res.json({ ok: true, list});
    } else {
      console.error(error);
      return res.status(500).json({ ok: false, error});
    }
  });
});

module.exports = router;
