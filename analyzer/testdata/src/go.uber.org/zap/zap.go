package zap

// Классический логгер
type Logger struct{}

// Sugar возвращает SugaredLogger
func (l *Logger) Sugar() *SugaredLogger { return &SugaredLogger{} }

// Методы логирования (классический API)
func (l *Logger) Debug(msg string, fields ...Field)  {}
func (l *Logger) Info(msg string, fields ...Field)   {}
func (l *Logger) Warn(msg string, fields ...Field)   {}
func (l *Logger) Error(msg string, fields ...Field)  {}
func (l *Logger) DPanic(msg string, fields ...Field) {}
func (l *Logger) Panic(msg string, fields ...Field)  {}
func (l *Logger) Fatal(msg string, fields ...Field)  {}

// Sugared логгер
type SugaredLogger struct{}

// Методы логирования (sugar API)
func (s *SugaredLogger) Debug(args ...any)  {}
func (s *SugaredLogger) Info(args ...any)   {}
func (s *SugaredLogger) Warn(args ...any)   {}
func (s *SugaredLogger) Error(args ...any)  {}
func (s *SugaredLogger) DPanic(args ...any) {}
func (s *SugaredLogger) Panic(args ...any)  {}
func (s *SugaredLogger) Fatal(args ...any)  {}

// structured sugar
func (s *SugaredLogger) Debugw(msg string, keysAndValues ...any)  {}
func (s *SugaredLogger) Infow(msg string, keysAndValues ...any)   {}
func (s *SugaredLogger) Warnw(msg string, keysAndValues ...any)   {}
func (s *SugaredLogger) Errorw(msg string, keysAndValues ...any)  {}
func (s *SugaredLogger) DPanicw(msg string, keysAndValues ...any) {}
func (s *SugaredLogger) Panicw(msg string, keysAndValues ...any)  {}
func (s *SugaredLogger) Fatalw(msg string, keysAndValues ...any)  {}

// printf-style sugar
func (s *SugaredLogger) Debugf(template string, args ...any)  {}
func (s *SugaredLogger) Infof(template string, args ...any)   {}
func (s *SugaredLogger) Warnf(template string, args ...any)   {}
func (s *SugaredLogger) Errorf(template string, args ...any)  {}
func (s *SugaredLogger) DPanicf(template string, args ...any) {}
func (s *SugaredLogger) Panicf(template string, args ...any)  {}
func (s *SugaredLogger) Fatalf(template string, args ...any)  {}

// printf-style sugar
func (s *SugaredLogger) Debugln(template string, args ...any)  {}
func (s *SugaredLogger) Infoln(template string, args ...any)   {}
func (s *SugaredLogger) Warnln(template string, args ...any)   {}
func (s *SugaredLogger) Errorln(template string, args ...any)  {}
func (s *SugaredLogger) DPanicln(template string, args ...any) {}
func (s *SugaredLogger) Panicln(template string, args ...any)  {}
func (s *SugaredLogger) Fatalln(template string, args ...any)  {}
func L() *Logger                                               { return &Logger{} }

func S() *SugaredLogger { return &SugaredLogger{} }
