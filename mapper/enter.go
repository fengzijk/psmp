package mapper

type BaseMapperGroup struct {
	ShortUrlRecordMapper
	UserMapper
	WxPushRecordMapper
	EmailRecordMapper
}

var MapperGroup = new(BaseMapperGroup)
