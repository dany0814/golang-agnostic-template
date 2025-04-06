package logger

func InitLogger() (ILogger, error) {
	var initErr error
	once.Do(func() {
		instance, initErr = newZapLogger()
	})
	return instance, initErr
}
