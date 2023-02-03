const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

const TICKET_PATH = "./grpc_clients/typeDefs/ticket.proto";
const AUTH_PATH = "./grpc_clients/typeDefs/auth.proto";
const options = require("./typeDefs/options");

const ticketDefinition = protoLoader.loadSync(TICKET_PATH, options);
const authDefinition = protoLoader.loadSync(AUTH_PATH, options);

const ticketProto = grpc.loadPackageDefinition(ticketDefinition);
const authProto = grpc.loadPackageDefinition(ticketDefinition);

const TicketService = ticketProto.TicketService;
const AuthService = authProto.AuthService;

const ticket_grpc = new TicketService('localhost:50050', grpc.credentials.createInsecure());
const auth_grpc = new AuthService('localhost:50051', grpc.credentials.createInsecure());


module.exports = {
    ticket_grpc,
    auth_grpc,
}