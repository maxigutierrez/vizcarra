package logs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// GeneralLogger exported
var InfoLogger *log.Logger
var DebugLogger *log.Logger
// ErrorLogger exported
var ErrorLogger *log.Logger
var AuditLogger *log.Logger
var SincLogger *log.Logger
var level int
func init() {
	absPath, err := filepath.Abs("./public/logs")
	if err != nil {
		fmt.Println("Error reading given path:", err)
	}

	generalLog, err := os.OpenFile(absPath+"/general.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
  auditLog, err := os.OpenFile(absPath+"/audit.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	sincroLog, err := os.OpenFile(absPath+"/sincro.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	SincLogger = log.New(sincroLog, "Sinc:\t", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(generalLog, "INFO:\t", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(generalLog, "Error:\t", log.Ldate|log.Ltime|log.Lshortfile)
  DebugLogger = log.New(generalLog, "Debug:\t", log.Ldate|log.Ltime|log.Lshortfile)
  AuditLogger = log.New(auditLog, "",log.Ldate|log.Ltime)
}
