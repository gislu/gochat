package utils

import (
	"database/sql"
	"os"
	"strings"

	"log"
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
