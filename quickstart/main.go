package main

import (
	"fmt"
	"github.com/shura1014/balance"
)

type Server struct {
	addr   string
	valid  bool
	weight int
}

func (s *Server) IsValid() bool {
	return s.valid
}

func (s *Server) GetWeight() int {
	return s.weight
}

func main() {
	//NoInstance()
	//Usable()
	//RoundRobin()
	//Random()
	Weight()
}

func NoInstance() {
	b := balance.GetBalance[*Server](balance.RoundRobin_)
	b.InitNodes(&Server{addr: "127.0.0.1:8081"}, &Server{addr: "127.0.0.1:8082"})
	for i := 0; i < 3; i++ {
		instance, err := b.Instance()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(instance.addr)
		}
	}
}

// Usable 部分可用情况
func Usable() {
	b := balance.GetBalance[*Server](balance.RoundRobin_)
	b.InitNodes(&Server{addr: "127.0.0.1:8081", valid: true}, &Server{addr: "127.0.0.1:8082"})
	for i := 0; i < 3; i++ {
		instance, err := b.Instance()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(instance.addr)
		}
	}
}

func RoundRobin() {
	b := balance.GetBalance[*Server](balance.RoundRobin_)
	b.InitNodes(&Server{addr: "127.0.0.1:8081", valid: true}, &Server{addr: "127.0.0.1:8082", valid: true})
	for i := 0; i < 3; i++ {
		instance, err := b.Instance()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(instance.addr)
		}
	}
}

func Random() {
	b := balance.GetBalance[*Server](balance.Random_)
	b.InitNodes(&Server{addr: "127.0.0.1:8081", valid: true}, &Server{addr: "127.0.0.1:8082", valid: true})
	for i := 0; i < 6; i++ {
		instance, err := b.Instance()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(instance.addr)
		}
	}
}

func Weight() {
	b := balance.NewWeight[*Server]()
	b.InitNodes(&Server{addr: "127.0.0.1:8081", valid: true, weight: 2}, &Server{addr: "127.0.0.1:8082", valid: true, weight: 4})
	for i := 0; i < 15; i++ {
		instance, err := b.Instance()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(instance.addr)
		}
	}
}
