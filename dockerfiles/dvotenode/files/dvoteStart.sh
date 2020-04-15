#!/bin/sh

GWARGS="\
${apiAllowPrivate:+ --apiAllowPrivate=${apiAllowPrivate}}\
${apiAllowedAddrs:+ --apiAllowedAddrs=${apiAllowedAddrs}}\
${apiRoute:+ --apiRoute=${apiRoute}}\
${censusApi:+ --censusApi=${censusApi}}\
${dataDir:+ --dataDir=${dataDir}}\
${dev:+ --dev=${dev}}\
${ethBootNodes:+ --ethBootNodes=${ethBootNodes}}\
${ethCensusSync:+ --ethCensusSync=${ethCensusSync}}\
${ethChain:+ --ethChain=${ethChain}}\
${ethChainLightMode:+ --ethChainLightMode=${ethChainLightMode}}\
${ethNoWaitSync:+ --ethNoWaitSync=${ethNoWaitSync}}\
${ethNodePort:+ --ethNodePort=${ethNodePort}}\
${ethProcessDomain:+ --ethProcessDomain=${ethProcessDomain}}\
${ethSigningKey:+ --ethSigningKey=${ethSigningKey}}\
${ethSubscribeOnly:+ --ethSubscribeOnly=${ethSubscribeOnly}}\
${ethTrustedPeers:+ --ethTrustedPeers=${ethTrustedPeers}}\
${fileApi:+ --fileApi=${fileApi}}\
${ipfsNoInit:+ --ipfsNoInit=${ipfsNoInit}}\
${ipfsSyncKey:+ --ipfsSyncKey=${ipfsSyncKey}}\
${ipfsSyncPeers:+ --ipfsSyncPeers=${ipfsSyncPeers}}\
${listenHost:+ --listenHost=${listenHost}}\
${listenPort:+ --listenPort=${listenPort}}\
${logLevel:+ --logLevel=${logLevel}}\
${logOutput:+ --logOutput=${logOutput}}\
${mode:+ --mode=${mode}}\
${resultsApi:+ --resultsApi=${resultsApi}}\
${saveConfig:+ --saveConfig=${saveConfig}}\
${sslDomain:+ --sslDomain=${sslDomain}}\
${vochainCreateGenesis:+ --vochainCreateGenesis=${vochainCreateGenesis}}\
${vochainGenesis:+ --vochainGenesis=${vochainGenesis}}\
${vochainMinerKey:+ --vochainMinerKey=${vochainMinerKey}}\
${vochainNodeKey:+ --vochainNodeKey=${vochainNodeKey}}\
${vochainLogLevel:+ --vochainLogLevel=${vochainLogLevel}}\
${vochainP2PListen:+ --vochainP2PListen=${vochainP2PListen}}\
${vochainPeers:+ --vochainPeers=${vochainPeers}}\
${vochainPublicAddr:+ --vochainPublicAddr=${vochainPublicAddr}}\
${vochainRPCListen:+ --vochainRPCListen=${vochainRPCListen}}\
${vochainSeedMode:+ --vochainSeedMode=${vochainSeedMode}}\
${vochainSeeds:+ --vochainSeeds=${vochainSeeds}}\
${voteApi:+ --voteApi=${voteApi}}\
${w3Enabled:+ --w3Enabled=${w3Enabled}}\
${w3HTTPHost:+ --w3HTTPHost=${w3HTTPHost}}\
${w3HTTPPort:+ --w3HTTPPort=${w3HTTPPort}}\
${w3Route:+ --w3Route=${w3Route}}\
${w3WsHost:+ --w3WsHost=${w3WsHost}}\
${w3WsPort:+ --w3WsPort=${w3WsPort}}\
"

CMD="/app/dvotenode $GWARGS $@"
echo "Executing $CMD"
$CMD

