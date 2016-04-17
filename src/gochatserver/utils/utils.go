package utils
import (
	"github.com/astaxie/beego"
)

const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

func Log(log interface{}, level int) {
	switch level {
	case LevelEmergency:
		beego.Emergency(log)
	case LevelAlert:
		beego.Alert(log)
	case LevelCritical:
		beego.Critical(log)
	case LevelError:
		beego.Error(log)
	case LevelWarning:
		beego.Warning(log)
	case LevelNotice:
		beego.Notice(log)
	case LevelInformational:
		beego.Informational(log)
	case LevelDebug:
		beego.Debug(log)
	default:
		beego.Notice(log)
	}
}

