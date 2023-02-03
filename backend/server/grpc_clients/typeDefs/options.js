// https://www.npmjs.com/package/@grpc/proto-loader

const options = {
    keepCase: true, // important to use true
    longs: String,
    enums: String,
    defaults: false,
    oneofs: true
};

module.exports = options;