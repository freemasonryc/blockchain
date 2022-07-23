package util

import (
	"freemasonry.cc/log"
	"github.com/sirupsen/logrus"
	"strings"
)


var (
	LmChainClient      = log.RegisterModule("ccli", logrus.InfoLevel)      
	LmChainType        = log.RegisterModule("chain-t", logrus.InfoLevel)   
	LmChainKeeper      = log.RegisterModule("chain-kp", logrus.InfoLevel)  
	LmChainMsgServer   = log.RegisterModule("chain-ms", logrus.InfoLevel)  
	LmChainRest        = log.RegisterModule("chain-re", logrus.InfoLevel)  
	LmChainMsgAnalysis = log.RegisterModule("chain-mas", logrus.InfoLevel) 
	LmChainUtil        = log.RegisterModule("chain-ut", logrus.InfoLevel)  
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
