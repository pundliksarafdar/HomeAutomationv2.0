package managers

import (
	"agent/Bean"
	"agent/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AnalyticsManager struct{

}

func (am *AnalyticsManager)GetHistoricalData(c *gin.Context){
	historicData,tempdata := GetData()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "historicData": historicData,"tempdata":tempdata})
}

func (am *AnalyticsManager)GetCurrentStatus(c *gin.Context){
	on :=0
	off := 0
	for _,val := range data.SwitchInfo{
		if val!=nil{
			for _,state := range val{
				if state == "ON"{
					on ++
				}else if state == "OFF"{
					off ++
				}
			}
		}
	}

	dashboard := Bean.Dashboard{CState:Bean.CurrentStatus{OFF:off,ON:on},Temperature: data.Temperature}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": dashboard})
}
