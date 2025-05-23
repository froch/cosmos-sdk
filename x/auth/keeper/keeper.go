package keeper

import (
	"context"
	"errors"
	"fmt"
	"time"

	gogotypes "github.com/cosmos/gogoproto/types"

	"cosmossdk.io/collections"
	"cosmossdk.io/collections/indexes"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/log"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeperI is the interface contract that x/auth's keeper implements.
type AccountKeeperI interface {
	// Return a new account with the next account number and the specified address. Does not save the new account to the store.
	NewAccountWithAddress(context.Context, sdk.AccAddress) sdk.AccountI

	// Return a new account with the next account number. Does not save the new account to the store.
	NewAccount(context.Context, sdk.AccountI) sdk.AccountI

	// Check if an account exists in the store.
	HasAccount(context.Context, sdk.AccAddress) bool

	// Retrieve an account from the store.
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI

	// Set an account in the store.
	SetAccount(context.Context, sdk.AccountI)

	// Remove an account from the store.
	RemoveAccount(context.Context, sdk.AccountI)

	// Iterate over all accounts, calling the provided function. Stop iteration when it returns true.
	IterateAccounts(context.Context, func(sdk.AccountI) bool)

	// Fetch the public key of an account at a specified address
	GetPubKey(context.Context, sdk.AccAddress) (cryptotypes.PubKey, error)

	// Fetch the sequence of an account at a specified address.
	GetSequence(context.Context, sdk.AccAddress) (uint64, error)

	// Fetch the next account number, and increment the internal counter.
	NextAccountNumber(context.Context) uint64

	// GetModulePermissions fetches per-module account permissions
	GetModulePermissions() map[string]types.PermissionsForAddress

	// AddressCodec returns the account address codec.
	AddressCodec() address.Codec
}

func NewAccountIndexes(sb *collections.SchemaBuilder) AccountsIndexes {
	return AccountsIndexes{
		Number: indexes.NewUnique(
			sb, types.AccountNumberStoreKeyPrefix, "account_by_number", collections.Uint64Key, sdk.AccAddressKey,
			func(_ sdk.AccAddress, v sdk.AccountI) (uint64, error) {
				return v.GetAccountNumber(), nil
			},
		),
	}
}

type AccountsIndexes struct {
	// Number is a unique index that indexes accounts by their account number.
	Number *indexes.Unique[uint64, sdk.AccAddress, sdk.AccountI]
}

func (a AccountsIndexes) IndexesList() []collections.Index[sdk.AccAddress, sdk.AccountI] {
	return []collections.Index[sdk.AccAddress, sdk.AccountI]{
		a.Number,
	}
}

// AccountKeeper encodes/decodes accounts using the go-amino (binary)
// encoding/decoding library.
type AccountKeeper struct {
	addressCodec address.Codec

	storeService store.KVStoreService
	cdc          codec.BinaryCodec
	permAddrs    map[string]types.PermissionsForAddress
	bech32Prefix string

	// enableUnorderedTxs enables unordered transaction support.
	// This boolean helps sigverify ante handlers to determine if they should process unordered transactions.
	enableUnorderedTxs bool

	// The prototypical AccountI constructor.
	proto func() sdk.AccountI

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string

	// State
	Schema          collections.Schema
	Params          collections.Item[types.Params]
	AccountNumber   collections.Sequence
	Accounts        *collections.IndexedMap[sdk.AccAddress, sdk.AccountI, AccountsIndexes]
	UnorderedNonces collections.KeySet[collections.Pair[int64, []byte]]
}

type InitOption func(*AccountKeeper)

// WithUnorderedTransactions enables unordered transaction support.
// When true, sigverify ante handlers will validate and process unordered transactions.
// When false, sigverify ante handlers will reject unordered transactions.
func WithUnorderedTransactions(enable bool) InitOption {
	return func(ak *AccountKeeper) {
		ak.enableUnorderedTxs = enable
	}
}

var _ AccountKeeperI = &AccountKeeper{}

// NewAccountKeeper returns a new AccountKeeperI that uses go-amino to
// (binary) encode and decode concrete sdk.Accounts.
// `maccPerms` is a map that takes accounts' addresses as keys, and their respective permissions as values. This map is used to construct
// types.PermissionsForAddress and is used in keeper.ValidatePermissions. Permissions are plain strings,
// and don't have to fit into any predefined structure. This auth module does not use account permissions internally, though other modules
// may use auth.Keeper to access the accounts permissions map.
func NewAccountKeeper(
	cdc codec.BinaryCodec, storeService store.KVStoreService, proto func() sdk.AccountI,
	maccPerms map[string][]string, ac address.Codec, bech32Prefix, authority string, opts ...InitOption,
) AccountKeeper {
	permAddrs := make(map[string]types.PermissionsForAddress)
	for name, perms := range maccPerms {
		permAddrs[name] = types.NewPermissionsForAddress(name, perms)
	}

	sb := collections.NewSchemaBuilder(storeService)

	ak := AccountKeeper{
		addressCodec:    ac,
		bech32Prefix:    bech32Prefix,
		storeService:    storeService,
		proto:           proto,
		cdc:             cdc,
		permAddrs:       permAddrs,
		authority:       authority,
		Params:          collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		AccountNumber:   collections.NewSequence(sb, types.GlobalAccountNumberKey, "account_number"),
		Accounts:        collections.NewIndexedMap(sb, types.AddressStoreKeyPrefix, "accounts", sdk.AccAddressKey, codec.CollInterfaceValue[sdk.AccountI](cdc), NewAccountIndexes(sb)),
		UnorderedNonces: collections.NewKeySet(sb, types.UnorderedNoncesKey, "unordered_nonces", collections.PairKeyCodec(collections.Int64Key, collections.BytesKey)),
	}
	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	ak.Schema = schema

	for _, opt := range opts {
		opt(&ak)
	}
	return ak
}

func (ak AccountKeeper) UnorderedTransactionsEnabled() bool {
	return ak.enableUnorderedTxs
}

// GetAuthority returns the x/auth module's authority.
func (ak AccountKeeper) GetAuthority() string {
	return ak.authority
}

// AddressCodec returns the x/auth account address codec.
// x/auth is tied to bech32 encoded user accounts
func (ak AccountKeeper) AddressCodec() address.Codec {
	return ak.addressCodec
}

// Logger returns a module-specific logger.
func (ak AccountKeeper) Logger(ctx context.Context) log.Logger {
	return sdk.UnwrapSDKContext(ctx).Logger().With("module", "x/"+types.ModuleName)
}

// GetPubKey Returns the PubKey of the account at address
func (ak AccountKeeper) GetPubKey(ctx context.Context, addr sdk.AccAddress) (cryptotypes.PubKey, error) {
	acc := ak.GetAccount(ctx, addr)
	if acc == nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownAddress, "account %s does not exist", addr)
	}
	return acc.GetPubKey(), nil
}

