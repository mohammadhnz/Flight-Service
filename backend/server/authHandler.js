const {auth_grpc} = require("./grpc_clients");


module.exports = (req, res, next) => {
    if (!req.cookies.tokens) {
        return res.status(401).json({ok: false, data: req.cookies})
    }
    let tokens = JSON.parse(req.cookies.tokens)
    auth_grpc.UserInfo({
        accessToken: tokens['accessToken'],
        refreshToken: tokens['refreshToken'],
    }, (error, response) => {
        if (!error) {
            req.user = response.user;
            next();
        } else {
            console.error(error);
            return res.sendStatus(401);
        }
    });
}
