package config

import (
	"crypto/ecdsa"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/yaml.v2"
)

var (
	SOLVER_PK  *ecdsa.PrivateKey
	SOLVER_EOA common.Address
)

// Config is the configuration struct for the bot.
// It is loaded from a YAML file and can be overridden by CLI flags and ENV vars.
// Precedence: ENV vars > CLI flags > YAML file
type Config struct {
	Mode string `yaml:"mode"`

	LogLevel           string `yaml:"log_level"`
	EthRpcUrl          string `yaml:"eth_rpc_url"`
	BloxrouteGrpcUrl   string `yaml:"bloxroute_grpc_url"`
	BloxrouteAuthToken string `yaml:"bloxroute_auth_token"`

	PoolsConfigFile  string `yaml:"pools_config_file"`
	PoolsBinFile     string `yaml:"pools_bin_file"`
	SwapPathsBinFile string `yaml:"swap_paths_bin_file"`

	MempoolConnectionEthMethod string `yaml:"mempool_connection_eth_method"`

	InterestingTokens []string `yaml:"interesting_tokens"`

	FastlaneOnlineAddress    string `yaml:"fastlane_online_address"`
	SolverContractAddress    string `yaml:"solver_contract_address"`
	AtlasAddress             string `yaml:"atlas_address"`
	AtlasVerificationAddress string `yaml:"atlas_verification_address"`
	WethAddress              string `yaml:"weth_address"`

	FrontrunGasMarginGwei int64 `yaml:"frontrun_gas_margin_gwei"`
	ProfitMarginx10000    int64 `yaml:"profit_margin_x10000"`

	MaxConcurrentRpcCalls int `yaml:"max_concurrent_rpc_calls"`
}

