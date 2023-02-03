const { Client } = require("pg");

const client = new Client({
    user: 'user1',
    database: 'project1',
    password:'password1',
    port: 5432,
    host: 'localhost',
    ssl: false
});

client.connect();
client.on('error', (e)=>{
    console.log(e);
})

module.exports = {
    query: (text, params, callback) => {
        return client.query(text, params, callback);
    },
};

