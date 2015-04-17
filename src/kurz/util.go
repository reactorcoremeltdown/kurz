package main

import (
//    "fmt"
    "io/ioutil"
//    "os"
    "encoding/json"
    "log"
)

type Chatroom struct {
    Jid string
    Nick string
}

type Config struct {
    Jid string
    Server string
    Password string
    Status string
    Socket string
    Script string
    Notls bool
    Debug bool
    Logging bool
    LogDirectory string
    AcceptSubscriptionRequests bool
    Chatrooms []Chatroom
}

func LoadConfig(cfgpath string) (c *Config, err error) {
    var bfile []byte
    if bfile, err = ioutil.ReadFile(cfgpath); err != nil {
        log.Fatalf("Error reading config file: %s\n", err.Error())
    }
    c = new(Config)
    err = json.Unmarshal(bfile, &c)
    return
}
