exports.socketSend = function(jid, type, text, socket) {
	var net = require('net');
	var client = net.connect({path: socket},
		function() {
			client.write(jid + '∙' + type + '∙' + text);
		});
}
