var PROTO_PATH = __dirname + '/../sparkgrpc/commands.proto';

var grpc = require('grpc');
var command_proto = grpc.load(PROTO_PATH).sparkgrpc;

function main() {
  var client = new command_proto.sparkusb('localhost:8001',
                                       grpc.credentials.createInsecure());
/*
client.connect({device:"/dev/ttyACM0"}, function(err, response) {
    console.log('Reponse:', response);
  });
*/



  client.getParameter({value : 5, parameter:0}, function(err, response) {
    console.log('Reponse:', response);
  });



}

main();
