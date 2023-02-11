const { AssertionError } = require('assert');
const assert = require('assert');

module.exports = {
    // default after: 2000:1:1
    // default before: 2050:1:1
    validateTimestamp(time, fieldName, messagePrefix="invalid timestamp ", before=2524608000000, after=946684800000) {
        assert(time<=before && time>=after, messagePrefix+fieldName);
    },
 
    validateStringLength(str, fieldName, l, r) {
        assert(str.length>=l && str.length<=r, "value "+ fieldName + "has invalid length");
    },
 
    validateNumberRange(x, fieldName, l, r) {
        assert(x>=l && x<=r, "value " + fieldName + " is not in range [" + l + ", " + r + "]");
    },

    validateIdentifier(x, fieldName) {
        assert(x>=1 && x<=1000000000, "invalid identifier "+fieldName);
    },
 
    validateStringOneOf(x, fieldName, alowedList) {
        assert(alowedList.contains(x), "invalid "+fieldName+": value should be one of {"+alowedList+"}");
    },

    validateSearchParams(searchParams) {
        if(searchParams.origin) {
            this.validateStringLength(searchParams.origin, 'origin', 2, 30);
        }
        if(searchParams.destination) {
            this.validateStringLength(searchParams.destination, 'destination', 2, 30);
        }
        if(searchParams.return_time) {
            this.validateTimestamp(searchParams.return_time, 'return_time');
        }
        if(searchParams.departure_time) {
            this.validateTimestamp(searchParams.departure_time, 'departure_time');
        }
        if(searchParams.number_of_passengers) {
            this.validateNumberRange(searchParams.number_of_passengers, 'number_of_passengers', 1, 30);
        }
    },

    validatePassengerData(passenger) {
        validators.validateStringLength(passenger.name, 'passenger name', 2, 30);
        validators.validateStringLength(passenger.name, 'passenger family', 2, 30);
        validators.validateStringLength(passenger.name, 'passenger passport', 2, 30);
    },

    validateCreateTicketRequest(body) {
        validators.validateIdentifier(req.body.flight_id, 'flight_id');
        validators.validateStringOneOf(req.body.class_name, ['First Class', 'Economy', 'Business']);
        assert(Array.isArray(body.passenger), 'passengers field must be an array!');
        assert(body.passenger.length>=1 && body.passenger.length<30, 'passengers field has illigal length');
        body.passenger.forEach(passenger => this.validatePassengerData(passenger));
    },

    validatorMiddleware(err, req, res, next) {
        if (err instanceof AssertionError) {
            return res.status(400).json({ ok: false, error: err.message});
        }
        next();
    }
};