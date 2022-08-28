package dep

// func LoadJson(s string) (map[string]int, error) {

// 	file, err := os.Open(s)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()
// 	fileinfo, err := file.Stat()
// 	if err != nil {
// 		return nil, err
// 	}
// 	cont := make([]byte, fileinfo.Size())
// 	_, err = file.Read(cont)
// 	if err != nil {
// 		return nil, err
// 	}
// 	jsonMap := make(map[string]int)
// 	err = json.Unmarshal(cont, &jsonMap)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return jsonMap, nil
// }

// func WriteJson(s string, m map[string]int) {
// 	cont, err := json.Marshal(m)
// 	if err != nil {
// 		panic(err)
// 	}
// 	file, err := os.OpenFile(s, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	_, err = file.Write(cont)
// 	if err != nil {
// 		panic(err)
// 	}
// }
