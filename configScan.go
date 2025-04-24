package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

func getConfig() map[string]interface{} {
	var config map[string]interface{}
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err)
	}
	return config
}
func getBotToken() string {
	return getChapter("telegram")["bot_token"].(string)

}
func getChatID() int {
	return int(getChapter("telegram")["chat_id"].(int64))
}
func getCheckInterval() int64 {
	return getChapter("settings")["check_interval"].(int64)
}
func getChapter(nameChapter string) map[string]interface{} {
	return getConfig()[nameChapter].(map[string]interface{})
}
func getUrls() []interface{} {
	return getChapter("sites")["urls"].([]interface{})
}
func getTimeout() int64 {
	return getChapter("settings")["timeout"].(int64)
}
