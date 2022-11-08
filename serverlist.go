package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mattn/go-mastodon"
)

const seedServer = "infosec.exchange"

type PeerList []string

type Server struct {
	Name      string             `json:"name,omitempty"`
	Peers     []string           `json:"peers,omitempty"`
	Status    bool               `json:"status,omitempty"`
	LastCheck time.Time          `json:"last_check,omitempty"`
	Info      *mastodon.Instance `json:"info,omitempty"`
}

type ServerList struct {
	Servers      map[string]Server
	SeedInstance string
	Peers        []string
}

var (
	SeedInfo = Server{}
	ctx      = context.Background()
)

func NewServerList(seed string) *ServerList {
	s := &ServerList{
		SeedInstance: seed,
		Peers:        PeerList{},
		Servers:      make(map[string]Server),
	}
	return s
}

func (s *ServerList) getSeedInfo() {
	extra := ""
	num := 0
	name := seedServer
	c := mastodon.NewClient(&mastodon.Config{
		Server: "https://" + s.SeedInstance,
	})
	c.Timeout = time.Second * 2
	SeedInfo.LastCheck = time.Now()
	info, err := c.GetInstance(ctx)
	if err != nil {
		s.printServerInfo(num, name, "-> "+err.Error())
		return
	}

	if !s.CheckServerStatus(info) {
		return
	}

	peers, err := c.GetInstancePeers(ctx)
	if err != nil {
		s.printServerInfo(num, name, "-> "+err.Error())
		return
	}
	SeedInfo.Name = info.Title
	SeedInfo.Info = info
	SeedInfo.Status = true
	SeedInfo.Peers = peers
	s.Servers[SeedInfo.Info.URI] = SeedInfo
	s.updateMasterPeerList(SeedInfo.Peers)
	extra += "-> Total Servers: " + fmt.Sprint(len(s.Servers)) + " ( +" + fmt.Sprint(len(peers)) + " = " + fmt.Sprint(len(s.Peers)) + " ) Peers " + info.Version
	s.printServerInfo(num, name, extra)
}

func (s *ServerList) printServerInfo(num int, domain string, extra string) {
	log.Println(fmt.Sprint(num) + " => " + fmt.Sprint(domain) + " " + extra)
}

func (s *ServerList) updateMasterPeerList(peers []string) {
	currentPeers := s.Peers
	// log.Println(fmt.Sprint(len(currentPeers)))
	// log.Println(fmt.Sprint(len(peers)))
	tempPeers := append(currentPeers, peers...)
	// log.Println(fmt.Sprint(len(tempPeers)))
	newPeers := removeDuplicate(tempPeers)
	// log.Println(fmt.Sprint(len(newPeers)))
	s.Peers = newPeers
}

func removeDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func (s *ServerList) getServerInfo(num int, name string) {
	extra := ""
	num = num + 1
	c := mastodon.NewClient(&mastodon.Config{
		Server: "https://" + name,
	})
	c.Timeout = time.Second * 2
	server := Server{}
	server.LastCheck = time.Now()
	info, err := c.GetInstance(ctx)
	if err != nil {
		s.printServerInfo(num, name, "-> "+err.Error())
		return
	}

	if !s.CheckServerStatus(info) {
		return
	}

	peers, err := c.GetInstancePeers(ctx)
	if err != nil {
		s.printServerInfo(num, name, "-> "+err.Error())
		return
	}
	server.Name = info.Title
	server.Info = info
	server.Status = true
	server.Peers = peers
	s.Servers[server.Name] = server
	s.updateMasterPeerList(peers)
	extra += "-> Total Servers: " + fmt.Sprint(len(s.Servers)) + " ( +" + fmt.Sprint(len(peers)) + " = " + fmt.Sprint(len(s.Peers)) + " ) Peers " + info.Version
	s.printServerInfo(num, name, extra)
}

func (s *ServerList) CheckServerStatus(i *mastodon.Instance) bool {
	return i != nil
}
