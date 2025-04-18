package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func SetupLogger() {
	// Diretório onde os logs serão armazenados
	logDir := "storage/log"

	// Garantir que o diretório existe
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("Erro ao criar diretório de logs: %v", err)
	}

	// Nome do arquivo baseado na data atual
	today := time.Now().Format("2006-01-02")
	logPath := filepath.Join(logDir, fmt.Sprintf("app-%s.log", today))

	// Abrir arquivo de log
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Erro ao abrir arquivo de log: %v", err)
	}

	// Configurar o logger
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Rotacionar logs mantendo apenas os últimos 15 dias
	go rotateOldLogs(logDir, 15)
}

func rotateOldLogs(logDir string, maxDays int) {
	// Listar todos os arquivos de log
	files, err := filepath.Glob(filepath.Join(logDir, "app-*.log"))
	if err != nil {
		log.Printf("Erro ao listar arquivos de log: %v", err)
		return
	}

	// Se não tivermos mais que o máximo de dias, não precisamos remover nada
	if len(files) <= maxDays {
		return
	}

	// Ordenar arquivos por data (do mais antigo para o mais recente)
	sort.Slice(files, func(i, j int) bool {
		return files[i] < files[j]
	})

	// Remover arquivos mais antigos para manter apenas os últimos 'maxDays'
	filesToRemove := files[:len(files)-maxDays]
	for _, file := range filesToRemove {
		if err := os.Remove(file); err != nil {
			log.Printf("Erro ao remover arquivo antigo %s: %v", file, err)
		} else {
			log.Printf("Arquivo antigo removido: %s", file)
		}
	}
}
