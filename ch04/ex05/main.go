package ex05

func RemoveDuplicate(s string) string {
	slice := []byte(s)
	for i := 0; i < len(slice)-1; i++ {
		a, b := slice[i], slice[i+1]
		if a == b {
			slice = remove(slice, i)
			i-- // 削除した分インデックスをひとつ戻す
		}
	}
	return string(slice)

}

func remove(s []byte, i int) []byte {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}
