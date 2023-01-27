// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

#![allow(dead_code)]

use crate::*;
use crate::coreaccounts::*;

pub struct DepositCall<'a> {
    pub func: ScFunc<'a>,
}

pub struct FoundryCreateNewCall<'a> {
    pub func:    ScFunc<'a>,
    pub params:  MutableFoundryCreateNewParams,
    pub results: ImmutableFoundryCreateNewResults,
}

pub struct FoundryDestroyCall<'a> {
    pub func:   ScFunc<'a>,
    pub params: MutableFoundryDestroyParams,
}

pub struct FoundryModifySupplyCall<'a> {
    pub func:   ScFunc<'a>,
    pub params: MutableFoundryModifySupplyParams,
}

pub struct HarvestCall<'a> {
    pub func:   ScFunc<'a>,
    pub params: MutableHarvestParams,
}

pub struct TransferAllowanceToCall<'a> {
    pub func:   ScFunc<'a>,
    pub params: MutableTransferAllowanceToParams,
}

pub struct WithdrawCall<'a> {
    pub func: ScFunc<'a>,
}

pub struct AccountNFTsCall<'a> {
    pub func:    ScView<'a>,
    pub params:  MutableAccountNFTsParams,
    pub results: ImmutableAccountNFTsResults,
}

pub struct AccountsCall<'a> {
    pub func:    ScView<'a>,
    pub results: ImmutableAccountsResults,
}

pub struct BalanceCall<'a> {
    pub func:    ScView<'a>,
    pub params:  MutableBalanceParams,
    pub results: ImmutableBalanceResults,
}

pub struct BalanceBaseTokenCall<'a> {
    pub func:    ScView<'a>,
    pub params:  MutableBalanceBaseTokenParams,
    pub results: ImmutableBalanceBaseTokenResults,
}

pub struct BalanceNativeTokenCall<'a> {
    pub func:    ScView<'a>,
    pub params:  MutableBalanceNativeTokenParams,
    pub results: ImmutableBalanceNativeTokenResults,
}

pub struct FoundryOutputCall<'a> {
    pub func:    ScView<'a>,
    pub params:  MutableFoundryOutputParams,
    pub results: ImmutableFoundryOutputResults,
}

pub struct GetAccountNonceCall<'a> {
    pub func:    ScView<'a>,
    pub params:  MutableGetAccountNonceParams,
    pub results: ImmutableGetAccountNonceResults,
}

pub struct GetNativeTokenIDRegistryCall<'a> {
    pub func:    ScView<'a>,
    pub results: ImmutableGetNativeTokenIDRegistryResults,
}

pub struct NftDataCall<'a> {
    pub func:    ScView<'a>,
    pub params:  MutableNftDataParams,
    pub results: ImmutableNftDataResults,
}

pub struct TotalAssetsCall<'a> {
    pub func:    ScView<'a>,
    pub results: ImmutableTotalAssetsResults,
}

pub struct ScFuncs {
}

impl ScFuncs {
    pub fn deposit(ctx: &impl ScFuncCallContext) -> DepositCall {
        DepositCall {
            func: ScFunc::new(ctx, HSC_NAME, HFUNC_DEPOSIT),
        }
    }

