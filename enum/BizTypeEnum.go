package enum

type BizTypeEnum int

const (
	Param = iota + 1 // 1

	Url
)

func (c BizTypeEnum) GetMsg() string {
	switch c {
	case Param:
		return "PARAM"
	case Url:
		return "URL"
	}
	return ""
}
