const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

const TICKET_PATH = "./grpc_clients/typeDefs/ticket.proto";
const options = require("./typeDefs/options");

const ticketDefinition = protoLoader.loadSync(
    TICKET_PATH,
    options
);

const ticketProto = grpc.loadPackageDefinition(ticketDefinition);

const TicketService = ticketProto.TicketService;


const ticket_grpc = new TicketService('localhost:50050', grpc.credentials.createInsecure());

module.exports = {
    ticket_grpc,
    auth_grpc: null,
}