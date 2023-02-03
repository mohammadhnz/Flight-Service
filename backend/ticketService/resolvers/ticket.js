const grpc = require('grpc');
const db = require('../data-access/db');
const { GetProduct, } = require("../classes/Ticket");

function buildQuery({ origin, destination, departure_time , return_time , number_of_passengers}) {
    let params = [];
    let sql = 'SELECT * FROM available_offers where true'
    if(origin){
        params.push(origin);
        sql += ' and origin = $' + params.length.toString();
    }
    if(destination){
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
    console.log(sql, params);
    return {
        text: sql,
        values: params,
    };
}

function splitTickets(request, row) {
    const results = [];
    const flight_classes = ['y', 'j', 'f'];
    const class_names = {'y': "Business", 'j': "Economy", 'f': "First Class"};
    for(const c of Object.values(flight_classes)) {
        if(row[c+'_class_free_capacity'] < request.number_of_passengers){
            continue;
        }
        results.push({
            flight_id: row.flight_id,
            origin: row.origin,
            destination: row.destination,
            duration: row.duration,
            equipment: row.equipment,
            class_name: class_names[c], 
            price: row[c+'_price'], 
            free_capacity: row[c+'_class_free_capacity'], 
            is_limited_capacity: (row[c+'_class_free_capacity'] <= 3 * request.number_of_passengers), 
            departure_local_time: row.departure_local_time.getTime(),
            arrival_local_time: row.arrival_local_time.getTime(),
        });
    }
    return results;
}

const searchTickets = async ({request}, callback) => {
    try {
        const query = buildQuery(request);
        const { rows } = await db.query(query);
        const result = rows.reduce((re, r) => re.concat(splitTickets(request, r)), [])
        if (rows.length !== 0) {
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
        title: "salam",
        imageUrl: "salam",
        redirectUrl: "FEOHOIEF",
   }]}); 
};

module.exports = {
    searchTickets,
    getNews,
};

/*

// grpc.status
// {
//     OK: 0,
//     CANCELLED: 1,
//     UNKNOWN: 2,
//     INVALID_ARGUMENT: 3,
//     DEADLINE_EXCEEDED: 4,
//     NOT_FOUND: 5,
//     ALREADY_EXISTS: 6,
//     PERMISSION_DENIED: 7,
//     RESOURCE_EXHAUSTED: 8,
//     FAILED_PRECONDITION: 9,
//     ABORTED: 10,
//     OUT_OF_RANGE: 11,
//     UNIMPLEMENTED: 12,
//     INTERNAL: 13,
//     UNAVAILABLE: 14,
//     DATA_LOSS: 15,
//     UNAUTHENTICATED: 16
// }

const create = async ({ request }, callback) => {
    const id = require('crypto').randomBytes(10).toString('hex');
    const sql = "INSERT INTO products(id, price_in_cents, title, description, discount.pct, discount.value_in_cents) VALUES($1, $2, $3, $4, $5, $6)";

    const { price_in_cents, title, description, pct } = request;
    const value_in_cents = Math.ceil(price_in_cents * pct); // Type match for int
    const query = {
        text: sql,
        values: [id, price_in_cents, title, description, pct, value_in_cents],
    };

    console.log(request);

    try {
        const { rowCount } = await db.query(query);

        if (rowCount === 1){
            console.log(`Create ${rowCount} product with id(${id}).`);
            callback(null, {
                id,
                price_in_cents,
                title,
                description,
                discount: {
                    pct,
                    value_in_cents,
                }
            });
        } else {
            callback({
                code: grpc.status.CANCELLED,
                details: "CANCELLED",
            })

            // This goes to catch part
            // callback({
            //     code: grpc.status.ALREADY_EXISTS,
            //     details: "ALREADY EXISTS",
            // })
        }
    } catch (e) {
        console.log(e);
        callback(e);
    }
};

const update = async ({ request }, callback) => {
    const { id, price_in_cents, title, description, pct } = request;
    const sql = "UPDATE products SET price_in_cents = $1, title = $2, description = $3, discount.pct = $4, discount.value_in_cents = $5 WHERE id = $6"

    const value_in_cents = Math.ceil(price_in_cents * pct);

    const query = {
        text: sql,
        values: [price_in_cents, title, description, pct, value_in_cents, id],
    };

    try {
        const { rowCount } = await db.query(query);

        if (rowCount === 1) {
            console.log(`Update ${rowCount} product with id(${id}).`);
            callback(null, {
                id,
                price_in_cents,
                title,
                description,
                discount: {
                    pct,
                    value_in_cents,
                }
            });
        } else {
            callback({
                code: grpc.status.NOT_FOUND,
                details: "Not Found"
            });
        }
    } catch (e) {
        console.log(e);
        callback(e);
    }
};

const getProduct = async ({ request }, callback) => {
    const { id } = request;

    const sql = 'SELECT * FROM products WHERE id = $1';
    const query = {
        text: sql,
        values: [id],
    };

    try {
        const { rows } = await db.query(query);

        if (rows.length !== 0) {
            console.log("\n[GET] Product\n")
            const product = rows[0]
            console.log(product)
            const payload = new GetProduct(id, product);

            callback(null, payload );
        }
        else {
            callback({
                code: grpc.status.NOT_FOUND,
                details: "Not Found"
            })
        }
    } catch (e) {
        console.log(e);
        callback(e);
    }
};

const deleteProduct = async ({ request }, callback) => {
    const { id } = request;
    const sql = 'DELETE FROM products WHERE id = $1';
    const query = {
        text: sql,
        values: [id],
    };

    try {
        const { rowCount } = await db.query(query);
        if (rowCount === 1) {
            console.log(`Remove ${rowCount} product with id(${id}).`);
            callback(null, {});
        } else {
            console.log(`There is no product with id(${id}) in database.`);
            callback({
                code: grpc.status.NOT_FOUND,
                details: "Not Found"
            });
        }
    } catch (e) {
        console.log(e);
        callback(e);
    }
};

const deleteProducts = async (_, callback) => {
    const sql = 'DELETE FROM products';
    const query = {
        text: sql,
    };

    try {
        const { rowCount } = await db.query(query);
        if (rowCount === 0) {
            callback({
                code: grpc.status.NOT_FOUND,
                details: "Not Found"
            })
        } else {
            console.log(`Remove ${rowCount} products`);
            callback(null, {})
        }
    } catch (e) {
        console.log(e);
        callback(e);
    }
};

module.exports = {
    getProducts,
    getProduct,
    create,
    update,
    deleteProduct,
    deleteProducts,
}*/