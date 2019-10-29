package data

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strconv"
)

var port = "8000"

//Temperature will be pulled from any one but only one device but will be present in datastore
func GetAndStoreDeviceInfo(ips []string)(map[string]map[string]string,map[string]map[string]string,int){
	switchData := map[string]map[string]string{}
	portData := map[string]map[string]string{}
	var temp int;
	for _,ip := range ips{
		d :=GetAndStoreDeviceInfoForIp(ip)
		port := GetAndStoreDevicePortForIp(ip)
		temperature := GetTemperatureForIp(ip)
		if len(d)!=0 {
			if temperature < 100 {
				temp = temperature
				d["temperature"] = strconv.Itoa(temperature)
			}
			switchData[ip] = d
			portData[ip] = port
		}
	}
	return switchData,portData,temp
}

func GetTemperatureForIp(ip string)int{
	var temp = 101
	dataM := map[string]int{}
	deviceTemperatureUrl := fmt.Sprintf("http://%s:%s/temp",ip,port)
	resp, err := http.Get(deviceTemperatureUrl)
	log.Info().Msg(deviceTemperatureUrl)
	defer func(){
		if nil != resp && nil != resp.Body{
			resp.Body.Close()}
	}()
	if err!=nil{
		log.Info().Msg("Error occured for device "+ip)
	}else{

		if resp.StatusCode == 200{
			body, err := ioutil.ReadAll(resp.Body)
			if err!=nil{
				log.Info().Msg("Error occured for device "+ip)
			}
			bodyMap := string(body)
			log.Print(bodyMap)
			er :=json.Unmarshal(body,&dataM)

			if er != nil{
				log.Info().Msg(er.Error())
			}
			temp = dataM["temperature"]
		}

	}
	return temp
}

func GetAndStoreDeviceInfoForIp(ip string)map[string]string{
	dataM := map[string]string{}
	deviceStatusUrl := fmt.Sprintf("http://%s:%s/status",ip,port)
	resp, err := http.Get(deviceStatusUrl)
	log.Info().Msg(deviceStatusUrl)
	defer func(){
		if nil != resp && nil != resp.Body{
			resp.Body.Close()}
		}()
	if err!=nil{
		log.Info().Msg("Error occured for device "+ip)
	}else{

			body, err := ioutil.ReadAll(resp.Body)
			if err!=nil{
				log.Info().Msg("Error occured for device "+ip)
			}
		bodyMap := string(body)
		log.Print(bodyMap)
		er :=json.Unmarshal(body,&dataM)

		if er != nil{
			log.Info().Msg(er.Error())
		}
	}
	return dataM
}

func GetAndStoreDevicePortForIp(ip string)map[string]string{
	port := "8000"
	dataM := map[string]string{}
	deviceStatusUrl := fmt.Sprintf("http://%s:%s/port",ip,port)
	resp, err := http.Get(deviceStatusUrl)
	log.Info().Msg(deviceStatusUrl)
	defer func(){
		if nil != resp && nil != resp.Body{
			resp.Body.Close()}
	}()
	if err!=nil{
		log.Info().Msg("Error occured for device "+ip)
	}else{

		body, err := ioutil.ReadAll(resp.Body)
		if err!=nil{
			log.Info().Msg("Error occured for device "+ip)
		}
		bodyMap := string(body)
		log.Print(bodyMap)
		er :=json.Unmarshal(body,&dataM)

		if er != nil{
			log.Info().Msg(er.Error())
		}
	}
	return dataM
}
