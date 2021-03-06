package repo

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/meshplus/bitxhub-kit/fileutil"

	"github.com/meshplus/bitxhub-kit/types"
	"github.com/spf13/viper"
)

// Config represents the necessary config data for starting pier
type Config struct {
	RepoRoot string
	Title    string   `toml:"title" json:"title"`
	Port     Port     `toml:"port" json:"port"`
	Bitxhub  Bitxhub  `toml:"bitxhub" json:"bitxhub"`
	Log      Log      `toml:"log" json:"log"`
	Appchain Appchain `toml:"appchain" json:"appchain"`
}

// Port are ports providing http and pprof service
type Port struct {
	Http  int64 `toml:"http" json:"http"`
	PProf int64 `toml:"pprof" json:"pprof"`
}

// Bitxhub are config about bitxhub
type Bitxhub struct {
	Addr       string   `toml:"addr" json:"addr"`
	Validators []string `toml:"validators" json:"validators"`
}

// GetValidators gets validator address of bitxhub
func (bxh *Bitxhub) GetValidators() []types.Address {
	validators := make([]types.Address, 0)
	for _, v := range bxh.Validators {
		validators = append(validators, types.String2Address(v))
	}
	return validators
}

// Log are config about log
type Log struct {
	Dir          string `toml:"dir" json:"dir"`
	Filename     string `toml:"filename" json:"filename"`
	ReportCaller bool   `mapstructure:"report_caller"`
	Level        string `toml:"level" json:"level"`
}

// Appchain are configs about appchain
type Appchain struct {
	Plugin string `toml:"plugin" json:"plugin"`
	Config string `toml:"config" json:"config"`
}

const (
	// Items
	LogLevel = "log.level"
)

// DefaultConfig returns config with default value
func DefaultConfig() *Config {
	return &Config{
		RepoRoot: ".pier",
		Title:    "pier configuration file",
		Port: Port{
			Http:  8987,
			PProf: 44555,
		},
		Bitxhub: Bitxhub{
			Addr: "localhost:60011",
			Validators: []string{
				"0x000f1a7a08ccc48e5d30f80850cf1cf283aa3abd",
				"0xe93b92f1da08f925bdee44e91e7768380ae83307",
				"0xb18c8575e3284e79b92100025a31378feb8100d6",
				"0x856E2B9A5FA82FD1B031D1FF6863864DBAC7995D",
			},
		},
		Log: Log{
			Dir:          "logs",
			Filename:     "pier.log",
			ReportCaller: false,
			Level:        "info",
		},
		Appchain: Appchain{
			Plugin: "fabric-client-1.4.so",
			Config: "fabric",
		},
	}
}

// UnmarshalConfig read from config files under config path
func UnmarshalConfig(repoRoot string) (*Config, error) {
	configPath := filepath.Join(repoRoot, ConfigName)

	if !fileutil.Exist(configPath) {
		return nil, fmt.Errorf("please initialize pier firstly")
	}

	viper.SetConfigFile(configPath)
	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("PIER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := DefaultConfig()

	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	config.RepoRoot = repoRoot

	return config, nil
}
