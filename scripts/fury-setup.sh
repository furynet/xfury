
#Setting up constants


FURY_HOME=$HOME/.fury
FURY_SRC=$FURY_HOME/src/fury
COSMOVISOR_SRC=$FURY_HOME/src/cosmovisor

FURY_VERSION="v1.0.1"
COSMOVISOR_VERSION="cosmovisor-v1.0.1"

echo "-----------setting constants---------------"
mkdir -p $FURY_HOME
mkdir -p $FURY_HOME/src
mkdir -p $FURY_HOME/bin
mkdir -p $FURY_HOME/logs
mkdir -p $FURY_HOME/cosmovisor/genesis/bin
mkdir -p $FURY_HOME/cosmovisor/upgrades/


echo "-----------setting environment settings---------------"
sudo apt update
sudo apt upgrade
sudo apt-get update
sudo apt-get upgrade
sudo apt install git build-essential ufw curl jq snapd wget --yes


set -eu

echo "--------------installing golang---------------------------"
curl https://dl.google.com/go/go1.19.1.linux-amd64.tar.gz --output $HOME/go.tar.gz
tar -C $HOME -xzf $HOME/go.tar.gz
rm $HOME/go.tar.gz
export PATH=$PATH:$HOME/go/bin
export GOPATH=$HOME/go
echo "export GOPATH=$HOME/go" >> ~/.bashrc
go version


echo "--------------installing homebrew---------------------------"
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
(echo; echo 'eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"') >> /home/adrian/.profile
eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"

brew install gcc


echo "-------------------installing cosmovisor-----------------------"
git clone -b $COSMOVISOR_VERSION https://github.com/onomyprotocol/onomy-sdk $COSMOVISOR_SRC
cd $COSMOVISOR_SRC
make cosmovisor
cp cosmovisor/cosmovisor $FURY_HOME/bin/cosmovisor


echo "----------------------installing fury---------------"
git clone https://github.com/furynet/xfury
cd xfury
make build && make install
mv ~/xfury/build/xfury $FURY_HOME/cosmovisor/genesis/bin/xfury



echo "-------------------adding binaries to path-----------------------"
chmod +x $FURY_HOME/bin/*
export PATH=$PATH:$FURY_HOME/bin
chmod +x $FURY_HOME/cosmovisor/genesis/bin/*
export PATH=$PATH:$FURY_HOME/cosmovisor/genesis/bin

echo "export PATH=$PATH" >> ~/.bashrc

# set the cosmovisor environments
echo "export DAEMON_HOME=$FURY_HOME/" >> ~/.bashrc
echo "export DAEMON_NAME=fury" >> ~/.bashrc
echo "export DAEMON_RESTART_AFTER_UPGRADE=true" >> ~/.bashrc


PASSWORD="F@nfuryG#n3sis@fury"
GAS_PRICES="0.000025utfury"
CHAIN_ID="fanfury-test-1"
NODE="(xfury tendermint show-node-id)"

xfury init fanfury-testnet --chain-id $CHAIN_ID --staking-bond-denom utfury


# Note: Add an account
yes $PASSWORD | xfury keys add fanfury-1
yes $PASSWORD | xfury add-genesis-acount fanfury-1 1000000000000000000utfury
yes $PASSWORD | xfury gentx fanfury-1 500000000000000000utfury
yes $PASSWORD | xfury collect-gentxs

# Set staking token (both bond_denom and mint_denom)
STAKING_TOKEN="utfury"
FROM="\"bond_denom\": \"stake\""
TO="\"bond_denom\": \"$STAKING_TOKEN\""
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/genesis.json
FROM="\"mint_denom\": \"stake\""
TO="\"mint_denom\": \"$STAKING_TOKEN\""
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/genesis.json

# Set fury token (both for gov min deposit and crisis constant fury)
FEE_TOKEN="utfury"
FROM="\"stake\""
TO="\"$FEE_TOKEN\""
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/genesis.json

# Set reserved bond tokens
RESERVED_BOND_TOKENS="" # example: " \"abc\", \"def\", \"ghi\" "
FROM="\"reserved_bond_tokens\": \[\]"
TO="\"reserved_bond_tokens\": \[$RESERVED_BOND_TOKENS\]"
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/genesis.json

# Set min-gas-prices (using fury token)
FROM="minimum-gas-prices = \"\""
TO="minimum-gas-prices = \"0.000002$FEE_TOKEN\""
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/app.toml

MAX_VOTING_PERIOD="90s" # example: "172800s"
FROM="\"voting_period\": \"172800s\""
TO="\"voting_period\": \"$MAX_VOTING_PERIOD\""
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/genesis.json

# Enable REST API
FROM="enable = false"
TO="enable = true"
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/app.toml

# Enable Swagger docs
FROM="swagger = false"
TO="swagger = true"
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/app.toml

# Broadcast node RPC endpoint
FROM="laddr = \"tcp:\/\/127.0.0.1:26657\""
TO="laddr = \"tcp:\/\/0.0.0.0:26657\""
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/config.toml

# Set timeouts to 1s for shorter block times
sed -i -e "s/timeout_commit = "5s"/timeout_commit = "1s"/g" "$HOME"/.fury/config/config.toml
sed -i -e "s/timeout_propose = "3s"/timeout_propose = "1s"/g" "$HOME"/.fury/config/config.toml

xfury validate-genesis


echo "
###############################################################################
###############################################################################
###############################################################################
###############################################################################
###                                											
###                                											
###                    
###                                											                                											
###                     ~!!~  Congratulations  ~!!~
###               The Fanfury Chain will now be started!!                  						                  											
###               
###                                											
###                                											
###
###############################################################################
###############################################################################
###############################################################################
###############################################################################
"

xfury start --seed-node true