    pub fn foundry_create_new(ctx: &impl ScFuncCallContext) -> FoundryCreateNewCall {
        let mut f = FoundryCreateNewCall {
            func:    ScFunc::new(ctx, HSC_NAME, HFUNC_FOUNDRY_CREATE_NEW),
            params:  MutableFoundryCreateNewParams { proxy: Proxy::nil() },
            results: ImmutableFoundryCreateNewResults { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        ScFunc::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn foundry_destroy(ctx: &impl ScFuncCallContext) -> FoundryDestroyCall {
        let mut f = FoundryDestroyCall {
            func:    ScFunc::new(ctx, HSC_NAME, HFUNC_FOUNDRY_DESTROY),
            params:  MutableFoundryDestroyParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn foundry_modify_supply(ctx: &impl ScFuncCallContext) -> FoundryModifySupplyCall {
        let mut f = FoundryModifySupplyCall {
            func:    ScFunc::new(ctx, HSC_NAME, HFUNC_FOUNDRY_MODIFY_SUPPLY),
            params:  MutableFoundryModifySupplyParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn harvest(ctx: &impl ScFuncCallContext) -> HarvestCall {
        let mut f = HarvestCall {
            func:    ScFunc::new(ctx, HSC_NAME, HFUNC_HARVEST),
            params:  MutableHarvestParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn transfer_allowance_to(ctx: &impl ScFuncCallContext) -> TransferAllowanceToCall {
        let mut f = TransferAllowanceToCall {
            func:    ScFunc::new(ctx, HSC_NAME, HFUNC_TRANSFER_ALLOWANCE_TO),
            params:  MutableTransferAllowanceToParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn withdraw(ctx: &impl ScFuncCallContext) -> WithdrawCall {
        WithdrawCall {
            func: ScFunc::new(ctx, HSC_NAME, HFUNC_WITHDRAW),
        }
    }

    pub fn account_nf_ts(ctx: &impl ScViewCallContext) -> AccountNFTsCall {
        let mut f = AccountNFTsCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_ACCOUNT_NF_TS),
            params:  MutableAccountNFTsParams { proxy: Proxy::nil() },
            results: ImmutableAccountNFTsResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn accounts(ctx: &impl ScViewCallContext) -> AccountsCall {
        let mut f = AccountsCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_ACCOUNTS),
            results: ImmutableAccountsResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn balance(ctx: &impl ScViewCallContext) -> BalanceCall {
        let mut f = BalanceCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_BALANCE),
            params:  MutableBalanceParams { proxy: Proxy::nil() },
            results: ImmutableBalanceResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn balance_base_token(ctx: &impl ScViewCallContext) -> BalanceBaseTokenCall {
        let mut f = BalanceBaseTokenCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_BALANCE_BASE_TOKEN),
            params:  MutableBalanceBaseTokenParams { proxy: Proxy::nil() },
            results: ImmutableBalanceBaseTokenResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn balance_native_token(ctx: &impl ScViewCallContext) -> BalanceNativeTokenCall {
        let mut f = BalanceNativeTokenCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_BALANCE_NATIVE_TOKEN),
            params:  MutableBalanceNativeTokenParams { proxy: Proxy::nil() },
            results: ImmutableBalanceNativeTokenResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn foundry_output(ctx: &impl ScViewCallContext) -> FoundryOutputCall {
        let mut f = FoundryOutputCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_FOUNDRY_OUTPUT),
            params:  MutableFoundryOutputParams { proxy: Proxy::nil() },
            results: ImmutableFoundryOutputResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn get_account_nonce(ctx: &impl ScViewCallContext) -> GetAccountNonceCall {
        let mut f = GetAccountNonceCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_GET_ACCOUNT_NONCE),
            params:  MutableGetAccountNonceParams { proxy: Proxy::nil() },
            results: ImmutableGetAccountNonceResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn get_native_token_id_registry(ctx: &impl ScViewCallContext) -> GetNativeTokenIDRegistryCall {
        let mut f = GetNativeTokenIDRegistryCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_GET_NATIVE_TOKEN_ID_REGISTRY),
            results: ImmutableGetNativeTokenIDRegistryResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn nft_data(ctx: &impl ScViewCallContext) -> NftDataCall {
        let mut f = NftDataCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_NFT_DATA),
            params:  MutableNftDataParams { proxy: Proxy::nil() },
            results: ImmutableNftDataResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn total_assets(ctx: &impl ScViewCallContext) -> TotalAssetsCall {
        let mut f = TotalAssetsCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_TOTAL_ASSETS),
            results: ImmutableTotalAssetsResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }
}
