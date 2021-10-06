package main

import (
	"flag"
	"github.com/robfig/cron/v3"
	"github.com/vaniairnanda/send-later/app"
	"github.com/vaniairnanda/send-later/config"
	"github.com/vaniairnanda/send-later/environment"
	"go.uber.org/zap"
)

var (
	// kafka
	kafkaBrokerUrl string
	kafkaVerbose   bool
	kafkaClientId  string
	kafkaTopic     string
)

type customLogger struct{}

func (customLogger) Printf(format string, args ...interface{}) {
	zap.S().Infof(format, args...)
}
func main(){
	loggerConfig := zap.NewDevelopmentConfig()
	zapLogger, _ := loggerConfig.Build()
	defer zapLogger.Sync()
	zap.ReplaceGlobals(zapLogger)

	flag.StringVar(&kafkaBrokerUrl, "kafka-brokers", "localhost:9092", "Kafka brokers in comma separated value")
	flag.BoolVar(&kafkaVerbose, "kafka-verbose", true, "Kafka verbose logging")
	flag.StringVar(&kafkaClientId, "kafka-client-id", "my-kafka-client", "Kafka client id to connect")
	flag.StringVar(&kafkaTopic, "kafka-topic", "foo", "Kafka topic to push")

	dbDisbursement := config.GetDBDisbursement()

	job := NewJob()
	job.InitializePublisher()
	env := environment.Load()
	var err error
	c := cron.New(cron.WithLogger(cron.VerbosePrintfLogger(customLogger{})))


	if _, err = c.AddFunc(env.MarkApprovalExpired, job.JobApprovalExpired); err != nil {
		zapLogger.Error(err.Error())
	}

	if _, err = c.AddFunc(env.ScheduledBatchDisbursement, job.JobScheduledBatchDisbursement); err != nil {
		zapLogger.Error(err.Error())
	}


	if _, err = c.AddFunc(env.SendApprovalReminder, job.JobApprovalReminder); err != nil {
		zapLogger.Error(err.Error())
	}

	c.Start()

	run := app.MakeHandler(dbDisbursement)
	run.Start()


}
