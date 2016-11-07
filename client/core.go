package client

import (
	"net"
	"errors"
	"fmt"
	"encoding/json"

	AntTcp "github.com/bannerchi/dorylus/tcp"
)
type ProcessState struct {
	Pid           int32 `json:"pid"`
	IsRunning     bool `json:"is_running"`
	MemoryPercent float32 `json:"memory_percent"`
	CpuPercent    float64 `json:"cpu_percent"`
}

type RetRunJob struct {
	Tid 		int `json:"tid"`
	Status 		int `json:"status"`
	Name        string `json:"name"`
	Pid 		int `json:"pid"`
}


const (
	LOAD_AVERAGE = "get_load_average"
	PROC_STATUS = "get_proc_status_%d"
	RUN_JOB = "run_task_%d"
	REMOVE_JOB = "rm_task_%d"
	READY_TO_RUN_JOB = "ready_to_run_jobs_%d"
	MEMORY = "get_memory"
	CPU_INFO = "get_cpu"
	PROC_STATUS_REGEX = `^get_proc_status_-?\d+$`
	NUMBER           = `-?\d+$`
)

/**
	@return json
 */
func GetLoadAverage(host string) (string, error) {
	res, err := getFromDorylus(host, LOAD_AVERAGE)
	if err != nil {
		return "", err
	}
	return res, nil
}

func RunJobByIdAndServerId(taskId int, host string) (string, error) {
	res, err := getFromDorylus(host, fmt.Sprintf(RUN_JOB, taskId))
	if err != nil {
		return "", err;
	}
	return res, nil
}

func RmJobByIdAndServerId(taskId int, host string) (string, error) {
	res, err := getFromDorylus(host, fmt.Sprintf(REMOVE_JOB, taskId))
	if err != nil {
		return "", err
	}
	return res, nil
}
/**
	@return json
 */
func GetProcStatusByPid(pid int, host string) (string, error) {
	res, err := getFromDorylus(host, fmt.Sprintf(PROC_STATUS, pid))
	if err != nil {
		return "", err
	}

	return res, nil
}
/**
 	@return json
 */
func GetMemory(host string) (string, error) {
	res, err := getFromDorylus(host, MEMORY)
	if err != nil {
		return "", err
	}
	return res, nil
}

// if size==0 all jobs will return
func GetReadyToRunJob(size int, host string) ([]RetRunJob, error) {
	var sliceEntry []RetRunJob
	res, err := getFromDorylus(host, fmt.Sprintf(READY_TO_RUN_JOB, size))

	if err != nil {
		return sliceEntry, err
	} else {
		json.Unmarshal([]byte(res), &sliceEntry)
	}
	return sliceEntry, nil
}
/**
  get json from dorylus
 */
func getFromDorylus(host string, section string) (string, error) {
	conn := GetConnection(host)

	resp, err1 := GetResponse(section, conn)
	if err1 != nil {
		return "", err1
	}

	return resp, nil
}

func GetConnection(dorylusDomain string) *net.TCPConn {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", dorylusDomain)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	return conn
}

func GetResponse(req string, conn *net.TCPConn) (string, error) {
	var resp string
	if req == "" {
		checkError(errors.New("request can not be empty"))
	}
	defer conn.Close()

	conn.Write(AntTcp.NewEchoPacket([]byte(req), false).Serialize())
	antProtocol := &AntTcp.EchoProtocol{}
	p, err := antProtocol.ReadPacket(conn)
	if err != nil {
		return resp, err
	} else {
		antPacket := p.(*AntTcp.EchoPacket)
		resp = string(antPacket.GetBody())
	}

	return resp, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Printf(err.Error())
	}
}

