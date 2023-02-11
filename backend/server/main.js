const express = require("express");
const bodyParser = require("body-parser");
const chalk = require("chalk");
const cookieParser = require("cookie-parser");

const app = express();

const ticketRoutes = require("./ticketRoutes");
const authHandler = require("./authHandler");
const authRoutes = require("./authRoutes");
const {validatorMiddleware} = require("./validators");

app.use(cookieParser());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());
app.use("/auth", authRoutes);
app.use(authHandler);
app.use("/", ticketRoutes);
app.use(validatorMiddleware);


app.listen(8081, () => {
  const blue = chalk.blue
  const target = blue(`http://localhost:8081`)
  console.log(`🚀 Server ready at ${target}`)
})
