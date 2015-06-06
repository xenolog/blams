package blams

import (
    "bufio"
    "fmt"
    //"os"
    //"runtime"
)

type discoverer struct {
    id  string
    key string
}

type node struct {
    //id     string,
    name   string
    ipaddr map[string]string
}

func (s *discoverer) UseMcast(if_name string, port uint16, mcast_group string) (id string, err error) {
    // add mcast sender/listener to Discoverer
    return id, nil
}

type Directory interface {
    UseMcast(if_name string, port uint16, mcast_group string)
    //UseBcast(if_name string, port uint16)
    //UseDirect(ipaddr string, port uint16)
}

func NewDirectory() Directory {
    return new(discoverer)
}
