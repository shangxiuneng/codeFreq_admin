package model

const (
	Easy   = 1
	Middle = 2
	Hard   = 3
)

var DifficultMapping = map[int]string{
	Easy:   "容易",
	Middle: "中等",
	Hard:   "困难",
}

const (
	BYTEDANCE = "BYTEDANCE"
	TENCENT   = "TENCENT"
	MEITUAN   = "MEITUAN"
)

var CompanyMapping = map[string]string{
	BYTEDANCE: "字节跳动",
	TENCENT:   "腾讯",
	MEITUAN:   "美团",
}
