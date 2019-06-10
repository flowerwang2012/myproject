package data

var datas []string

func Add(str string) string {
	b := []byte(str)
	s := string(b)
	datas = append(datas, s)
	return s
}
