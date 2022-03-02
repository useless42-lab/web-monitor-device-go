package jobs

import (
	"WebMonitorDevice/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type CheckServerJob struct {
}

func (checkServerJob CheckServerJob) Run() {
	HttpPostForm()
}

func HttpPostForm() {
	cpuUser, cpuSytem, cpuIdle, cpuPercent := utils.GetCpuInfo()
	memoryTotal, memoryAvailable, memoryUsed, memoryUsedPercent := utils.GetMemoryInfo()
	diskTotal, diskFree, diskUsed, diskUsedPercent := utils.GetDiskInfo()
	netSent, netRecv := utils.GetNetInfo()
	resp, err := http.PostForm(os.Getenv("API_URL")+"/"+os.Getenv("TOKEN"),
		url.Values{
			"token":               {os.Getenv("TOKEN")},
			"cpu_user":            {strconv.FormatFloat(cpuUser, 'f', 6, 64)},
			"cpu_system":          {strconv.FormatFloat(cpuSytem, 'f', 6, 64)},
			"cpu_idle":            {strconv.FormatFloat(cpuIdle, 'f', 6, 64)},
			"cpu_percent":         {strconv.FormatFloat(cpuPercent, 'f', 6, 64)},
			"memory_total":        {strconv.FormatUint(memoryTotal, 10)},
			"memory_available":    {strconv.FormatUint(memoryAvailable, 10)},
			"memory_used":         {strconv.FormatUint(memoryUsed, 10)},
			"memory_used_percent": {strconv.FormatFloat(memoryUsedPercent, 'f', 6, 64)},
			"disk_total":          {strconv.FormatUint(diskTotal, 10)},
			"disk_free":           {strconv.FormatUint(diskFree, 10)},
			"disk_used":           {strconv.FormatUint(diskUsed, 10)},
			"disk_used_percent":   {strconv.FormatFloat(diskUsedPercent, 'f', 6, 64)},
			"net_sent":            {strconv.FormatUint(netSent, 10)},
			"net_recv":            {strconv.FormatUint(netRecv, 10)},
		})

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))

}
