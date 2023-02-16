const { Client } = require("pg");
const redis = require('redis');

const postgresClient = new Client({
    user: 'user1',
    database: 'project1',
    password:'password1',
    port: 5432,
    host: 'postgres',
    ssl: false
});

const redisClient = redis.createClient({url: 'redis://redis:6379'});

redisClient.on('connect', function() {
    console.log('Connected to redis!');
});

redisClient.on('error', function(e) {
    console.log('error to redis!');
    console.log(e);
});

postgresClient.connect();
redisClient.connect();

postgresClient.on('error', (e)=>{
    console.log(e);
});

function createKey(text, values) {
    return text+"_$$$_"+values.reduce((line, value) => line + "_"+ value, "");
}

const Indexes = {
    AvailableOffersByFlightId: "SELECT * FROM available_offers where flight_id = $1",
    FlightByFlightSerial:  "SELECT * FROM flight where flight_serial = $1",
    FlightByFlightId:  "SELECT * FROM flight where flight_id = $1",
    SuggestOriginDestination: 'select * from origin_destination '
                                + 'where city ILIKE CONCAT(\'%\', cast($1 as text), \'%\') '
                                + 'or county ILIKE CONCAT(\'%\', cast($1 as text), \'%\') '
                                + 'or airport ILIKE CONCAT(\'%\', cast($1 as text), \'%\') '
                                + 'or iata ILIKE CONCAT(\'%\', cast($1 as text), \'%\') '
                                + 'limit 10',
    PurchaseByTrackingCode:  'select * from purchase where tracking_code = $1',
    PurchaseByUserId: 'select * from purchase where corresponding_user_id = $1',
    
}

async function removeCache(index, values) {
    try {
        if(values){
            const queryKey = createKey(index, values);
            console.log("REMOVING", queryKey);
            await redisClient.del(queryKey);
        }else{
            const keys = await redisClient.keys(index + "_$$$_*");
            for(const key of keys) {
                await redisClient.del(key);
            }
        }
    } catch (Error) {
        return;
    }
}

async function redis_get(index, values) {
    try {
        return redisClient.get(queryKey);
    } catch (Error) {
        return null;
    }
}


const functions = {
    update: async ({text, values, dirtyIndexes}) => {
        const result = await postgresClient.query(text, values);
        for(const {index, values} of dirtyIndexes) {
            await removeCache(index, values);
        }
        return result;
    },
    insert: async ({text, values, dirtyIndexes}) => {
        const result = await postgresClient.query(text, values);
        for(const {index, values} of dirtyIndexes) {
            await removeCache(index, values);
        }
        return result;
    },
    queryIndex: async (index, ...values) => {
        const queryKey = createKey(index, values);
        const cacheResult = await redisClient.get(queryKey);
        if (cacheResult) {
            console.log('Cache hit!');
            return {rows: JSON.parse(cacheResult)};
        }
        console.log('Cache miss!');
        console.log(values, Indexes[index]);
        const result = await postgresClient.query(Indexes[index], values);
        console.log(JSON.stringify(result.rows));
        await redisClient.set(queryKey, JSON.stringify(result.rows), 'EX', 60 * 60);
        return result;
    },
    complexSelect: async ({text, values}) => {
        const queryKey = createKey(text, values);
        const cacheResult = await redisClient.get(queryKey);
        if (cacheResult) {
            console.log('Cache hit!');
            return {rows: JSON.parse(cacheResult)};
        }
        console.log('Cache miss!');
        const result = await postgresClient.query(text, values);
        await redisClient.set(queryKey, JSON.stringify(result.rows), 'EX', 60 * 60);
        return result;
    },
};

module.exports = functions;

