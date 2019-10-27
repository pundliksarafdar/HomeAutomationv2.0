package managers

import (
	"agent/data"
	"bytes"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func PostDataToEs() {
	dataJson := map[string]interface{}{}
	dataForES := make([]map[string]string, 0)
	for ip, data := range data.SwitchInfo {
		dataCpy := map[string]string{}
		//reflect.Copy(&dataCpy,&data)
		for k,v := range data{
			dataCpy[k] = v
		}
		dataCpy["ip"] = ip
		dataForES = append(dataForES, dataCpy)
	}
	dataJson["timestamp"] = time.Now().Unix()
	dataJson["data"] = dataForES
	PostData(dataJson)
}

func PostData(data map[string]interface{}){
	esUrl := "http://datastore:9200/switchinfo/doc"
	dByte,_ := json.Marshal(data)
	dataByte := bytes.NewBuffer(dByte)
	_, err:= http.Post(esUrl,"application/json",dataByte)
	if err != nil{
		log.Info().Msg("Error occured while posting data to es")
	}
}
