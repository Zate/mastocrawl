package main

import (
	"fmt"
	"log"
	"time"
)

// Channels

// GetServerChan to submit server name
var GetServerChan = make(chan struct {
	int
	string
}, 5)

// ParseFilesChan to parse file
// var ParseServerChan = make(chan string, 5)

// DoneChan to signal intel import is done.
// var DoneChan = make(chan bool, 1)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	s := NewServerList(seedServer)
	s.getSeedInfo()
	// go s.GetServerNames()
	// go s.ParseServer()
	// x := 0
	// for i, p := range s.Peers {
	// 	if x > 10 {
	// 		os.Exit(0)
	// 	}
	// 	s.getServerInfo(i, p)
	// 	x++
	// }

	defer close(GetServerChan)
	// defer close(ParseServerChan)
	// defer close(DoneChan)
	go GetServerNames(s)
	go ParseServer(s)
	// go Stage2()
}

func ParseServer(s *ServerList) {
	log.Println("ParseServer Started")
	for {
		thing := <-GetServerChan
		time.Sleep(time.Second * 10)
		log.Println("Processing: "+fmt.Sprint(thing.int), thing.string)
	}
	// return nil
}

func GetServerNames(s *ServerList) {
	log.Println("GetServerNames Started")

	for i, p := range s.Peers {
		GetServerChan <- struct {
			int
			string
		}{i, p}
		log.Println("Sent: "+fmt.Sprint(i), p)
	}
}

// // Stage2 is a goroutine to kick off stage 2
// func Stage2() {
// 	for {
// 		done := <-DoneChan
// 		if done == true {
// 			GetData()
// 		}
// 	}
// }

// // GetIntel function to download an Intel file
// func GetIntel() (err error) {
// 	log.Println("GetIntel Started")
// 	num := 0
// 	var f feed

// 	for num >= 0 {
// 		start := time.Now()
// 		f = <-GetFileChan
// 		filepath := "ipsets/" + f.Path
// 		//for f = range feed {
// 		log.Printf("Starting to process %v", filepath)
// 		//log.Println(f.Path)
// 		//log.Println(f.URL)
// 		//log.Println(f.Name)
// 		if _, err := os.Stat("ipsets"); os.IsNotExist(err) {
// 			err = os.MkdirAll("ipsets", 0755)
// 			if err != nil {
// 				log.Println(err)
// 			}
// 		}

// 		out, err := os.Create(filepath)
// 		if err != nil {
// 			log.Println(err)

// 		}
// 		//defer out.Close()
// 		//log.Println(out)

// 		resp, err := http.Get(f.URL)
// 		//log.Println(resp)
// 		if err != nil {
// 			log.Println(err)

// 		}

// 		defer resp.Body.Close()

// 		_, err = io.Copy(out, resp.Body)
// 		if err != nil {
// 			log.Println(err)

// 		}
// 		//}
// 		ParseFilesChan <- filepath
// 		dur := time.Since(start)
// 		log.Printf("Completed getting %v and it took %v", filepath, dur)
// 	}
// 	return err
// }

// // ParseFile takes a filepath and grabs some values out of it.
// // Only for files from github.com/firehol/ipsets
// func ParseFile() (err error) {
// 	log.Println("ParseFile Started")
// 	// #
// 	// # tor_exits_30d
// 	// #
// 	// # ipv4 hash:ip ipset
// 	// #
// 	// # [TorProject.org] (https://www.torproject.org) list of all
// 	// # current TOR exit points (TorDNSEL)
// 	// #
// 	// # Maintainer      : TorProject.org
// 	// # Maintainer URL  : https://www.torproject.org/
// 	// # List source URL : https://check.torproject.org/exit-addresses
// 	// # Source File Date: Tue Dec 29 08:02:26 UTC 2020
// 	// #
// 	// # Category        : anonymizers
// 	// # Version         : 14747
// 	// #
// 	// # This File Date  : Tue Dec 29 08:08:05 UTC 2020
// 	// # Update Frequency: 5 mins
// 	// # Aggregation     : 30 days
// 	// # Entries         : 2633 unique IPs
// 	// #
// 	// # Full list analysis, including geolocation map, history,
// 	// # retention policy, overlaps with other lists, etc.
// 	// # available at:
// 	// #
// 	// #  http://iplists.firehol.org/?ipset=tor_exits_30d
// 	// #
// 	// # Generated by FireHOL's update-ipsets.sh
// 	// # Processed with FireHOL's iprange
// 	var (
// 		//fileComment = regexp.MustCompile(`^#.*`)
// 		fileTag = regexp.MustCompile(`#\s\shttp\://iplists\.firehol\.org\/\?ipset=(.*)$`)
// 		//feedInfo   = regexp.MustCompile(`^#\s\s(?P<stuff>http.*)`)
// 		maintainer = regexp.MustCompile(`^#\sMaintainer\s\s.*:\s(?P<stuff>.*)`)
// 		maintURL   = regexp.MustCompile(`^#\sMaintainer\sURL\s.*:\s(?P<stuff>.*)`)
// 		category   = regexp.MustCompile(`^#\sCategory\s.*:\s(?P<stuff>.*)`)
// 		done       = regexp.MustCompile(`^#\sProcessed\swith\sFireHOL.*`)
// 		isip       = regexp.MustCompile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`)
// 		tag        string
// 		//infourl    string
// 		maint string
// 		murl  string
// 		cat   string
// 		fid   int64
// 		num   int
// 	)
// 	err = DB.Ping()

