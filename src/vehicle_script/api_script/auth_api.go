package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

func main()  {

	req_url:="http://localhost:7001/auth"

	bodyParams := map[string]interface{}{
		"user_name":"safly",
		"password":"838facfab6e49cd2d54d064864ceeb00",
	}
	resp,_:=tool.PostForm(req_url,bodyParams)
	respMarshal ,_:= json.Marshal(resp)
	fmt.Printf("resp %+v",string(respMarshal))
}

