package pas

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/pas"
	"testing"

	"gotest.tools/assert"
)

func TestCreateKpi(t *testing.T) {

	setUp(t)
	var tags []pas.Tag
	global.GVA_DB.Find(&tags)
	err := CreateKpi(pas.Kpi{
		Name:        "Test",
		Description: "This is a Test KPI",
		Status:      true,
		Category:    "TEST",
		Tags:        tags,
	})
	assert.Equal(t, err, nil)
}
