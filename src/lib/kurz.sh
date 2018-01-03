function kurz_send_message() {
  chat_jid=$1
  chat_type=$2
  message_text=$3
  kurz_socket=$4

  echo '{ "actionType": "SendMessage", "actionSettings": { "Remote": '$chat_jid', "Type": '$chat_type', "Text": '$message_text' } }' | socat stdio unix-connect:$4
}
