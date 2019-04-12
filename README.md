# Block Atlas by Trust Wallet

[![Build Status](https://dev.azure.com/TrustWallet/Trust%20Backend/_apis/build/status/TrustWallet.blockatlas?branchName=master)](https://dev.azure.com/TrustWallet/Trust%20Backend/_build/latest?definitionId=25&branchName=master)

Clean explorer API for crypto currencies.

__Supported Coins__

<a href="https://binance.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/714.png" width="32" /></a>
<a href="https://nimiq.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/242.png" width="32" /></a>
<a href="https://ripple.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/144.png" width="32" /></a>
<a href="https://stellar.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/148.png" width="32" /></a>
<a href="https://kin.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/2017.png" width="32" /></a>
<a href="https://tezos.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/1729.png" width="32" /></a>
<a href="https://aion.network" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/425.png" width="32" /></a>
<a href="https://ethereum.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/60.png" width="32" /></a>
<a href="https://ethereumclassic.github.io" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/61.png" width="32" /></a>
<a href="https://poa.network" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/178.png" width="32" /></a>
<a href="https://callisto.network" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/820.png" width="32" /></a>
<a href="https://gochain.io" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/6060.png" width="32" /></a>
<a href="https://wanchain.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/5718350.png" width="32" /></a>
<a href="https://thundercore.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/1001.png" width="32" /></a>

## Setup

#### Quick start

Deploy it in less than 30 seconds!

__From Source__ (Go Toolchain required)

```shell
go get -u github.com/trustwallet/blockatlas
~/go/bin/blockatlas
```

__With Docker__

`docker run -it -p 8420:8420 trustwallet/blockatlas`

## Configuration

Block Atlas can run just fine without configuration.

If you want to use custom RPC endpoints, or enable coins without public RPC (like Nimiq),
you can configure Block Atlas over `config.yml` or environment variables.

__Config File__

By default, `config.yml` is loaded from the working directory.
Live reload is supported across the app.

Example (`config.yml`):
```yaml
nimiq:
  api: http://localhost:8648
#...
```

__Environment__

The rest gets loaded from the environment variables.
Every config option is available under the `ATLAS_` prefix.

Example:
```shell
ATLAS_NIMIQ_API=http://localhost:8648 \
blockatlas
```

#### Supported platforms

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://www.heroku.com/deploy/?template=https://github.com/TrustWallet/blockatlas)

[![Docker](https://img.shields.io/docker/cloud/build/trustwallet/blockatlas.svg?style=for-the-badge)](https://hub.docker.com/r/trustwallet/blockatlas)
