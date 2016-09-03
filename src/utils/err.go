package utils

import (
	"os"
	"log"
	"fmt"
)


func Log(v ...interface{}) {

	logfile,err:= os.OpenFile("server.log",os.O_RDWR|os.O_APPEND|os.O_CREATE,0);
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return    }
	log.Println(v...)
	logger := log.New(logfile,"\r\n",log.Ldate|log.Ltime);
	logger.SetPrefix("[Info]")
	logger.Println(v...)
	defer logfile.Close();
}