// GetSequence Returns the Sequence of the account at address
func (ak AccountKeeper) GetSequence(ctx context.Context, addr sdk.AccAddress) (uint64, error) {
	acc := ak.GetAccount(ctx, addr)
	if acc == nil {
		return 0, errorsmod.Wrapf(sdkerrors.ErrUnknownAddress, "account %s does not exist", addr)
	}

	return acc.GetSequence(), nil
}

func (ak AccountKeeper) getAccountNumberLegacy(ctx context.Context) (uint64, error) {
	store := ak.storeService.OpenKVStore(ctx)
	b, err := store.Get(types.LegacyGlobalAccountNumberKey)
	if err != nil {
		return 0, fmt.Errorf("failed to get legacy account number: %w", err)
	}
	v := new(gogotypes.UInt64Value)
	if err := v.Unmarshal(b); err != nil {
		return 0, fmt.Errorf("failed to unmarshal legacy account number: %w", err)
	}
	return v.Value, nil
}

// NextAccountNumber returns and increments the global account number counter.
// If the global account number is not set, it initializes it with value 0.
func (ak AccountKeeper) NextAccountNumber(ctx context.Context) uint64 {
	n, err := collections.Item[uint64](ak.AccountNumber).Get(ctx)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		// this won't happen in the tip of production network,
		// but can happen when query historical states,
		// fallback to old key for backward-compatibility.
		// for more info, see https://github.com/cosmos/cosmos-sdk/issues/23741
		n, err = ak.getAccountNumberLegacy(ctx)
	}

	if err != nil {
		panic(err)
	}

	if err := ak.AccountNumber.Set(ctx, n+1); err != nil {
		panic(err)
	}

	return n
}

// GetModulePermissions fetches per-module account permissions.
func (ak AccountKeeper) GetModulePermissions() map[string]types.PermissionsForAddress {
	return ak.permAddrs
}

// ValidatePermissions validates that the module account has been granted
// permissions within its set of allowed permissions.
func (ak AccountKeeper) ValidatePermissions(macc sdk.ModuleAccountI) error {
	permAddr := ak.permAddrs[macc.GetName()]
	for _, perm := range macc.GetPermissions() {
		if !permAddr.HasPermission(perm) {
			return fmt.Errorf("invalid module permission %s", perm)
		}
	}

	return nil
}

