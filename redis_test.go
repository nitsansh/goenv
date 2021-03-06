package goenv

import (
	"testing"
)

func TestGetRedis(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "redis", "nil")
	host, port, db := goenv.GetRedis()
	if host != "ecr" || port != "340" || db != 16 {
		t.Error("redis != ecr, 340, 16")
	}
}

func TestGetRedisNotFound(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "nonexistent", "nil")
	host, port, db := goenv.GetRedis()
	if host != "localhost" || port != "6379" || db != -1 {
		t.Error("redis != localhost, 6379, -1")
	}
}

func TestGetNamedRedis(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "redis", "nil")
	host, port, db := goenv.GetNamedRedis("custom")
	if host != "ruo" || port != "114" || db != 81 {
		t.Error("custom != ruo, 114, 81")
	}
}

func TestGetNamedRedisNotFound(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "redis", "nil")
	host, port, db := goenv.GetNamedRedis("nonexistent")
	if host != "localhost" || port != "6379" || db != -1 {
		t.Error("nonexistent != localhost, 6379, -1")
	}
}
