package server

import (
	"github.com/Coinmaxify/Coinmaxify/consensus"
	consensusDev "github.com/Coinmaxify/Coinmaxify/consensus/dev"
	consensusDummy "github.com/Coinmaxify/Coinmaxify/consensus/dummy"
	consensusIBFT "github.com/Coinmaxify/Coinmaxify/consensus/ibft"
	"github.com/Coinmaxify/Coinmaxify/secrets"
	"github.com/Coinmaxify/Coinmaxify/secrets/awsssm"
	"github.com/Coinmaxify/Coinmaxify/secrets/hashicorpvault"
	"github.com/Coinmaxify/Coinmaxify/secrets/local"
)

type ConsensusType string

const (
	DevConsensus   ConsensusType = "dev"
	IBFTConsensus  ConsensusType = "ibft"
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:   consensusDev.Factory,
	IBFTConsensus:  consensusIBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