func NewConfig() *Config {
	var conf Config

	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(configFile, &conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Parse CLI args
	fs := flag.NewFlagSet("BotConfig", flag.ExitOnError)
	fs.StringVar(&conf.Mode, "mode", conf.Mode, "Mode")
	fs.StringVar(&conf.LogLevel, "log_level", conf.LogLevel, "Log level")
	fs.StringVar(&conf.EthRpcUrl, "eth_rpc_url", conf.EthRpcUrl, "Ethereum RPC URL")
	fs.StringVar(&conf.BloxrouteGrpcUrl, "bloxroute_grpc_url", conf.BloxrouteGrpcUrl, "Bloxroute gRPC URL")
	//bloxroute auth token can only be passed via env var
	fs.StringVar(&conf.PoolsConfigFile, "pools_config_file", conf.PoolsConfigFile, "Pools config file")
	fs.StringVar(&conf.PoolsBinFile, "pools_bin_file", conf.PoolsBinFile, "Pools bin file")
	fs.StringVar(&conf.SwapPathsBinFile, "swap_paths_bin_file", conf.SwapPathsBinFile, "Swap paths bin file")
	fs.StringVar(&conf.MempoolConnectionEthMethod, "mempool_connection_eth_method", conf.MempoolConnectionEthMethod, "Mempool connection ETH method")
	fs.StringVar(&conf.FastlaneOnlineAddress, "fastlane_online_address", conf.FastlaneOnlineAddress, "Fastlane online address")
	fs.StringVar(&conf.SolverContractAddress, "solver_contract_address", conf.SolverContractAddress, "Solver contract address")
	fs.StringVar(&conf.AtlasAddress, "atlas_address", conf.AtlasAddress, "Atlas address")
	fs.StringVar(&conf.AtlasVerificationAddress, "atlas_verification_address", conf.AtlasVerificationAddress, "Atlas verification address")
	fs.StringVar(&conf.WethAddress, "weth_address", conf.WethAddress, "WETH address")
	fs.Int64Var(&conf.FrontrunGasMarginGwei, "frontrun_gas_margin_gwei", conf.FrontrunGasMarginGwei, "Frontrun gas margin in gwei")
	fs.Int64Var(&conf.ProfitMarginx10000, "profit_margin_x10000", conf.ProfitMarginx10000, "Profit margin x10000")
	fs.IntVar(&conf.MaxConcurrentRpcCalls, "max_concurrent_rpc_calls", conf.MaxConcurrentRpcCalls, "Max concurrent RPC calls")
	fs.Parse(os.Args[1:])

	// Override with ENV vars
	if env := os.Getenv("SOLVER_PK"); env != "" {
		if pk, err := crypto.HexToECDSA(env); err == nil {
			SOLVER_PK = pk
			SOLVER_EOA = crypto.PubkeyToAddress(SOLVER_PK.PublicKey)
		} else {
			log.Fatalf("error: invalid SOLVER_PK env var")
		}
	}
	if env := os.Getenv("MODE"); env != "" {
		conf.Mode = env
	}
	if env := os.Getenv("LOG_LEVEL"); env != "" {
		conf.LogLevel = env
	}
	if env := os.Getenv("ETH_RPC_URL"); env != "" {
		conf.EthRpcUrl = env
	}
	if env := os.Getenv("BLOXROUTE_GRPC_URL"); env != "" {
		conf.BloxrouteGrpcUrl = env
	}
	if env := os.Getenv("BLOXROUTE_AUTH_TOKEN"); env != "" {
		conf.BloxrouteAuthToken = env
	}
	if env := os.Getenv("POOLS_CONFIG_FILE"); env != "" {
		conf.PoolsConfigFile = env
	}
	if env := os.Getenv("POOLS_BIN_FILE"); env != "" {
		conf.PoolsBinFile = env
	}
	if env := os.Getenv("SWAP_PATHS_BIN_FILE"); env != "" {
		conf.SwapPathsBinFile = env
	}
	if env := os.Getenv("MEMPOOL_CONNECTION_ETH_METHOD"); env != "" {
		conf.MempoolConnectionEthMethod = env
	}
	if env := os.Getenv("FASTLANE_ONLINE_ADDRESS"); env != "" {
		conf.FastlaneOnlineAddress = env
	}
	if env := os.Getenv("SOLVER_CONTRACT_ADDRESS"); env != "" {
		conf.SolverContractAddress = env
	}
	if env := os.Getenv("ATLAS_ADDRESS"); env != "" {
		conf.AtlasAddress = env
	}
	if env := os.Getenv("ATLAS_VERIFICATION_ADDRESS"); env != "" {
		conf.AtlasVerificationAddress = env
	}
	if env := os.Getenv("WETH_ADDRESS"); env != "" {
		conf.WethAddress = env
	}
	if env := os.Getenv("FRONTRUN_GAS_MARGIN_GWEI"); env != "" {
		if margin, err := strconv.Atoi(env); err == nil {
			conf.FrontrunGasMarginGwei = int64(margin)
		} else {
			log.Fatalf("error: invalid FRONTRUN_GAS_MARGIN_GWEI env var")
		}
	}
	if env := os.Getenv("PROFIT_MARGIN_X10000"); env != "" {
		if margin, err := strconv.Atoi(env); err == nil {
			conf.ProfitMarginx10000 = int64(margin)
		} else {
			log.Fatalf("error: invalid PROFIT_MARGIN_X10000 env var")
		}
	}
	if env := os.Getenv("MAX_CONCURRENT_RPC_CALLS"); env != "" {
		if max, err := strconv.Atoi(env); err == nil {
			conf.MaxConcurrentRpcCalls = max
		} else {
			log.Fatalf("error: invalid MAX_CONCURRENT_RPC_CALLS env var")
		}
	}

	return &conf
}

func (c *Config) PoolsConfig() ([]*PoolsConfig, error) {
	return parsePoolsConfig(c.PoolsConfigFile)
}

func (c *Config) GetInterestingTokens() []common.Address {
	var tokens []common.Address
	for _, token := range c.InterestingTokens {
		tokens = append(tokens, common.HexToAddress(token))
	}
	return tokens
}
