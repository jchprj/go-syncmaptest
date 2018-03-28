package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"golang.org/x/sync/syncmap"
)

//Client ... Client
type Client struct {
	ID     string
	Server string
}

var clients map[string]Client
var m *syncmap.Map

const test = true

func main() {
	fmt.Println("dd")
	clients = make(map[string]Client)

	m = new(syncmap.Map)
	fmt.Println(m)

	for j := 0; j <= 10000; j++ {
		// fmt.Println("j", j)
		go write()
		go read()
	}
}

func write() {
	client := Client{"ss", "server"}
	i := rand.Intn(88)
	fmt.Println("write", i)
	istr := strconv.Itoa(i)
	if test {
		m.Store(istr, client)
	} else {
		clients[istr] = client
	}
}
func read() {
	i := rand.Intn(88)
	istr := strconv.Itoa(i)
	var client Client
	var ok bool
	if test {
		tmpclient, tmpok := m.Load(istr)
		ok = tmpok
		if ok {
			client = tmpclient.(Client)
		}
	} else {
		tmpclient, tmpok := clients[istr]
		ok = tmpok
		if ok {
			client = tmpclient
		}
	}
	if ok {
		fmt.Println("read", i, client.Server)
	} else {
		fmt.Println("read", i)
	}

}
