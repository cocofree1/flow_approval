package lib

const (
	ERR_JSON_UNMARSHAL = 2000 + iota
	ERR_INSERT_TABLE
	ERR_DELETE_TABLE
	ERR_UPDATE_TABLE
	ERR_SEARCH_TABLE
	ERR_STRING_TO_INT
)

var ErrnoMap = map[int]string{
	ERR_JSON_UNMARSHAL:    "json解码失败",
	ERR_INSERT_TABLE:      "插入表失败",
	ERR_DELETE_TABLE:      "删除表失败",
	ERR_UPDATE_TABLE:      "更新表失败",
	ERR_SEARCH_TABLE:      "查询表失败",
	ERR_STRING_TO_INT:     "string转换int失败",
}
