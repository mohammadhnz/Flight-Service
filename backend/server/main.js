const express = require("express");
const bodyParser = require("body-parser");
const chalk = require("chalk");

const app = express();

const ticketRoutes = require("./ticketRoutes");

app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());

app.use("/", ticketRoutes);

app.listen(8081, () => {
  const blue = chalk.blue
  const target = blue(`http://localhost:8081`)
  console.log(`🚀 Server ready at ${target}`)
})
