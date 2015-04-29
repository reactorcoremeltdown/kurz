# KURZ 1 2015-04-28

## NAME

kurz - universal xmpp bot

## SYNOPSIS

**kurz** [OPTION]...

## DESCRIPTION

kurz is a simple multipurpose xmpp bot that can handle and log direct messages or multi-user chat.

**-c, --config FILE**
       Path to JSON config file. Default file path is /etc/kurz/default.json

**--help**
       show usage

## CONFIGURATION

kurz aims to be a framework for building xmpp-oriented services. You can run kurz as a service providing different config files. Here is the description of JSON config fields:

**jid**
       Jabber ID of the bot

**server**
       XMPP server in format *fqdn:port*

**password**
       Password of the XMPP account

**status**
       XMPP status message

**notls**
       Ignore TLS encryption. Takes a boolean value(*true* or *false*)

**debug**
       Dump all XMPP traffic(in XML) to stderr for debugging. Takes a boolean value(*true* or *false*)

**socket**
       Path to AF_UNIX socket of the bot. You can send a message to any Jabber ID in the roster of the bot by sending a text to the socket in the format of: *jid*∙*type*∙*text*, where *jid* is a Jabber ID of a recipient, *type* can be *chat* or *groupchat* and *text* is actually the text you want to send to recipient. kurz uses "∙" symbol as the delimiter

**script**
       Path to a script that should handle incoming messages. kurz invokes the script with three arguments: Jabber ID, type of chat and the message text. The script should return processing results via the socket. Please refer to the *LIBRARIES* section for more information

**logging**
       Enable or disable plaintext chat logging. Takes a boolean value(*true* or *false*)

**logDirectory**
       Path to the directory where the plaintext log files will be stored

**acceptSubscriptionRequests**
       Automatically accept and send subscription requests to new recipients. Takes a boolean value(*true* or *false*)

**whitelistEnabled**
       Enable or disable whitelist. Takes a boolean value(*true* or *false*)

**whitelist**
       A list of Jabber IDs who are allowed to invoke the script

**chatrooms**
       A list of chatrooms to join. Each item contains a room JID and a nickname for the bot

## LIBRARIES

There are some inter-process communication libraries which already ship with kurz. They are intended to help communicating with kurz via its socket. You can easily send messages to recipients using them:

**bash**

```
source /usr/lib/kurz/socket_send.sh
socket_send "username@example.org" "chat" "Hello, buddy" "/tmp/example.socket"
```

**node.js**

```
var kurz = require('/usr/lib/kurz/socket_send.js');
kurz.socketSend("username@example.org", "chat", "Hello, buddy", "/tmp/example.socket");
```

## TIPS AND TRICKS

There is a special systemd target unit that ships with kurz. If you want to restart your services when kurz updates, you may place Requires-style dependencies from your services on the target. Here is an example:

```
[Unit]
Description=Example bot
After=kurz.target
Requires=kurz.target

[Service]
User=somebody
ExecStart=/usr/bin/kurz -c /etc/kurz/example.json
```

kurz builds for Debian and Ubuntu already ship with postinstall script which restart *kurz.target* during updates.

## BUGS

Please send your bugreports here: https://github.com/Like-all/kurz/issues

## AUTHOR

Written by Azer Abdullaev aka Like-all

## SEE ALSO

systemd.target(5)
