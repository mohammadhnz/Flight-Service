const grpc = require('grpc');
const db = require('../data_access/db');
const request = require('request');
const uuid = require('uuid').v4;


function buildQuery({ origin, destination, departure_time , return_time , number_of_passengers}) {
    let params = [];
    let sql = 'SELECT * FROM available_offers where true'
    if(origin && origin.toLowerCase()!="all"){
        params.push(origin);
        sql += ' and origin = $' + params.length.toString();
    }
    if(destination && destination.toLowerCase()!="all"){
        params.push(destination);
        sql += ' and destination = $' + params.length.toString();
    }
    if(departure_time){
        params.push(new Date(parseInt(departure_time)).toISOString());
        sql += ' and date_trunc(\'day\', departure_local_time) = date_trunc(\'day\', cast($'+params.length.toString()+' as timestamp))';
    }
    if(return_time){
        params.push(new Date(parseInt(return_time)).toISOString());
        sql += ' and date_trunc(\'day\', arrival_local_time) = date_trunc(\'day\', cast($'+params.length.toString()+' as timestamp))';
    }
    if(number_of_passengers){
        params.push(number_of_passengers);
        sql += ' and (y_class_free_capacity > $' + params.length.toString();
        sql += ' or j_class_free_capacity > $' + params.length.toString();
        sql += ' or f_class_free_capacity > $' + params.length.toString() + ')';
    }
    sql += ' limit 30';
    return {
        text: sql,
        values: params,
    };
}

function getFlightClassData(row, class_name, number_of_passengers){
    const class_names = {"Business": 'y', 'Economy': "j", 'First Class': "f"};
    const c = class_names[class_name];
    return {
        flight_id: row.flight_id,
        flight_serial: row.flight_serial,
        origin: row.origin,
        destination: row.destination,
        duration: row.duration,
        equipment: row.equipment,
        class_name: class_name, 
        price: row[c+'_price'], 
        free_capacity: row[c+'_class_free_capacity'], 
        is_limited_capacity: (row[c+'_class_free_capacity'] <= 3 * number_of_passengers), 
        departure_local_time: new Date(row.departure_local_time).getTime(),
        arrival_local_time: new Date(row.arrival_local_time).getTime(),
    };
}

async function getFlightData(flight_id, class_name, number_of_passengers) {
    const { rows } = await db.queryIndex('AvailableOffersByFlightId', flight_id);
    if (rows.length === 0){
        return null;
    }
    return getFlightClassData(rows[0], class_name, number_of_passengers);
}

async function getFlightDataWithSerial(flight_serial, class_name) {
    const { rows } = await db.queryIndex('FlightByFlightSerial', flight_serial);
    if (rows.length === 0){
        return null;
    }
    return getFlightData(rows[0].flight_id, class_name, 1);
}


async function getFlightSerial(flight_id) {
    const { rows } = await db.queryIndex('FlightByFlightId', flight_id);
    if (rows.length === 0){
        return null;
    }
    return rows[0].flight_serial;
}

function getFullFlightData(request, row) {
    return {
        flight_id: row.flight_id,
        flight_serial: row.flight_serial,
        origin: row.origin,
        destination: row.destination,
        duration: row.duration,
        equipment: row.equipment,
        y_price: row.y_price, 
        j_price: row.j_price, 
        f_price: row.f_price, 
        y_class_free_capacity: row.y_class_free_capacity, 
        j_class_free_capacity: row.j_class_free_capacity, 
        f_class_free_capacity: row.f_class_free_capacity, 
        y_is_limited_capacity: (row.y_class_free_capacity <= 3 * request.number_of_passengers), 
        j_is_limited_capacity: (row.j_class_free_capacity <= 3 * request.number_of_passengers), 
        f_is_limited_capacity: (row.f_class_free_capacity <= 3 * request.number_of_passengers), 
        departure_local_time: new Date(row.departure_local_time).getTime(),
        arrival_local_time: new Date(row.arrival_local_time).getTime(),
    };
}

const searchFlights = async ({request}, callback) => {
    try {
        const query = buildQuery(request);
        const { rows } = await db.complexSelect(query);
        if (rows.length !== 0) {
            const result = rows.map(r => getFullFlightData(request, r));
            callback(null, { list: result}); 
        }
        else {
            callback({
                code: grpc.status.NOT_FOUND,
                details: "Not Found"
            })
        }
    } catch(e) {
        console.log(e);
        callback(e);
    }
};

const getNews = async (_, callback) => {
    callback(null,  { list: [ {
        title: "?????? ??",
        image_url: "http://localhost:8081/image/news1.png",
        redirect_url: "http://localhost:7777/blog/1",
    },{
        title: "?????? ??",
        image_url: "http://localhost:8081/image/news2.png",
        redirect_url: "http://localhost:7777/blog/2",
    },{
        title: "?????? ??",
        image_url: "http://localhost:8081/image/news3.png",
        redirect_url: "http://localhost:7777/blog/3",
    },{
        title: "?????? ??",
        image_url: "http://localhost:8081/image/news4.png",
        redirect_url: "http://localhost:7777/blog/4",
    }]}); 
};


