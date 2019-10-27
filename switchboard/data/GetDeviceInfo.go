package data

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

func GetAndStoreDeviceInfo(ips []string)(map[string]map[string]string,map[string]map[string]string){
	switchData := map[string]map[string]string{}
	portData := map[string]map[string]string{}
	for _,ip := range ips{
		d :=GetAndStoreDeviceInfoForIp(ip)
		port := GetAndStoreDevicePortForIp(ip)
		if len(d)!=0 {
			switchData[ip] = d
			portData[ip] = port
		}
	}
	return switchData,portData
}

func GetAndStoreDeviceInfoForIp(ip string)map[string]string{
	port := "8000"
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
