package lib

const (
	ERR_JSON_UNMARSHAL = 2000 + iota
	ERR_INSERT_TABLE
	ERR_SEARCH_TABLE
)

var ErrnoMap = map[int]string{
	ERR_JSON_UNMARSHAL:    "json解码失败",
	ERR_INSERT_TABLE:      "插入表失败",
	ERR_SEARCH_TABLE:      "查询表失败",
}
