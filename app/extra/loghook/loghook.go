package loghook

import "github.com/xtls/xray-core/common/log"

type AppLogHook interface {
	HandleAccessMessage(log.Message)
	HandleDNSLog(log.Message)
	HandleGeneralMessage(log.Message)
}

var appLogHook AppLogHook = &NoopAppLogHook{}

func HandleAccessMessage(msg log.Message) {
	appLogHook.HandleAccessMessage(msg)
}

func HandleDNSLog(msg log.Message) {
	appLogHook.HandleDNSLog(msg)
}

func HandleGeneralMessage(msg log.Message) {
	appLogHook.HandleGeneralMessage(msg)
}

type NoopAppLogHook struct {}

func (h *NoopAppLogHook) HandleAccessMessage(log.Message) {}
func (h *NoopAppLogHook) HandleDNSLog(log.Message) {}
func (h *NoopAppLogHook) HandleGeneralMessage(log.Message) {}

func SetAppLogHook(h AppLogHook) {
	appLogHook = h
}