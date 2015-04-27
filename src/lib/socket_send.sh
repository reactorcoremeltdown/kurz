function socket_send() {
    echo -n "$1∙$2∙$3" | socat stdio unix-connect:$4
}
