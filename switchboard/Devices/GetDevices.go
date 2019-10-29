package Devices

import (
	"agent/data"
	"agent/managers"
	"bytes"
	"github.com/rs/zerolog/log"
	"net"
	"reflect"

	"time"
)

var (
	ip1 = net.ParseIP("192.168.1.1")
	ip2 = net.ParseIP("192.168.1.100")
)

func check(ip string) bool {
	trial := net.ParseIP(ip)
	if trial.To4() == nil {
		//log.Printf("%v is not an IPv4 address\n", trial)
		return false
	}
	if bytes.Compare(trial, ip1) >= 0 && bytes.Compare(trial, ip2) <= 0 {
		//log.Printf("%v is between %v and %v\n", trial, ip1, ip2)
		return true
	}
	//log.Printf("%v is NOT between %v and %v\n", trial, ip1, ip2)
	return false
}

func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		if check(ip.String()){
			ips = append(ips, ip.String())
		}
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

//  http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

type Pong struct {
	Ip    string
	Alive bool
}

func ping(pingChan <-chan string, pongChan chan<- Pong) {
	for ip := range pingChan {
		//_, err := exec.Command("ping", ip).Output()
		timeOut := time.Duration(5) * time.Second

		conn, err := net.DialTimeout("tcp", ip+":8000", timeOut)
		if err != nil {

		}else{
			log.Print("Address is "+conn.RemoteAddr().String())
		}
		var alive bool
		if err != nil {
			alive = false
		} else {
			alive = true
		}
		pongChan <- Pong{Ip: ip, Alive: alive}
	}
}

func receivePong(pongNum int, pongChan <-chan Pong, doneChan chan<- []Pong) {
	var alives []Pong
	for i := 0; i < pongNum; i++ {
		pong := <-pongChan
		//  log.Println("received:", pong)
		if pong.Alive {
			alives = append(alives, pong)
		}
	}
	doneChan <- alives
}

func GetLiveConnection(){
	log.Printf("Starting schedular function")
	hosts, _ := Hosts("192.168.1.0/24")

	concurrentMax := 1000
	pingChan := make(chan string, concurrentMax)
	pongChan := make(chan Pong, len(hosts))
	doneChan := make(chan []Pong)

	for i := 0; i < concurrentMax; i++ {
		go ping(pingChan, pongChan)
	}

	go receivePong(len(hosts), pongChan, doneChan)

	for _, ip := range hosts {
		pingChan <- ip
	}

	alives := <-doneChan
	log.Print(alives)
	listIp := make([]string,0)
	for _,alive := range alives{
		log.Printf("Alive items "+alive.Ip)
		listIp = append(listIp,alive.Ip)
	}

	close(pingChan)
	close(pongChan)
	close(doneChan)
	//Set beans for access
	dataNew,port,temp := data.GetAndStoreDeviceInfo(listIp)
	//data.SwitchInfo =
	isEqual := reflect.DeepEqual(dataNew,data.SwitchInfo)
	data.ConnectedDevices = listIp
	data.SwitchPortInfo = port
	data.Temperature = temp
	if !isEqual{
		//Update SwitchInfo sience difference found // in case of ip change you will find see difference
		log.Info().Msg("Pushing data to ES")
		data.SwitchInfo = dataNew
		//Post data to ES
		managers.PostDataToEs()
	}else{
		log.Info().Msg("No differenct found in switch hence not pushing data to ES")
	}
}

