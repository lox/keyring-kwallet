keyring-kwallet
===============
[![CI](https://github.com/lox/keyring-kwallet/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/lox/keyring-kwallet/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/lox/keyring-kwallet.svg)](https://pkg.go.dev/github.com/lox/keyring-kwallet)

KWallet provider for [`github.com/lox/keyring/v2`](https://github.com/lox/keyring).

## Usage

```bash
go get github.com/lox/keyring-kwallet
```

```go
import (
	"context"

	"github.com/lox/keyring/v2"
	kwallet "github.com/lox/keyring-kwallet"
)

ctx := context.Background()

ring, err := keyring.Open(ctx,
	keyring.WithServiceName("example"),
	keyring.WithProvider(kwallet.Provider()),
)
```

`kwallet.Provider` accepts `AppID` and `Folder` options. On non-Linux
platforms, it returns `keyring.ErrUnavailable` during open.
