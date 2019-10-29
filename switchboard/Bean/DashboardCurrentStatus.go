package Bean

type Dashboard struct {
	CState CurrentStatus
	Temperature int
}

type CurrentStatus struct {
	ON int
	OFF int
}

type HistoricalData struct {

}
