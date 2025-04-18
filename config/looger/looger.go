package logger

import (
	"io"
	"log"
	"os"
)

func SetupLogger() {
	logFile, err := os.OpenFile("storage/log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("Erro ao abrir arquivo de log: %v", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
