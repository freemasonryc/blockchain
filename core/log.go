package core

import (
	"freemasonry.cc/log"
	"github.com/sirupsen/logrus"
	"strings"
)


var (
	LmChainClient      = log.RegisterModule("vc-cli", logrus.InfoLevel) 
	LmChainType        = log.RegisterModule("vc-ty", logrus.InfoLevel)  
	LmChainKeeper      = log.RegisterModule("vc-kp", logrus.InfoLevel)  
	LmChainMsgServer   = log.RegisterModule("vc-ms", logrus.InfoLevel)  
	LmChainRest        = log.RegisterModule("vc-re", logrus.InfoLevel)  
	LmChainMsgAnalysis = log.RegisterModule("vc-mas", logrus.InfoLevel) 
	LmChainUtil        = log.RegisterModule("vc-ut", logrus.InfoLevel)  
)


func BuildLog(funcName string, modules ...log.LogModule) *logrus.Entry {
	moduleName := ""
	for _, v := range modules {
		if moduleName != "" {
			moduleName += "/"
		}
		moduleName += string(v)
	}
	logEntry := log.Log.WithField("module", strings.ToLower(moduleName))
	if funcName != "" {
		logEntry = logEntry.WithField("method", strings.ToLower(funcName))
	}
	return logEntry
}
