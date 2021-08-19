package util

import "net/http"

func SetHeader(w http.ResponseWriter){
	w.Header().Set("Content-Type", "application/json")
}

//func getParam(vals map[string]string) interface{} {
//	v, ok := vals["type"]
//	var pt string
//	if ok {
//		if len(v) >= 1 {
//			pt = v[0]
//		}
//	}
//}