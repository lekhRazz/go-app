package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"sample_go_app/src/config"
	"time"

	"github.com/labstack/echo"
)

var (
	WarningLogger    *log.Logger
	InfoLogger       *log.Logger
	ErrorLogger      *log.Logger
	DebugLogger      *log.Logger
	LogPath          string = "src/logs"
	DateFormatLayout string = "2006-01-02"
)

type LogMessage struct {
	When        time.Time `json:"@timestamp"`
	Message     string    `json:"message"`
	DebugId     string    `json:"debugId"`
	UserAgent   string    `json:"userAgent"`
	Origin      string    `json:"origin"`
	Host        string    `json:"host"`
	OriginalUrl string    `json:"originalUrl"`
	Method      string    `json:"method"`
	StackTrace  string    `json:"stackTrace"`
	HandlerName string    `json:"handlerName"`
	LogLevel    string    `json:"logLevel"`
}

func getErrorMsg(handlerName string, message string, debug config.Debug, sysErr error, logLevel string) LogMessage {
	e := LogMessage{
		When:        time.Now(),
		Message:     message,
		DebugId:     debug.DebugId,
		UserAgent:   debug.UserAgent,
		Origin:      debug.Origin,
		Host:        debug.Host,
		OriginalUrl: debug.OriginalUrl,
		Method:      debug.Method,
		HandlerName: handlerName,
		StackTrace:  fmt.Sprintf("%v", sysErr),
		LogLevel:    logLevel,
	}

	return e
	// fmt.Sprintf("[ Time : %v] [Message : %s] [DebugId : %s] [UserAgent : %s] [Origin : %s] [Host : %s] [OriginalUrl : %s] [Method : %s] [HandlerName : %s] [StackTrace : %s]",
	// 	e.When, e.Message, e.DebugId, e.UserAgent, e.Origin, e.Host, e.OriginalUrl, e.Method, e.HandlerName, e.StackTrace)
}

func getMainLogMsg(handlerName string, message string, debug config.Debug, logLevel string) LogMessage {
	e := LogMessage{
		When:        time.Now(),
		Message:     message,
		DebugId:     debug.DebugId,
		UserAgent:   debug.UserAgent,
		Origin:      debug.Origin,
		Host:        debug.Host,
		OriginalUrl: debug.OriginalUrl,
		Method:      debug.Method,
		HandlerName: handlerName,
		LogLevel:    logLevel,
	}
	return e
}

/*
*  @author lekhrazz
*  @params (logLevel) String
*  This method creates folders to maintain logs according to log level
 */

func createLogRepos(logLevel string) (string, error) {
	var dirPath string = path.Join(LogPath, logLevel)
	if err := os.MkdirAll(dirPath, 0777); err != nil {
		// s := err.Error()
		// fmt.Printf("type: %T; value: %q\n", s, s)
		return "", err
	}

	var logFileName string = logLevel + "-" + time.Now().Format(DateFormatLayout) + ".log"
	var logFilePath string = path.Join(dirPath, logFileName)
	return logFilePath, nil
}

/*
* @author lekhrazz
* @params (message String)
 */
func Info(ctx echo.Context, message string) {
	logFilePath, fileErr := createLogRepos("info")
	if fileErr != nil {
		log.Fatal(fileErr)
	}

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	debug := ctx.Get("debug").(config.Debug)

	logMsgStruct := getMainLogMsg(ctx.Get("HandlerName").(string), message, debug, "INFO")
	logJsonMsg, _ := json.Marshal(logMsgStruct)
	InfoLogger.Println(fmt.Sprint(string(logJsonMsg)))

	InsertDocs(&logMsgStruct)

}

func Debug(ctx echo.Context, message string) {
	logFilePath, fileErr := createLogRepos("debug")
	if fileErr != nil {
		log.Fatal(fileErr)
	}

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	DebugLogger = log.New(file, "DEBUG: ", log.Lshortfile)
	debug := ctx.Get("debug").(config.Debug)

	logMsgStruct := getMainLogMsg(ctx.Get("HandlerName").(string), message, debug, "DEBUG")
	logJsonMsg, _ := json.Marshal(logMsgStruct)
	DebugLogger.Println(fmt.Sprint(string(logJsonMsg)))

	InsertDocs(&logMsgStruct)
}

/*
* function name {Warn}
* @params (context, message)
*
 */
func Warn(ctx echo.Context, message string) {
	logFilePath, fileErr := createLogRepos("warn")
	if fileErr != nil {
		log.Fatal(fileErr)
	}

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	WarningLogger = log.New(file, "WARN: ", log.Lshortfile)
	debug := ctx.Get("debug").(config.Debug)

	logMsgStruct := getMainLogMsg(ctx.Get("HandlerName").(string), message, debug, "WARNING")
	logJsonMsg, _ := json.Marshal(logMsgStruct)
	WarningLogger.Println(fmt.Sprint(string(logJsonMsg)))

	InsertDocs(&logMsgStruct)
}

func Error(ctx echo.Context, message string, sysErr error) {
	logFilePath, fileErr := createLogRepos("error")
	if fileErr != nil {
		log.Fatal(fileErr)
	}

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	ErrorLogger = log.New(file, "ERROR: ", log.Lshortfile)
	debug := ctx.Get("debug").(config.Debug)

	logMsgStruct := getErrorMsg(ctx.Get("HandlerName").(string), message, debug, sysErr, "ERROR")
	logJsonMsg, _ := json.Marshal(logMsgStruct)
	WarningLogger.Println(fmt.Sprint(string(logJsonMsg)))

	InsertDocs(&logMsgStruct)
}
