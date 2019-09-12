package vochain

import (
	"fmt"
	"os"
	"path/filepath"

	//"time"

	//abci "github.com/tendermint/tendermint/abci/types"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	codec "github.com/cosmos/cosmos-sdk/codec"
	//abci "github.com/tendermint/tendermint/abci/types"
	cfg "github.com/tendermint/tendermint/config"
	tmflags "github.com/tendermint/tendermint/libs/cli/flags"
	cmn "github.com/tendermint/tendermint/libs/common"
	tlog "github.com/tendermint/tendermint/libs/log"
	nm "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	privval "github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/proxy"
	tmtypes "github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
	dbm "github.com/tendermint/tm-db"
	vlog "gitlab.com/vocdoni/go-dvote/log"
	vochain "gitlab.com/vocdoni/go-dvote/vochain/app"

	//test "gitlab.com/vocdoni/go-dvote/vochain/test"
	vochaintypes "gitlab.com/vocdoni/go-dvote/vochain/types"
)

// Start starts a new vochain validator node
func Start(configFilePath string, db *dbm.GoLevelDB) (*vochain.BaseApplication, *nm.Node) {
	// PUT ON GATEWAY CONFIG
	/*
		flag.StringVar(&configFile, "config", "/home/jordi/vocdoni/go-dvote/vochain/config/config.toml", "Path to config.toml")
		flag.StringVar(&appdbName, "appdbname", "vochaindb", "Application database name")
		flag.StringVar(&appdbDir, "appdbdir", "/home/jordi/vocdoni/go-dvote/vochain/data", "Path where the application database will be located")
	*/
	// create application db
	vlog.Info("Initializing Vochain")

	// creating new vochain app
	app := vochain.NewBaseApplication(db)
	//flag.Parse()
	vlog.Info("Creating node and application")
	node, err := newTendermint(app, configFilePath)
	if err != nil {
		vlog.Info(err)
		return app, node
	}
	node.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
	return app, node
}

// we need to set init (first time validators and oracles)
func newTendermint(app *vochain.BaseApplication, configFile string) (*nm.Node, error) {
	// create node config
	config := cfg.DefaultConfig()
	config.RootDir = filepath.Dir(filepath.Dir(configFile))
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper failed to read config file")
	}
	if err := viper.Unmarshal(config); err != nil {
		return nil, errors.Wrap(err, "viper failed to unmarshal config")
	}
	if err := config.ValidateBasic(); err != nil {
		return nil, errors.Wrap(err, "config is invalid")
	}

	// create data dir
	newpath := filepath.Join(config.RootDir, "data")
	os.MkdirAll(newpath, os.ModePerm)

	// create logger
	logger := tlog.NewTMLogger(tlog.NewSyncWriter(os.Stdout))
	var err error
	//config.LogLevel = "none"
	logger, err = tmflags.ParseLogLevel(config.LogLevel, logger, cfg.DefaultLogLevel())
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse log level")
	}

	// read or create private validator
	pv := privval.LoadOrGenFilePV(
		config.PrivValidatorKeyFile(),
		config.PrivValidatorStateFile(),
	)

	// read or create node key
	nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
	if err != nil {
		return nil, errors.Wrap(err, "failed to load node's key")
	}

	// read or create genesis file
	genFile := config.GenesisFile()
	if cmn.FileExists(genFile) {
		vlog.Info("Found genesis file", "path", genFile)
	} else {
		vlog.Info("Creating genesis file")
		genDoc := tmtypes.GenesisDoc{
			ChainID:         "0x1",
			GenesisTime:     tmtime.Now(),
			ConsensusParams: tmtypes.DefaultConsensusParams(),
		}

		// create app state getting validators and oracle keys from eth
		// one oracle needs to exist
		state := &vochaintypes.State{
			Oracles:    getOraclesFromEth(), // plus existing oracle ?
			Validators: getValidatorsFromEth(*pv),
			Processes:  make(map[string]*vochaintypes.Process, 0),
		}

		// set validators from eth smart contract
		genDoc.Validators = state.Validators

		// amino marshall state
		genDoc.AppState = codec.Cdc.MustMarshalJSON(*state)

		// save genesis
		if err := genDoc.SaveAs(genFile); err != nil {
			panic(fmt.Sprintf("Cannot load or generate genesis file: %v", err))
		}
		logger.Info("Generated genesis file", "path", genFile)
		vlog.Info("genesis file: %+v", genFile)
	}

	// create node
	node, err := nm.NewNode(
		config,
		pv,                               // the node val
		nodeKey,                          // node val key
		proxy.NewLocalClientCreator(app), // Note we use proxy.NewLocalClientCreator here to create a local client instead of one communicating through a socket or gRPC.
		nm.DefaultGenesisDocProviderFunc(config),
		nm.DefaultDBProvider,
		nm.DefaultMetricsProvider(config.Instrumentation),
		logger)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new Tendermint node")
	}

	return node, nil
}

// temp function
func getValidatorsFromEth(nodeKey privval.FilePV) []tmtypes.GenesisValidator {
	// TODO oracle doing stuff
	// oracle returns a list of validators... then
	list := make([]tmtypes.GenesisValidator, 0)
	list = append(list, tmtypes.GenesisValidator{

		Address: nodeKey.GetPubKey().Address(),
		PubKey:  nodeKey.GetPubKey(),
		Power:   10,
	},
	)
	return list
}

// temp func
func getOraclesFromEth() []string {
	// TODO oracle doing stuff
	// oracle returns a list of trusted oracles
	//"0xF904848ea36c46817096E94f932A9901E377C8a5"
	list := make([]string, 0)
	list = append(list, "0xF904848ea36c46817096E94f932A9901E377C8a5")
	vlog.Infof("Oracles return: %v", list[0])
	return list
}