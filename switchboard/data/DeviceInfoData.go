package data

var(
	SwitchInfo = map[string]map[string]string{}
	SwitchPortInfo = map[string]map[string]string{}
	ConnectedDevices = make([]string,0)
)

func GetConnectedDevicesIps()[]string{
	return ConnectedDevices
}

func GetConnectedDeviceInfo()map[string]map[string]string{
	return SwitchInfo
}
