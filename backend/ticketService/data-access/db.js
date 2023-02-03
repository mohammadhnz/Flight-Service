// require('dotenv').config({ path: ".env" }) ;

const { Client } = require("pg");

const client = new Client('localhost:5432');
client.connect();

module.exports = {
    query: (text, params, callback) => {
        return client.query(text, params, callback)
    },
};

