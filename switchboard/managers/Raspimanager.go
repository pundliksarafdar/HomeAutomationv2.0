package managers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"runtime"
	time2 "time"
)

func GetSwitchData(c *gin.Context){
	timeStr := time2.Now().String()
	timeNano := time2.Now().UnixNano()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	alloc := bToMb(m.Alloc)
	totalAlloc :=  bToMb(m.TotalAlloc)
	sysMem := bToMb(m.Sys)
	gcNum := m.NumGC
	c.JSON(http.StatusOK, gin.H{"version": os.Getenv("build"),"time":timeStr,"timeNano":timeNano,"alloc":alloc,"totalAlloc":totalAlloc,"sysMem":sysMem,"gcNum":gcNum})
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
