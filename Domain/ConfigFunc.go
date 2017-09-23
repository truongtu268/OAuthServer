package Domain

func GetConfigFile() chan DataConfig {
	channel:= make(chan DataConfig)
	go func() {
		channel<- LoadConfig("config.json")
	}()
	return channel
}