// 	if err != nil {
// 		log.Fatalf("Error: %v\n", err)
// 	}
// 	for num >= 0 {
// 		counter := 0
// 		start := time.Now()
// 		filepath := <-ParseFilesChan
// 		log.Printf("Starting parsing %v", filepath)
// 		file, err := os.Open(filepath)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer file.Close()

// 		scanner := bufio.NewScanner(file)
// 		for scanner.Scan() {
// 			counter++

// 			text := scanner.Text()
// 			switch {
// 			case fileTag.MatchString(text):
// 				tag = fileTag.FindStringSubmatch(text)[1]
// 				//log.Printf("\tTag: %s\n", tag)
// 				//log.Printf("\tInfo URL: http://iplists.firehol.org/?ipset=%s", tag)
// 			//case feedInfo.MatchString(text):
// 			//	infourl = feedInfo.FindStringSubmatch(text)[1]
// 			//	log.Printf("\tInfo URL: %s\n", infourl)
// 			case maintainer.MatchString(text):
// 				maint = maintainer.FindStringSubmatch(text)[1]
// 				//log.Printf("\tMaintainer: %s\n", maint)
// 			case maintURL.MatchString(text):
// 				murl = maintURL.FindStringSubmatch(text)[1]
// 				//log.Printf("\tMaintainer URL: %s\n", murl)
// 			case category.MatchString(text):
// 				cat = category.FindStringSubmatch(text)[1]
// 				//log.Printf("\tCategory: %s\n", cat)
// 			case done.MatchString(text):
// 				infourl := "http://iplists.firehol.org/?ipset=" + tag
// 				stmt, err := DB.Prepare("INSERT INTO feeds (tag, category, feed_name, feed_url, info_url) VALUES (?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE tag=VALUES(tag), category=VALUES(category), feed_name=VALUES(feed_name), feed_url=VALUES(feed_url), info_url=VALUES(info_url), updated=CURRENT_TIMESTAMP()")
// 				if err != nil {
// 					log.Fatal(err)
// 				}

// 				res, err := stmt.Exec(tag, cat, maint, murl, infourl)
// 				if err != nil {
// 					log.Println(err)
// 					res, err = stmt.Exec(tag, cat, maint, murl, infourl)
// 				}
// 				fid, err = res.LastInsertId()
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 				//log.Printf("\tFeed id: %v\n", fid)
// 				stmt.Close()
// 			case isip.MatchString(text):
// 				var ipid int64
// 				ipaddr := ip2int(net.ParseIP(text))

// 				// insert / update each ip, get last id as ip_id
// 				stmt, err := DB.Prepare("INSERT INTO ip (ipaddr) VALUES (?) ON DUPLICATE KEY UPDATE updated=CURRENT_TIMESTAMP()")
// 				if err != nil {
// 					log.Fatal(err)
// 				}

// 				res, err := stmt.Exec(ipaddr)
// 				if err != nil {
// 					log.Println(err)
// 					res, err = stmt.Exec(ipaddr)
// 				}
// 				ipid, err = res.LastInsertId()
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 				//log.Printf("\tIP : %v %v %v %v", scanner.Text(), ipaddr, ipid, id)
// 				//insert each entry as intel with feed_id and ip_id
// 				//log.Printf("\tIP: %v\n", scanner.Text())
// 				stmt.Close()
// 				stmt, err = DB.Prepare("INSERT INTO intel (ip_id, feed_id) VALUES (?, ?) ON DUPLICATE KEY UPDATE updated=CURRENT_TIMESTAMP()")
// 				if err != nil {
// 					log.Fatal(err)
// 				}

// 				res, err = stmt.Exec(ipid, fid)
// 				if err != nil {
// 					log.Println(err)
// 					res, err = stmt.Exec(ipid, fid)

// 				}
// 				stmt.Close()
// 			default:
// 			}

// 			if counter%5000 == 0 {
// 				log.Printf("\tProcessed: %v of %v", counter, filepath)
// 			}

// 		}
// 		dur := time.Since(start)
// 		log.Printf("Completed %v with feed_id %v in %v", filepath, fid, dur)
// 		GetData()
// 		DoneChan <- true
// 	}
// 	return nil
// }
