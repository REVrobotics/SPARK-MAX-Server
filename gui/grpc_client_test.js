var PROTO_PATH = __dirname + '/../sparkgrpc/commands.proto';

var grpc = require('grpc');
var command_proto = grpc.load(PROTO_PATH).sparkgrpc;

function main() {
  var client = new command_proto.sparkusb('localhost:8001',
                                       grpc.credentials.createInsecure());
  var user;
  if (process.argv.length >= 3) {
    user = process.argv[2];
  } else {
    user = 'world';
  }
  client.List({}, function(err, response) {
    console.log('Reponse:', response.deviceList);
  });
}

main();
