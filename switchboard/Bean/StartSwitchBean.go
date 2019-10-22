package Bean

type SwitchToggle struct {
	Port   string `json:"port" binding:"required"`
	Status string `json:"status" binding:"required"`
	Ip     string `json:"ip" binding:"required"`
}

type SWITCH_STATUS string

const ON SWITCH_STATUS = "on"
const OFF SWITCH_STATUS = "off"