// GetModuleAddress returns an address based on the module name
func (ak AccountKeeper) GetModuleAddress(moduleName string) sdk.AccAddress {
	permAddr, ok := ak.permAddrs[moduleName]
	if !ok {
		return nil
	}

	return permAddr.GetAddress()
}

// GetModuleAddressAndPermissions returns an address and permissions based on the module name
func (ak AccountKeeper) GetModuleAddressAndPermissions(moduleName string) (addr sdk.AccAddress, permissions []string) {
	permAddr, ok := ak.permAddrs[moduleName]
	if !ok {
		return addr, permissions
	}

	return permAddr.GetAddress(), permAddr.GetPermissions()
}

// GetModuleAccountAndPermissions gets the module account from the auth account store and its
// registered permissions
func (ak AccountKeeper) GetModuleAccountAndPermissions(ctx context.Context, moduleName string) (sdk.ModuleAccountI, []string) {
	addr, perms := ak.GetModuleAddressAndPermissions(moduleName)
	if addr == nil {
		return nil, []string{}
	}

	acc := ak.GetAccount(ctx, addr)
	if acc != nil {
		macc, ok := acc.(sdk.ModuleAccountI)
		if !ok {
			panic("account is not a module account")
		}
		return macc, perms
	}

	// create a new module account
	macc := types.NewEmptyModuleAccount(moduleName, perms...)
	maccI := (ak.NewAccount(ctx, macc)).(sdk.ModuleAccountI) // set the account number
	ak.SetModuleAccount(ctx, maccI)

	return maccI, perms
}

// GetModuleAccount gets the module account from the auth account store, if the account does not
// exist in the AccountKeeper, then it is created.
func (ak AccountKeeper) GetModuleAccount(ctx context.Context, moduleName string) sdk.ModuleAccountI {
	acc, _ := ak.GetModuleAccountAndPermissions(ctx, moduleName)
	return acc
}

// SetModuleAccount sets the module account to the auth account store
func (ak AccountKeeper) SetModuleAccount(ctx context.Context, macc sdk.ModuleAccountI) {
	ak.SetAccount(ctx, macc)
}

// add getter for bech32Prefix
func (ak AccountKeeper) getBech32Prefix() (string, error) {
	return ak.bech32Prefix, nil
}

// GetParams gets the auth module's parameters.
func (ak AccountKeeper) GetParams(ctx context.Context) (params types.Params) {
	params, err := ak.Params.Get(ctx)
	if err != nil && !errors.Is(err, collections.ErrNotFound) {
		panic(err)
	}
	return params
}

// -------------------------------------
// Unordered Nonce management methods
// -------------------------------------

// ContainsUnorderedNonce reports whether the sender has used this timeout already.
func (ak AccountKeeper) ContainsUnorderedNonce(ctx sdk.Context, sender []byte, timeout time.Time) (bool, error) {
	return ak.UnorderedNonces.Has(ctx, collections.Join(timeout.UnixNano(), sender))
}

// TryAddUnorderedNonce tries to add a new unordered nonce for the sender.
// If the sender already has an entry with the provided timeout, an error is returned.
func (ak AccountKeeper) TryAddUnorderedNonce(ctx sdk.Context, sender []byte, timeout time.Time) error {
	alreadyHas, err := ak.ContainsUnorderedNonce(ctx, sender, timeout)
	if err != nil {
		return fmt.Errorf("failed to check unordered nonce in storage: %w", err)
	}
	if alreadyHas {
		return fmt.Errorf("sender %s has already used timeout %d", sdk.AccAddress(sender).String(), timeout.UnixNano())
	}

	return ak.UnorderedNonces.Set(ctx, collections.Join(timeout.UnixNano(), sender))
}

// RemoveExpiredUnorderedNonces removes all unordered nonces that have a timeout value before
// the current block time.
func (ak AccountKeeper) RemoveExpiredUnorderedNonces(ctx sdk.Context) error {
	blkTime := ctx.BlockTime().UnixNano()
	it, err := ak.UnorderedNonces.Iterate(ctx, collections.NewPrefixUntilPairRange[int64, []byte](blkTime))
	if err != nil {
		return err
	}
	defer it.Close()

	keys, err := it.Keys()
	if err != nil {
		return err
	}

	for _, key := range keys {
		if err := ak.UnorderedNonces.Remove(ctx, key); err != nil {
			return err
		}
	}

	return nil
}
