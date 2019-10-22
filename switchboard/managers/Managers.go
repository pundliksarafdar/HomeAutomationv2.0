package managers

import (
	"agent/Bean"
	"agent/data"
	"agent/post"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type SwitchManager struct {

}
func (sM *SwitchManager)GetSwitchData(c *gin.Context){
	switchData := data.SwitchInfo
	switchPortData := data.SwitchPortInfo
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": switchData,"port":switchPortData})
}

func (sM *SwitchManager)GetSwitchIps(c *gin.Context){
	switchData := data.ConnectedDevices
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": switchData})
}


func (sM *SwitchManager)ToggleSwitch(c *gin.Context){
	var sT Bean.SwitchToggle
	er:= c.BindJSON(&sT)
	if er!=nil{
		log.Print("Error occured...")
	}
	success := post.ToggleSwitch(sT)
	if success{
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	}else{
		c.JSON(http.StatusOK, gin.H{"status": http.StatusExpectationFailed})
	}

}
