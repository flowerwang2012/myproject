package main

import (
	"bytes"
	"mime/multipart"
	"net/textproto"
	"fmt"
)

func main() {
	test()
}
// 多部分文档文件 multipart 实现http文件传输
func test() (err error) {
	b := []byte("hello world")
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "media", "abc"))
	h.Set("Content-Type", "image/png")
	writer, _ := w.CreatePart(h)
	if _, err = writer.Write(b); err != nil {
		return
	}
	if err = w.Close(); err != nil {
		return
	}

	fmt.Println(buf.String())
	return nil
}
