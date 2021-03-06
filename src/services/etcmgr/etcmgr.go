package etcmgr

import (
	"INServer/src/common/global"
	"INServer/src/common/logger"
	"INServer/src/proto/etc"
	"INServer/src/proto/msg"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gogo/protobuf/proto"
	"gopkg.in/yaml.v3"
)

var Instance *ETC

type (
	ETC struct {
		servers      []*etc.Server
		database     *etc.Database
		basic        *etc.BasicConfig
		zones        []*etc.Zone
		zoneLocation map[int32][]int32 // 每个游戏区都在哪些WorldServer里面
		ok           bool
	}
)

func New() *ETC {
	e := new(ETC)
	if global.CurrentServerID == global.CenterID {
		dir, err := os.Getwd()
		if err != nil {
			logger.Fatal(err)
			return nil
		}
		e.Load(dir + "/etc")
	}
	return e
}

func (e *ETC) Load(path string) {
	var ok bool
	basic := e.loadBasic(path)
	if e.checkBasic(basic) == false {
		logger.Debug("ETC:加载basic失败")
		return
	}
	database := e.loadDatabase(path)
	if e.checkDatabase(database) == false {
		logger.Debug("ETC:加载database失败")
		return
	}
	zones := e.loadZones(path)
	if zones, ok = e.checkZones(zones); ok == false {
		logger.Debug("ETC:加载zones失败")
		return
	}
	servers := e.loadServers(path)
	if servers, ok = e.checkServers(servers); ok == false {
		logger.Debug("ETC:加载servers失败")
		return
	}
	if e.checkAllConfig(basic, database, zones, servers) == false {
		logger.Debug("ETC:检查失败")
		return
	}
	e.basic = basic
	e.database = database
	e.zones = zones
	e.servers = servers
	e.makeConfig()
	e.ok = true
}

func (e *ETC) checkAllConfig(basic *etc.BasicConfig, database *etc.Database, zones []*etc.Zone, servers []*etc.Server) bool {
	return true
}

func (e *ETC) makeConfig() {
	e.zoneLocation = make(map[int32][]int32)
	for _, svr := range e.servers {
		if svr.ServerType == global.WorldServer {
			for _, zone := range svr.ServerConfig.WorldConfig.Zones {
				_, ok := e.zoneLocation[zone.ZoneID]
				if ok == false {
					e.zoneLocation[zone.ZoneID] = make([]int32, 0)
				}
				find := false
				for _, serverID := range e.zoneLocation[zone.ZoneID] {
					if serverID == svr.ServerID {
						find = true
						break
					}
				}
				if find == false {
					e.zoneLocation[zone.ZoneID] = append(e.zoneLocation[zone.ZoneID], svr.ServerID)
				}
			}
		}
	}
}

func (e *ETC) loadServers(path string) []*etc.Server {
	servers := &etc.ServerList{}
	e.load(path, "servers", servers)
	return servers.Servers
}

func (e *ETC) checkServers(servers []*etc.Server) ([]*etc.Server, bool) {
	var maxServerID int32
	maxServerID = -1
	for index, server := range servers {
		if index != int(server.ServerID) {
			proto.Merge(servers[server.ServerID], server)
		}
		if server.ServerID > maxServerID {
			maxServerID = server.ServerID
		}
	}
	return servers[:maxServerID+1], true
}

func (e *ETC) loadDatabase(path string) *etc.Database {
	database := &etc.Database{}
	e.load(path, "database", database)
	return database
}

func (e *ETC) checkDatabase(*etc.Database) bool {
	return true
}

