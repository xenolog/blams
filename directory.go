package blams

import (
	"bufio"
	"fmt"
	//"os"
	//"runtime"
)

type discoverer struct {
	id string
}

type node struct {
	//id     string,
	name string
}

func (s *Discoverer) UseMcast(if_name string, port uint16, mcast_group string) (id string, err error) {
	// add mcast sender/listener to Discoverer
	return id, nil
}

type Directory interface {
	UseMcast(if_name string, port uint16, mcast_group string)
	UseBcast(if_name string, port uint16)
	Use(if_name string, port uint16)
}
