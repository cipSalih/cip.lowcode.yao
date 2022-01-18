package importer

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaoapp/gou"
	"github.com/yaoapp/xiang/config"
)

func TestProcessMapping(t *testing.T) {
	simple := filepath.Join(config.Conf.Root, "imports", "assets", "simple.xlsx")
	args := []interface{}{"order", simple}
	response := gou.NewProcess("xiang.import.Mapping", args...).Run()
	_, ok := response.(*Mapping)
	assert.True(t, ok)
}

func TestProcessMappingSetting(t *testing.T) {
	args := []interface{}{"order"}
	response := gou.NewProcess("xiang.import.MappingSetting", args...).Run()
	assert.Nil(t, response)
}

func TestProcessData(t *testing.T) {
	simple := filepath.Join(config.Conf.Root, "imports", "assets", "simple.xlsx")
	mapping := gou.NewProcess("xiang.import.Mapping", "order", simple).Run()
	args := []interface{}{"order", simple, 1, 2, mapping}
	response := gou.NewProcess("xiang.import.Data", args...).Run()
	_, ok := response.(map[string]interface{})
	assert.True(t, ok)
}

func TestProcessDataSetting(t *testing.T) {
	args := []interface{}{"order"}
	response := gou.NewProcess("xiang.import.DataSetting", args...).Run()
	assert.Nil(t, response)
}

func TestProcessRun(t *testing.T) {
	simple := filepath.Join(config.Conf.Root, "imports", "assets", "simple.xlsx")
	mapping := gou.NewProcess("xiang.import.Mapping", "order", simple).Run()
	args := []interface{}{"order", simple, mapping}
	response := gou.NewProcess("xiang.import.Run", args...).Run()
	_, ok := response.(map[string]int)
	assert.True(t, ok)
}

func TestProcessRules(t *testing.T) {
	args := []interface{}{"order"}
	response := gou.NewProcess("xiang.import.Rules", args...).Run()
	assert.Nil(t, response)
}
