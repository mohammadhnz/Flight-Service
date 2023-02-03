// const { auth_grpc } = require("./grpc_clients");

const auth_grpc = {
  authenticate: (data, callback) => {
    // return callback(null, {success: false, code: 403});
    return callback("error");
    // return callback(null, {success: true, user: {id: 34}});
  },
}

module.exports = (req, res, next) => {
  auth_grpc.authenticate({path: req.path, headers: req.headers}, (error, response) => {
    if (!error && response.success) {
      req.user = response.user;
      next();
    } else {
      return res.send(error ? 500: response.code);
    }
  });
}
