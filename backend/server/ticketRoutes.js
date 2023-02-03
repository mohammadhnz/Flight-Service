const express = require('express')
const router = express.Router()

const { ticket_grpc, auth_grpc } = require("./grpc_clients");

router.get("/news/", (req, res) => {
  console.log(ticket_grpc.getNews);
  ticket_grpc.getNews({}, (error, response) => {
    // const userId = req.headers.hasOwnProperty("x-user-id");
    if (!error) {
      const { list } = response;
      console.log('[GET] products', response);
      return res.json({ payload: list });
    } else {
      console.error(error);
      return res.json({ payload: null });
    }
  });
});

module.exports = router;
