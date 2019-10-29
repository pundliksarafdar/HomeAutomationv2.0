package managers

import (
	"agent/data"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

	unixNano := time.Now().UnixNano()
	umillisec := unixNano / 1000000
	dataJson["timestamp"] = umillisec
	dataJson["data"] = dataForES
	PostData(dataJson)
}

func PostData(data map[string]interface{}){
	if data["timestamp"] == nil{
		data["timestamp"] = time.Now().Unix()
	}

	if _, err := os.Stat("/switch"); os.IsNotExist(err) {
		os.Mkdir("/switch", os.ModePerm)
	}
	//esUrl := "http://datastore:9200/switchinfo/doc"
	fileName := fmt.Sprintf("/switch/switch_%v.data",data["timestamp"])
	dByte,_ := json.Marshal(data)
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0600)

	fullPath1, _ := filepath.Abs(fileName)
	log.Info().Msg(fullPath1)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.Write(dByte); err != nil {
		panic(err)
	}


	//Elasticsearch
	/*
		dataByte := bytes.NewBuffer(dByte)
		_, err:= http.Post(esUrl,"application/json",dataByte)
		if err != nil{
			log.Info().Msg("Error occured while posting data to es")
		}*/

}

func GetData() ([]map[string]interface{},[]map[string]interface{}){
	allFilesData := make([]map[string]interface{},0)
	filepath.Walk("/switch", func(path string, info os.FileInfo, err error) error {
		if strings.HasPrefix(info.Name(),"switch") && strings.HasSuffix(info.Name(),".data"){
			file, err := os.Open(path)
			if err != nil{
				log.Error().Msg(err.Error())
			}
			b, err := ioutil.ReadAll(file)
			if err != nil{
				log.Error().Msg(err.Error())
			}
			dataM := map[string]interface{}{}
			json.Unmarshal(b,&dataM)
			allFilesData = append(allFilesData,dataM)
		}
		return nil
	})
	dataTOSend := FormatData(allFilesData)
	return BuildHistoricalGraph(dataTOSend)
}

func FormatData(data []map[string]interface{}) map[string][][]interface{}{
	formatedSwitchData := map[string][][]interface{}{}
	for _,d := range data{
		switchData := d["data"]
		timestamp := d["timestamp"]
		switchDataMap := switchData.([]interface {})

		for _,swDa := range switchDataMap{
			dMap := swDa.(map[string]interface{})

			for switchName,switchStatus := range dMap{
				if switchName == "ip" || switchName == "switchName"{
					continue
				}
				statusData := make([]interface{},0)
				if formatedSwitchData[switchName] == nil{
					formatedSwitchData[switchName] = [][]interface{}{}
				}

				if switchStatus == "ON"{
					statusData = append(statusData,timestamp)
					statusData = append(statusData,1)
				}else if switchStatus == "OFF"{
					statusData = append(statusData,timestamp)
					statusData = append(statusData,0)
				}else {
					if temp, err := strconv.Atoi(switchStatus.(string)); err == nil {
						statusData = append(statusData,timestamp)
						statusData = append(statusData,temp)
					}
				}

				if len(statusData)>0{
					formatedSwitchData[switchName] = append(formatedSwitchData[switchName],statusData)
				}

			}
		}
	}

	return formatedSwitchData
}

func BuildHistoricalGraph(formatedSwitchData map[string][][]interface{})([]map[string]interface{},[]map[string]interface{}){
	graphDataList := make([]map[string]interface{},0)
	graphTempDataList := make([]map[string]interface{},0)
	for key,data := range formatedSwitchData{
		graphData := map[string]interface{}{}
		graphData["name"] = key
		graphData["data"] = data

		if key == "temperature"{
			graphTempDataList = append(graphTempDataList,graphData)
		}else{
			graphDataList = append(graphDataList,graphData)
		}

	}
	return graphDataList,graphTempDataList
}
