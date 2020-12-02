package pas

import (
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"gin-vue-admin/model/pas"
	"testing"

	"gotest.tools/assert"
)

func setUp(t *testing.T) {
	global.GVA_VP = core.Viper()          // 初始化Viper
	global.GVA_LOG = core.Zap()           // 初始化zap日志库
	global.GVA_DB = initialize.Gorm()     // gorm连接数据库
	initialize.MysqlTables(global.GVA_DB) // 初始化表
}

func TestCreateTag(t *testing.T) {
	setUp(t)
	err := CreateTag(pas.Tag{
		Name:     "Test Tag",
		Category: "Category 1",
		Parentid: 0,
	})

	assert.Equal(t, err, nil)
}
