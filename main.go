package main

import (
	"INServer/src/common/global"
	"INServer/src/common/logger"
	"INServer/src/common/profiler"
	"INServer/src/modules/center"
	"INServer/src/modules/database"
	"INServer/src/modules/etcmgr"
	"INServer/src/modules/gate"
	"INServer/src/modules/gps"
	"INServer/src/modules/login"
	"INServer/src/modules/node"
	"INServer/src/modules/web"
	"INServer/src/modules/world"
	"flag"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var serverID = flag.Int("id", -1, "本服务器ID(范围0~65535)")
var centerIP = flag.String("cip", "127.0.0.1", "中心服务器IP")
var profile = flag.Bool("profile", true, "开启性能监控")

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	runtime.GOMAXPROCS(1)
	flag.Parse()
	global.ServerID = int32(*serverID)
	global.CenterIP = *centerIP

	if global.ServerID == -1 {
		log.Fatalln("必须使用参数(-id ?)指定本服务器ID")
	} else if global.ServerID > global.SERVER_ID_MAX || global.ServerID < 0 {
		log.Fatalln("服务器ID范围0~999")
	}

	logger.Setup()

	if *profile {
		profiler.Start()
	}

	if global.ServerID == global.CenterID {
		go startCenter()
	} else {
		go startNode()
	}

	for {
		sig := <-sigs
		logger.Info(sig)
		if sig.String() == "interrupt" {
			break
		}
	}

	logger.Info("Shut Down!")
}

func startCenter() {
	global.ServerType = global.CenterServer
	dir, err := os.Getwd()
	if err != nil {
		logger.Fatal(err)
		return
	}
	etcmgr.Instance = etcmgr.New()
	etcmgr.Instance.Load(dir + "/etc")
	center.Instance = center.New()
	center.Instance.Start()
}

func startNode() {
	node.Instance = node.New()
	node.Instance.Prepare()
	prepareServer()
	node.Instance.Start()
}

func prepareServer() {
	switch global.ServerType {
	case global.GateServer:
		gate.Instance = gate.New()
		gate.Instance.Start()
		break
	case global.LoginServer:
		login.Instance = login.New()
		login.Instance.Start()
	case global.DatabaseServer:
		database.Instance = database.New()
		database.Instance.Start()
	case global.WebServer:
		web.Instance = web.New()
		web.Instance.Start()
	case global.WorldServer:
		world.Instance = world.New()
		world.Instance.Start()
	case global.GPSServer:
		gps.Instance = gps.New()
		gps.Instance.Start()
	default:
		logger.Fatal("不支持的服务器类型:" + global.ServerType)
	}
}
