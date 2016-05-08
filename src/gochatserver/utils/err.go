package utils

import (
	"database/sql"
	"os"
	"strings"

	"log"
	"fmt"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func CheckNotFoundErr(err error) bool {
	if err != nil {
		if strings.HasSuffix(err.Error(), "not found") {
			return true
		} else {
			log.Fatal(err)
			panic(err)
		}
	}
	return false
}

func CheckNoRowInSetErr(err error) bool {
	if err != nil {
		if err == sql.ErrNoRows {
			return true
		} else {
			log.Fatal(err)
			panic(err)
		}
	}
	return false
}

func CheckNoFileErr(err error) bool {
	if err != nil {
		if os.IsNotExist(err) {
			return true
		} else {
			log.Fatal(err)
			panic(err)
		}
	}
	return false
}

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