func (e *ETC) load(path string, name string, config proto.Message) {
	lconfig := proto.Clone(config)
	f, err := os.Open(fmt.Sprintf("%s/%s.yaml", path, name))
	if err == nil {
		buf, err := ioutil.ReadAll(f)
		if err != nil {
			logger.Fatal(err)
		}
		err = yaml.Unmarshal(buf, config)
		if err != nil {
			logger.Fatal(err)
		}

		lf, err := os.Open(fmt.Sprintf("%s/local.%s.yaml", path, name))
		if err == nil {
			defer lf.Close()
			lbuf, err := ioutil.ReadAll(lf)
			if err != nil {
				logger.Fatal(err)
			}
			err = yaml.Unmarshal(lbuf, lconfig)
			if err != nil {
				logger.Fatal(err)
			}
			proto.Merge(config, lconfig)
		}
	} else {
		f, err := os.Open(fmt.Sprintf("%s/%s.json", path, name))
		if err != nil {
			logger.Fatal(err)
		}
		buf, err := ioutil.ReadAll(f)
		if err != nil {
			logger.Fatal(err)
		}
		err = json.Unmarshal(buf, config)
		if err != nil {
			logger.Fatal(err)
		}

		lf, err := os.Open(fmt.Sprintf("%s/local.%s.json", path, name))
		if err == nil {
			defer lf.Close()
			lbuf, err := ioutil.ReadAll(lf)
			if err != nil {
				logger.Fatal(err)
			}
			err = json.Unmarshal(lbuf, lconfig)
			if err != nil {
				logger.Fatal(err)
			}
			proto.Merge(config, lconfig)
		}
	}
}

func (e *ETC) loadBasic(path string) *etc.BasicConfig {
	basic := &etc.BasicConfig{}
	e.load(path, "basic", basic)
	return basic
}

func (e *ETC) checkBasic(*etc.BasicConfig) bool {
	return true
}

func (e *ETC) loadZones(path string) []*etc.Zone {
	zones := &etc.ZoneList{}
	e.load(path, "zones", zones)
	return zones.Zones
}

func (e *ETC) checkZones(zones []*etc.Zone) ([]*etc.Zone, bool) {
	var maxZoneID int32
	maxZoneID = -1
	for index, zone := range zones {
		if index != int(zone.ZoneID) {
			proto.Merge(zones[zone.ZoneID], zone)
		}
		if zone.ZoneID > maxZoneID {
			maxZoneID = zone.ZoneID
		}
	}
	return zones[:maxZoneID+1], true
}

func (e *ETC) GetServerType(serverID int32) string {
	if len(e.servers) >= int(serverID) {
		return e.servers[int(serverID)].ServerType
	}
	return global.InvalidServer
}

func (e *ETC) GetServerConfig(serverID int32) *etc.ServerConfig {
	if len(e.servers) >= int(serverID) {
		server := e.servers[int(serverID)]
		switch server.ServerType {
		case global.LoginServer:
			if server.ServerConfig.LoginConfig.Database == nil {
				server.ServerConfig.LoginConfig.Database = e.database
			}
			return server.ServerConfig
		case global.DatabaseServer:
			if server.ServerConfig.DatabaseConfig.Database == nil {
				server.ServerConfig.DatabaseConfig.Database = e.database
			}
			return server.ServerConfig
		default:
			return server.ServerConfig
		}
	}
	return nil
}

func (e *ETC) BasicConfig() *etc.BasicConfig {
	return e.basic
}

func (e *ETC) Database() *etc.Database {
	return e.database
}

func (e *ETC) Servers() []*etc.Server {
	return e.servers
}

func (e *ETC) Zones() []*etc.Zone {
	return e.zones
}

func (e *ETC) GetZoneLocation(zoneID int32) []int32 {
	if list, ok := e.zoneLocation[zoneID]; ok {
		return list
	}
	return nil
}

// OK 配置是否加载完成，对于center是读文件完成，对于node是center返回数据
func (e *ETC) OK() bool {
	return e.ok
}

func (e *ETC) HANDLE_ETC_SYNC_NTF(header *msg.MessageHeader, buffer []byte) {
	ntf := &msg.ETCSyncNTF{}
	err := proto.Unmarshal(buffer, ntf)
	if err != nil {
		logger.Error(err)
		return
	}
	e.servers = ntf.ServerList.Servers
	e.database = ntf.Database
	e.basic = ntf.BasicConfig
	e.zones = ntf.ZoneList.Zones
	e.makeConfig()
	e.ok = true
}

func (e *ETC) GenerateETCSyncNTF() *msg.ETCSyncNTF {
	ntf := &msg.ETCSyncNTF{
		BasicConfig: e.BasicConfig(),
		Database:    e.Database(),
		ServerList: &etc.ServerList{
			Servers: e.Servers(),
		},
		ZoneList: &etc.ZoneList{
			Zones: e.Zones(),
		},
	}
	return ntf
}
