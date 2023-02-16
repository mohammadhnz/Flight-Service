const express = require('express');
const router = express.Router();

const path = require('path');
const {ticket_grpc, auth_grpc} = require("./grpc_clients");
const authHandler = require("./authHandler");

router.post("/signup/", (req, res) => {
    // console.log(auth_grpc.Signup);
    auth_grpc.Signup({
        email: req.body.email,
        first_name: req.body.first_name,
        last_name: req.body.last_name,
        gender: req.body.gender == "male" ? 0 : 1,
        phone_number: req.body.phone_number,
        password: req.body.password,
    }, (error, response) => {
        if (!error) {
            const ticket = response;
            return res.json({ok: true, data: ticket});
        } else {
            console.error(error);
            return res.status(500).json({ok: false, time: "ali", error});
        }
    })
});


router.post("/signin/", (req, res) => {
    auth_grpc.SignIn({
        email: req.body.email,
        phone_number: req.body.phone_number,
        password: req.body.password,
    }, (error, response) => {
        if (!error) {
            res.cookie('tokens', JSON.stringify(response))
            return res.cookie('tokens', JSON.stringify(response)).send()
        } else {
            console.error(error);
            return res.status(500).json({ok: false, time: "ali", error});
        }
    })
});


router.post("/signout/", (req, res) => {
    if (!req.cookies.tokens) {
        return res.status(500).json({ok: false, data: req.cookies})
    }
    let tokens = JSON.parse(req.cookies.tokens)
    auth_grpc.SignOut({
        accessToken: tokens['accessToken'],
        refreshToken: tokens['refreshToken'],
    }, (error, response) => {
        if (!error) {
            res.cookie('tokens', JSON.stringify(response))
            return res.cookie('tokens', req.cookies.tokens).status(200).json({ok: true});
        } else {
            console.error(error);
            return res.cookie('tokens', req.cookies.tokens).status(400).json({ok: false, error});
        }
    })


});


router.get("/userinfo/", (req, res) => {
    if (!req.cookies.tokens) {
        return res.status(500).json({ok: false, data: req.cookies})
    }
    let tokens = JSON.parse(req.cookies.tokens)
    auth_grpc.UserInfo({
        accessToken: tokens['accessToken'],
        refreshToken: tokens['refreshToken'],
    }, (error, response) => {
        if (!error) {
            res.cookie('tokens', JSON.stringify(response))
            return res.cookie('tokens', req.cookies.tokens).status(200).json({data: response});
        } else {
            console.error(error);
            return res.cookie('tokens', req.cookies.tokens).status(400).json({ok: false, error});
        }
    })

});

router.get("/refresh/", (req, res) => {
    if (!req.cookies.tokens) {
        return res.status(500).json({ok: false, data: req.cookies})
    }
    let tokens = JSON.parse(req.cookies.tokens)
    auth_grpc.Refresh({
        accessToken: tokens['accessToken'],
        refreshToken: tokens['refreshToken'],
    }, (error, response) => {
        if (!error) {
            return res.cookie('tokens', JSON.stringify({
                accessToken: response["accessToken"],
                refreshToken: response["refreshToken"],
            })).status(200).json({data: response});
        } else {
            console.error(error);
            return res.cookie('tokens', req.cookies.tokens).status(400).json({ok: false, error});
        }
    })

});

module.exports = router;
