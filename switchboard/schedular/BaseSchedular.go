package schedular

import (
	"agent/Devices"
	"fmt"
)

func start(){
fmt.Print("Yes.....")
}

func StartSchedule(){
	StartSchedular(SECOND,5,false,Devices.GetLiveConnection)
}
