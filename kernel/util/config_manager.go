package util

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v3"

	"github.com/billadm/kernel/constant"
	"github.com/billadm/kernel/logger"
)

var (
	configManager     *ConfigManager
	configManagerOnce sync.Once
)

// ConfigManager 管理配置，提供参数查询功能
type ConfigManager struct {
	configMap map[string]string
}

func GetConfigManager() *ConfigManager {
	if configManager != nil {
		return configManager
	}

	configManagerOnce.Do(func() {
		configManager = newConfigManager()
	})

	return configManager
}

// 创建一个新的ConfigManager实例
func newConfigManager() *ConfigManager {
	globalConfig := make(map[string]string)
	modeConfig := make(map[string]string)

	filePath := filepath.Join(GetConfDir(), constant.YamlFile)
	data, err := os.ReadFile(filePath)
	if err != nil {
		logger.Error("读取文件%s失败: %v", filePath, err)
		panic(fmt.Errorf("读取文件%s失败: %v", filePath, err))
	}

	// 解析YAML内容
	var fileContent map[string]map[string]string
	if err := yaml.Unmarshal(data, &fileContent); err != nil {
		logger.Error("解析YAML文件%s失败: %v", filePath, err)
		panic(fmt.Errorf("解析YAML文件%s失败: %v", filePath, err))
	}

	// 合并global配置
	if globalSection, exists := fileContent["global"]; exists {
		mergeMaps(globalConfig, globalSection)
	}

	// 合并mode配置
	if modeSection, exists := fileContent[constant.Mode]; exists {
		mergeMaps(modeConfig, modeSection)
	}

	// 合并全局和模式配置（模式配置优先）
	finalConfig := make(map[string]string)
	mergeMaps(finalConfig, globalConfig)
	mergeMaps(finalConfig, modeConfig)

	return &ConfigManager{
		configMap: finalConfig,
	}
}

// Get 获取配置参数，不存在时返回默认值
func (cm *ConfigManager) Get(key, defaultValue string) string {
	if val, exists := cm.configMap[key]; exists {
		return val
	}
	return defaultValue
}

// mergeMaps 将源map合并到目标map（源map优先级更高）
func mergeMaps(dest, src map[string]string) {
	for k, v := range src {
		dest[k] = v
	}
}
