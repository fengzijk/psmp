package service

type BaseServiceGroup struct {
	EmailService
	ShortService
	WxPushService
	JwtService
	LocalRemoteIPService
	DingTalkService
}

var ServiceGroup = new(BaseServiceGroup)
