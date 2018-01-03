exports.sendMessage = function(chat_jid, chat_type, message_text, kurz_socket) {
	var net = require('net');
	var action = {
		actionType: "SendMessage",
		actionSettings: {
			Remote: chat_jid,
			Type: chat_type,
			Text: message_text
		}
	}
	var client = net.connect({path: kurz_socket},
		function() {
			client.write(JSON.stringify(action));
		});
}
