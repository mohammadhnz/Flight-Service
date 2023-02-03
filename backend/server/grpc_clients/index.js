const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

const TICKET_PATH = "./grpc_clients/type_defs/ticket.proto";
const AUTH_PATH = "./grpc_clients/type_defs/auth.proto";
const options = require("./type_defs/options");

const ticketDefinition = protoLoader.loadSync(TICKET_PATH, options);
const authDefinition = protoLoader.loadSync(AUTH_PATH, options);

const ticketProto = grpc.loadPackageDefinition(ticketDefinition);
const authProto = grpc.loadPackageDefinition(authDefinition);

const TicketService = ticketProto.TicketService;
const AuthService = authProto.AuthService;

const ticket_grpc = new TicketService('ticket:50050', grpc.credentials.createInsecure());
const auth_grpc = null; // new AuthService('auth:50051', grpc.credentials.createInsecure());


module.exports = {
    ticket_grpc,
    auth_grpc,
}