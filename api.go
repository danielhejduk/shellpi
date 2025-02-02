package main

import (
    "github.com/tidwall/gjson"
    "net/http"
    "io"
    "fmt"
)

type api struct {
    ip string
}

type shelly_data struct {
    a_power int64
    b_power int64
    c_power int64
}

func (api *api) getdata_request() shelly_data {
    resp, err := http.Get(fmt.Sprintf("http://%s/rpc/Shelly.GetStatus", api.ip))
    if err != nil {
        fmt.Printf("[ERROR] %s\n", err.Error())
        return shelly_data{}
    }

    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)

    var shelly shelly_data

    shelly.a_power = gjson.GetBytes(body, "em:0.a_act_power").Int()
    shelly.b_power = gjson.GetBytes(body, "em:0.b_act_power").Int()
    shelly.c_power = gjson.GetBytes(body, "em:0.c_act_power").Int()

    return shelly
}
