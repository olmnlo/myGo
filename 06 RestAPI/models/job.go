package models

// job struct
type Job struct {
	Msg          string `json:"msg"`
	TimeStart    string `json:"timeStart"`
	TimeEnd      string `json:"timeEnd"`
	TimeDuration string `json:"timeDuration"`
}

const (
	Broker = "localhost:9092"
)

// global variables for Job struct
var myJobs = []Job{}

func (j Job) Save() {
	myJobs = append(myJobs, j)
}

func GetAllJobs() []Job {
	return myJobs
}
