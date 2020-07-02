package keeper

import (
	"fmt"

	//"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/PaddyMc/poa/x/poa/types"
)

// Implements ValidatorSet interface
var _ types.ValidatorSet = Keeper{}

// Keeper of the poa store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	paramspace types.ParamSubspace
}

// NewKeeper creates a poa keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramspace types.ParamSubspace) Keeper {
	keeper := Keeper{
		storeKey:   key,
		cdc:        cdc,
		paramspace: paramspace.WithKeyTable(ParamKeyTable()),
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetValidator(ctx sdk.Context, addr sdk.ValAddress) (found bool) {
	//store := ctx.KVStore(k.storeKey)
	fmt.Sprintf("x/%s", types.ModuleName)
	//value := store.Get(types.GetValidatorKey(addr))
	//if value == nil {
	//	return validator, false
	//}

	//// If these amino encoded bytes are in the cache, return the cached validator
	//strValue := string(value)
	//if val, ok := k.validatorCache[strValue]; ok {
	//	valToReturn := val.val
	//	// Doesn't mutate the cache's value
	//	valToReturn.OperatorAddress = addr
	//	return valToReturn, true
	//}

	//// amino bytes weren't found in cache, so amino unmarshal and add it to the cache
	//validator = types.MustUnmarshalValidator(k.cdc, value)
	//cachedVal := newCachedValidator(validator, strValue)
	//k.validatorCache[strValue] = newCachedValidator(validator, strValue)
	//k.validatorCacheList.PushBack(cachedVal)

	//// if the cache is too big, pop off the last element from it
	//if k.validatorCacheList.Len() > aminoCacheSize {
	//	valToRemove := k.validatorCacheList.Remove(k.validatorCacheList.Front()).(cachedValidator)
	//	delete(k.validatorCache, valToRemove.marshalled)
	//}

	//validator = types.MustUnmarshalValidator(k.cdc, value)
	return true
}