const suggestOriginDestination = async ({request: {name}}, callback) => {
    try {
        const { rows } = await db.queryIndex('SuggestOriginDestination', name);
        if (rows.length !== 0) {
            callback(null, { list: rows}); 
        }
        else {
            callback({
                code: grpc.status.NOT_FOUND,
                details: "Not Found"
            })
        }
    } catch(e) {
        console.log(e);
        callback(e);
    }
};

function getPurchaseTitle(flight) {
    return "???????????? ?????????? ???????? ?????????? ?????????? "+flight.flight_serial+", ??????????: "+flight.price;
}

function makeTransaction(amount, receipt_id) {
    const tracking_code = uuid();
    return new Promise(function (resolve, reject) {
      request.post('http://bank:8080/transaction/', { 
        json: { 
            amount,
            receipt_id,
            callback: 'http://localhost:8081/payment/callback/' + tracking_code
        }
    }, function (error, res, body) {
        if (!error && res.statusCode === 201) {
          resolve({...body, tracking_code});
        } else {
          reject(error);
        }
      });
    });
}

async function getTicketData(row) {
    const flight = await getFlightDataWithSerial(row.flight_serial, row.offer_class);
    return {
        transaction_id: row.transaction_id, 
        flight, 
        passengers: [{
            name: row.first_name,
            family: row.last_name,
        }],
        payment_url: row.transaction_result === 1 ? undefined : "http://localhost:8080/payment/" + row.transaction_id + "/",
    };
}

async function getTicket(tracking_code) {
    const {rows} = await db.queryIndex('PurchaseByTrackingCode', tracking_code);
    if(rows.length===0){
        return null;
    }
    return getTicketData(rows[0]);
}

const createTicket = async ({request: {user_id, flight_id, class_name, passengers}}, callback) => {
    try {
        const flight = await getFlightData(flight_id, class_name, passengers.length);
        if (flight === null){
            return callback({
                code: grpc.status.NOT_FOUND,
                details: "Not Found"
            });
        }
        const flight_serial = await getFlightSerial(flight.flight_id);

        const {id: transaction_id, tracking_code} = await makeTransaction(flight.price, flight_serial);
        const query = {
            text: 'insert into purchase ('
                + 'corresponding_user_id,'
                + 'title,'
                + 'first_name,'
                + 'last_name,'
                + 'flight_serial,'
                + 'offer_price,'
                + 'offer_class,'
                + 'transaction_id,'
                + 'tracking_code'
            + ') values ($1, $2, $3, $4, $5, $6, $7, $8, $9)',
            values: [
                user_id,
                getPurchaseTitle(flight), 
                passengers[0].name, 
                passengers[0].family, 
                flight_serial, 
                flight.price, 
                flight.class_name,
                transaction_id,
                tracking_code,
            ],
            dirtyIndexes: [{
                index: 'PurchaseByTrackingCode',
                values: [tracking_code],
            },{
                index: 'PurchaseByUserId',
                values: [user_id],
            }]
        };
        const {rowCount} = await db.insert(query);
        if (rowCount === 1) {
            callback(null, await getTicket(tracking_code)); 
        }
        else {
            callback({
                code: grpc.status.NOT_FOUND,
                details: "Not Found"
            })
        }
    } catch(e) {
        console.log(e);
        callback(e);
    }
};


const parchase = async ({request: {tracking_code, status}}, callback) => {
    try {
        const query = {
            text: 'update purchase set transaction_result = $1 where (tracking_code = $2)',
            values: [
                status,
                tracking_code, 
            ],
            dirtyIndexes: [{
                index: 'PurchaseByTrackingCode',
                values: [tracking_code],
            },{
                index: 'PurchaseByUserId'
            }]
        };
        const {rowCount} = await db.update(query);
        if (rowCount === 1) {
            callback(null, await getTicket(tracking_code)); 
        }
        else {
            callback({
                code: grpc.status.NOT_FOUND,
                details: "Not Found"
            })
        }
    } catch(e) {
        console.log(e);
        callback(e);
    }
};


const getUsersTickets = async ({request: {user_id}}, callback) => {
    try {
        const {rows} = await db.queryIndex('PurchaseByUserId', user_id);
        if (rows.length !== 0) {
            const list = [];
            for(const r of rows){
                list.push(await getTicketData(r));
            }
            callback(null, {list}); 
        }
        else {
            callback({
                code: grpc.status.NOT_FOUND,
                details: "Not Found"
            })
        }
    } catch(e) {
        console.log(e);
        callback(e);
    }
};

module.exports = {
    searchFlights,
    getNews,
    suggestOriginDestination,
    createTicket,
    parchase,
    getUsersTickets,
};
