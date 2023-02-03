const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');
const chalk = require("chalk");

const ticketResolvers = require("./resolvers/ticket");

const PROTO_PATH = "./typeDefs/ticket.proto";
const options = require("./typeDefs/options");

const packageDefinition = protoLoader.loadSync(PROTO_PATH, options);
const productproto = grpc.loadPackageDefinition(packageDefinition);

const main = () => {
    const server = new grpc.Server()
    console.log(productproto.TicketService.service);

    server.addService(productproto.TicketService.service, ticketResolvers);
    const port = "0.0.0.0:50050";

    server.bind(port, grpc.ServerCredentials.createInsecure());
    const blue = chalk.blue
    const target = blue(`http://${port}`)

    console.log(`🚀 gRPC ticket server ready at ${target}`);
    server.start();
}

main();