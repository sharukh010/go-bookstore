package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request,x interface{}){
	if body, err := io.ReadAll(r.Body); err == nil {
		//converting json -> go struct
		if err := json.Unmarshal([]byte(body),x);err != nil {
			return 
		}
	}
}