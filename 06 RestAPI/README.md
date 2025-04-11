# Welcome To Our RestAPI
## our application:
### main.go have some functions:

[x]`getJobs()`: http://localhost:8080/jobs/{topicname} `GET`

[x]`createJobs()` http://localhost:8080/jobs/{topicname} `POST`

### models
1. job.go: 
- * `Job` Struct
- * `Broker`
- * `func (j Job) Save()` => assume it save to database
- * `func GetAllJobs()`
2. sendReport.go
- * `func SendEmail(subject string, body Job, to string) error` => after consume send report to email
3. consumer.go
