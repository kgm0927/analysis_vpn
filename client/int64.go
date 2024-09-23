package client

import "encoding/json"

func InterfaceToInt64(item interface{}) int64 {
	var i int64
	switch item := item.(type) {
	case json.Number:
		i, _ = item.Int64()
	case int64:
		i = item
	}

	return i
}

// 숫자일 시 i 값 설정, json 파일이면 숫자화해서 i에 저장.
