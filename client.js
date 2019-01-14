var net = require('net');

var server = net.createServer(function(socket) {
	socket.on('data', function(data) {
		socket.write('Echo server\r\n');
		console.log('Received: ' + data);
	});
});

server.listen(4545);


