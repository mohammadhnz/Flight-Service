const express = require('express')
const router = express.Router()

const { ticket_grpc, auth_grpc } = require("./grpc_clients");

router.get("/news/", (req, res) => {
  console.log(ticket_grpc);
  ticket_grpc.getNews({}, (error, response) => {
    // const userId = req.headers.hasOwnProperty("x-user-id");
    if (!error) {
      const { news } = response;
      console.log('[GET] products');
      return res.json({ payload: news });
    } else {
      console.error(error);
      return res.json({ payload: null });
    }
  });
});

module.exports = router;
