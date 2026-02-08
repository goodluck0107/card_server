package wuhan

import (
	"bytes"
	"cmp"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/open-source/game/chess.git/pkg/models"
	public "github.com/open-source/game/chess.git/pkg/static"
	syslog "github.com/open-source/game/chess.git/pkg/xlog"
	"net/http"
	"runtime/debug"
	"slices"
)

func SuperOpt(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                               // 允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept") // header的类型
	w.Header().Set("content-type", "application/json")
	defer func() {
		x := recover()
		if x != nil {
			syslog.Logger().Errorln(x, string(debug.Stack()))
		}
	}() // 返回数据格式是json

	data := req.FormValue("token")
	if data != "d16783f09ea7" {
		w.Write([]byte("你在干吗"))
	} else {
		superId := public.HF_Atoi(req.FormValue("super"))
		if superId <= 0 {
			w.Write([]byte("你在干吗"))
		}
		opt := req.FormValue("opt")
		if opt == "1" {
			rate := public.HF_Atoi(req.FormValue("rate"))
			if rate < 0 || rate > 100 {
				w.Write([]byte("你在干吗"))
			}
			err := GetDBMgr().GetDBrControl().RedisV2.HSet("superman", req.FormValue("super"), req.FormValue("rate")).Err()
			if err != nil {
				syslog.Logger().Error(err)
				w.Write([]byte("数据库异常"))
			} else {
				w.Write([]byte("添加成功"))
			}
		} else if opt == "0" {
			err := GetDBMgr().GetDBrControl().RedisV2.HDel("superman", req.FormValue("super")).Err()
			if err != nil {
				syslog.Logger().Error(err)
				w.Write([]byte("数据库异常"))
			} else {
				w.Write([]byte("删除成功"))
			}
		} else {
			w.Write([]byte("你在干吗"))
		}
	}
}

func FakerOpt(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                               // 允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept") // header的类型
	w.Header().Set("content-type", "application/json")
	defer func() {
		x := recover()
		if x != nil {
			syslog.Logger().Errorln(x, string(debug.Stack()))
		}
	}() // 返回数据格式是json

	data := req.FormValue("token")
	if data != "d16783f09ea7" {
		w.Write([]byte("你在干吗"))
	} else {
		fakerId := public.HF_Atoi(req.FormValue("faker"))
		if fakerId <= 0 {
			w.Write([]byte("你在干吗"))
		}
		opt := req.FormValue("opt")
		if opt == "1" {
			err := GetDBMgr().GetDBrControl().RedisV2.SAdd("faker_admin", fakerId).Err()
			if err != nil {
				syslog.Logger().Error(err)
				w.Write([]byte("数据库异常"))
			} else {
				w.Write([]byte("添加战绩查看权限成功"))
			}
		} else if opt == "0" {
			err := GetDBMgr().GetDBrControl().RedisV2.SRem("faker_admin", fakerId).Err()
			if err != nil {
				syslog.Logger().Error(err)
				w.Write([]byte("数据库异常"))
			} else {
				w.Write([]byte("删除战绩查看权限成功"))
			}
		} else {
			w.Write([]byte("你在干吗"))
		}
	}
}

func OwnerOpt(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                               // 允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept") // header的类型
	w.Header().Set("content-type", "application/json")
	defer func() {
		x := recover()
		if x != nil {
			syslog.Logger().Errorln(x, string(debug.Stack()))
		}
	}() // 返回数据格式是json

	data := req.FormValue("token")
	if data != "d16783f09ea7" {
		w.Write([]byte("你在干吗?"))
	} else {
		hid := public.HF_Atoi64(req.FormValue("hid"))
		if hid <= 0 {
			w.Write([]byte("你在干吗??"))
			return
		}

		var house models.House
		err := GetDBMgr().GetDBmControl().First(&house, "hid = ?", hid).Error
		if err != nil {
			syslog.Logger().Error(err)
			w.Write([]byte("数据库异常:" + err.Error()))
			return
		}

		key := fmt.Sprintf("houseOwner:%d:%d", house.HId, house.UId)
		cli := GetDBMgr().GetDBrControl().RedisV2
		date := req.FormValue("date")
		res := make(map[string]string)
		if date == "" {
			res, err = cli.HGetAll(key).Result()
		} else {
			res[date], err = cli.HGet(key, date).Result()
		}
		syslog.Logger().Warningf("owner query: %s, result: %v", key, res)
		if err != nil && !errors.Is(err, redis.Nil) {
			syslog.Logger().Error(err)
			w.Write([]byte("数据库异常:" + err.Error()))
			return
		}

		type Res struct {
			Date   string `json:"date"`
			Income string `json:"income"`
		}

		resSlice := make([]Res, 0, len(res))
		for k, v := range res {
			resSlice = append(resSlice, Res{
				Date:   k,
				Income: v,
			})
		}
		slices.SortFunc(resSlice, func(a, b Res) int {
			return cmp.Compare(a.Date, b.Date)
		})

		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("当前圈号: %d\n", house.HId))
		buf.WriteString(fmt.Sprintf("当前圈名: %s\n", house.Name))
		buf.WriteString(fmt.Sprintf("当前房主: %d\n", house.UId))
		buf.WriteString("收益情况: \n")
		for _, v := range resSlice {
			buf.WriteString(fmt.Sprintf("\t日期:%s => 收益:%s\n", v.Date, v.Income))
		}
		w.Write([]byte(buf.String()))
		return
	}
}
