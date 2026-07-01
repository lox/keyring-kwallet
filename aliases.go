package kwallet

import "github.com/lox/keyring/v2"

type Item = keyring.Item
type Metadata = keyring.Metadata

const KWalletBackend = keyring.KWalletBackend

var (
	ErrKeyNotFound              = keyring.ErrKeyNotFound
	ErrMetadataNeedsCredentials = keyring.ErrMetadataNeedsCredentials
	ErrUnavailable              = keyring.ErrUnavailable
)
