package protocol

import (
	"strconv"
	"strings"
)

func GetRequest(args []string) []byte {
	req := []string{
		// *<参数数量> CR LF
		"*" + strconv.Itoa(len(args)),
	}

	for _, arg := range args {
		// $<参数1的字节数量> CR LF
		req = append(req, "$"+strconv.Itoa(len(arg)))
		// <参数1的数据> CR LF
		req = append(req, arg)
	}

	str := strings.Join(req, "\r\n")
	// 该协议下所有发送至Redis服务器的参数都是二进制安全(binary safe)的
	return []byte(str + "\r\n")
}
