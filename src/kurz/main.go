package main

import (
    "crypto/tls"
    "fmt"
    "os"
    "log"
    "os/exec"
    "os/signal"
    "syscall"
    "strings"
    "time"
    "net"
    "github.com/Like-all/go-xmpp"
    goopt "github.com/droundy/goopt"
)

var param_cfgpath = goopt.String([]string{"-c", "--config"}, "/etc/kurz/default.json", "set config file path")

func serverName(host string) string {
    return strings.Split(host, ":")[0]
}

func jidInWhitelist(jid string, whitelist []string) bool {
    present := false
    for _, item := range whitelist {
        if jid == item {
            present = true
        }
    }
    return present
}

func main() {
    goopt.Description = func() string {
        return "Kurz - universal xmpp bot"
    }

    goopt.Version = "0.1"
    goopt.Summary = "kurz -c [config]"
    goopt.Parse(nil)

    CfgParams, _ := LoadConfig(*param_cfgpath)

    msgbus := make(chan string)

    if !CfgParams.Notls {
        xmpp.DefaultConfig = tls.Config{
            ServerName:         serverName(CfgParams.Server),
            InsecureSkipVerify: false,
        }
    }

    var talk *xmpp.Client
    var err error
    options := xmpp.Options {
        Host:       CfgParams.Server,
        User:       CfgParams.Jid,
        Password:   CfgParams.Password,
        NoTLS:      CfgParams.Notls,
        Debug:      CfgParams.Debug,
        Session:    false,
        Status:     "chat",
        StatusMessage: CfgParams.Status,
    }

    talk, err = options.NewClient()

    if err != nil {
        log.Fatalf("Error at connection: %s\n", err.Error())
    }

    for _, chatroom :=range CfgParams.Chatrooms {
        talk.JoinMUC(chatroom.Jid, chatroom.Nick)
    }

    go func() {
        for {
            chat, err := talk.Recv()
            if err != nil {
                log.Fatalf("Error at: %s\n", err.Error())
            }
            switch v := chat.(type) {
                case xmpp.Chat:
                    if v.Type == "groupchat" {
                        from := strings.Split(v.Remote, "/")
                        nick := ""
                        if len(from) == 2 {
                            nick = from[1]
                        } else {
                            nick = from[0]
                        }
                        if CfgParams.Logging {
                            logFilename := CfgParams.LogDirectory + "/" + from[0] + ".log"
                            file, err := os.OpenFile(logFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
                            if err != nil {
                                log.Fatalf("Error at: %s\n", err.Error())
                            }
                            _, err = file.WriteString("[" + time.Now().Format("2006-01-02T15:04:05-07:00") + "] <" + nick + "> " + v.Text + "\n")
                            if err != nil {
                                log.Fatalf("Error at: %s\n", err.Error())
                            }
                            file.Close()
                        }
                    } else {
                        from := strings.Split(v.Remote, "/")
                        if CfgParams.Logging {
                            logFilename := CfgParams.LogDirectory + "/" + from[0] + ".log"
                            file, err := os.OpenFile(logFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
                            if err != nil {
                                log.Fatalf("Error at: %s\n", err.Error())
                            }
                            _, err = file.WriteString("[" + time.Now().Format("2006-01-02T15:04:05-07:00") + "] <" + from[0] + "> " + v.Text + "\n")
                            if err != nil {
                                log.Fatalf("Error at: %s\n", err.Error())
                            }
                            file.Close()
                        }
                        if !CfgParams.WhitelistEnabled || jidInWhitelist(from[0], CfgParams.Whitelist) && v.Text != "" {
                            cmd := exec.Command(CfgParams.Script, v.Remote, v.Type, v.Text)
                            err := cmd.Start()
                            if err != nil {
                                log.Fatalf("Error at: %s\n", err.Error())
                            }
                        }
                    }
                case xmpp.Presence:
                    if v.Type == "subscribe" && CfgParams.AcceptSubscriptionRequests {
                        talk.ApproveSubscription(v.From)
                        talk.RequestSubscription(v.From)
                    }
            }
        }
    }()

    go func() {
        for {
            for _, chatroom := range CfgParams.Chatrooms {
                talk.PingC2S(CfgParams.Jid, chatroom.Jid + "/" + chatroom.Nick)
            }
            time.Sleep(5 * time.Second)
        }
    }()

    go func() {
        l, err := net.ListenUnix("unix", &net.UnixAddr{CfgParams.Socket, "unix"})
        if err != nil {
            log.Fatalf("Error at: %s\n", err.Error())
        }
        for {
            conn, err := l.AcceptUnix()
            if err != nil {
                log.Fatalf("Error at: %s\n", err.Error())
            }
            var buf [1024]byte
            n, err := conn.Read(buf[:])
            if err != nil {
                log.Fatalf("Error at: %s\n", err.Error())
            }
            msgbus <- string(buf[:n])
        }
    }()

    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGINT,
        syscall.SIGTERM,
        syscall.SIGQUIT)
    go func() {
        for sig := range c {
            os.Remove(CfgParams.Socket)
            fmt.Printf("Captured %v, Exiting\n", sig)
            os.Exit(0)
        }
    }()

    for {
        msg := <-msgbus
        msgparts := strings.Split(msg, "∙")
        talk.Send(xmpp.Chat{Remote: msgparts[0], Type: msgparts[1], Text: msgparts[2]})
        if CfgParams.Logging {
            logFilename := CfgParams.LogDirectory + "/" + strings.Split(msgparts[0], "/")[0] + ".log"
            file, err := os.OpenFile(logFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
            if err != nil {
                log.Fatalf("Error at: %s\n", err.Error())
            }
            _, err = file.WriteString("[" + time.Now().Format("2006-01-02T15:04:05-07:00") + "] <" + CfgParams.Jid + "> " + msgparts[2] + "\n")
            if err != nil {
                log.Fatalf("Error at: %s\n", err.Error())
            }
            file.Close()
        }
    }
}
