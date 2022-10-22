package response

var zhCNText = map[int]string{
	//ServerError:        "内部服务器错误",
	//TooManyRequests:    "请求过多",
	//ParamBindError:     "参数信息错误",
	AuthorizationError:   "签名信息错误",
	UrlSignError:         "参数签名错误",
	CacheSetError:        "设置缓存失败",
	CacheGetError:        "获取缓存失败",
	CacheDelError:        "删除缓存失败",
	CacheNotExist:        "缓存不存在",
	ResubmitError:        "请勿重复提交",
	HashIdsEncodeError:   "HashID 加密失败",
	HashIdsDecodeError:   "HashID 解密失败",
	IllegalAccess:        "非法访问",
	AuthorizationExpired: "token已经过期",

	TokenExpired:     "令牌过期",
	TokenNotValidYet: "令牌尚未生效",
	TokenMalformed:   "令牌格式错误",
}
