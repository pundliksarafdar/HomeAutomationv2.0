package post

import (
	"agent/Bean"
	"fmt"
	"net/http"
)

func ToggleSwitch(toggle Bean.SwitchToggle)bool{
	toggleUrl := fmt.Sprintf("http://%s:%s/%s",toggle.Ip,toggle.Port,toggle.Status)
	r,e := http.Get(toggleUrl)

	if e!=nil || r == nil || r.StatusCode != 200{
		return true
	}else{
		return true
	}
}
