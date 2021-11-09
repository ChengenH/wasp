// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]

use std::ptr;

use crate::*;
use crate::coreblocklog::*;

pub struct ControlAddressesCall {
	pub func: ScView,
	pub results: ImmutableControlAddressesResults,
}

pub struct GetBlockInfoCall {
	pub func: ScView,
	pub params: MutableGetBlockInfoParams,
	pub results: ImmutableGetBlockInfoResults,
}

pub struct GetEventsForBlockCall {
	pub func: ScView,
	pub params: MutableGetEventsForBlockParams,
	pub results: ImmutableGetEventsForBlockResults,
}

pub struct GetEventsForContractCall {
	pub func: ScView,
	pub params: MutableGetEventsForContractParams,
	pub results: ImmutableGetEventsForContractResults,
}

pub struct GetEventsForRequestCall {
	pub func: ScView,
	pub params: MutableGetEventsForRequestParams,
	pub results: ImmutableGetEventsForRequestResults,
}

pub struct GetLatestBlockInfoCall {
	pub func: ScView,
	pub results: ImmutableGetLatestBlockInfoResults,
}

pub struct GetRequestIDsForBlockCall {
	pub func: ScView,
	pub params: MutableGetRequestIDsForBlockParams,
	pub results: ImmutableGetRequestIDsForBlockResults,
}

pub struct GetRequestReceiptCall {
	pub func: ScView,
	pub params: MutableGetRequestReceiptParams,
	pub results: ImmutableGetRequestReceiptResults,
}

pub struct GetRequestReceiptsForBlockCall {
	pub func: ScView,
	pub params: MutableGetRequestReceiptsForBlockParams,
	pub results: ImmutableGetRequestReceiptsForBlockResults,
}

pub struct IsRequestProcessedCall {
	pub func: ScView,
	pub params: MutableIsRequestProcessedParams,
	pub results: ImmutableIsRequestProcessedResults,
}

pub struct ScFuncs {
}

impl ScFuncs {
    pub fn control_addresses(_ctx: & dyn ScViewCallContext) -> ControlAddressesCall {
        let mut f = ControlAddressesCall {
            func: ScView::new(HSC_NAME, HVIEW_CONTROL_ADDRESSES),
            results: ImmutableControlAddressesResults { id: 0 },
        };
        f.func.set_ptrs(ptr::null_mut(), &mut f.results.id);
        f
    }
    pub fn get_block_info(_ctx: & dyn ScViewCallContext) -> GetBlockInfoCall {
        let mut f = GetBlockInfoCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_BLOCK_INFO),
            params: MutableGetBlockInfoParams { id: 0 },
            results: ImmutableGetBlockInfoResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }
    pub fn get_events_for_block(_ctx: & dyn ScViewCallContext) -> GetEventsForBlockCall {
        let mut f = GetEventsForBlockCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_EVENTS_FOR_BLOCK),
            params: MutableGetEventsForBlockParams { id: 0 },
            results: ImmutableGetEventsForBlockResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }
    pub fn get_events_for_contract(_ctx: & dyn ScViewCallContext) -> GetEventsForContractCall {
        let mut f = GetEventsForContractCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_EVENTS_FOR_CONTRACT),
            params: MutableGetEventsForContractParams { id: 0 },
            results: ImmutableGetEventsForContractResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }
    pub fn get_events_for_request(_ctx: & dyn ScViewCallContext) -> GetEventsForRequestCall {
        let mut f = GetEventsForRequestCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_EVENTS_FOR_REQUEST),
            params: MutableGetEventsForRequestParams { id: 0 },
            results: ImmutableGetEventsForRequestResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }
    pub fn get_latest_block_info(_ctx: & dyn ScViewCallContext) -> GetLatestBlockInfoCall {
        let mut f = GetLatestBlockInfoCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_LATEST_BLOCK_INFO),
            results: ImmutableGetLatestBlockInfoResults { id: 0 },
        };
        f.func.set_ptrs(ptr::null_mut(), &mut f.results.id);
        f
    }
    pub fn get_request_i_ds_for_block(_ctx: & dyn ScViewCallContext) -> GetRequestIDsForBlockCall {
        let mut f = GetRequestIDsForBlockCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_REQUEST_I_DS_FOR_BLOCK),
            params: MutableGetRequestIDsForBlockParams { id: 0 },
            results: ImmutableGetRequestIDsForBlockResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }
    pub fn get_request_receipt(_ctx: & dyn ScViewCallContext) -> GetRequestReceiptCall {
        let mut f = GetRequestReceiptCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_REQUEST_RECEIPT),
            params: MutableGetRequestReceiptParams { id: 0 },
            results: ImmutableGetRequestReceiptResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }
    pub fn get_request_receipts_for_block(_ctx: & dyn ScViewCallContext) -> GetRequestReceiptsForBlockCall {
        let mut f = GetRequestReceiptsForBlockCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_REQUEST_RECEIPTS_FOR_BLOCK),
            params: MutableGetRequestReceiptsForBlockParams { id: 0 },
            results: ImmutableGetRequestReceiptsForBlockResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }
    pub fn is_request_processed(_ctx: & dyn ScViewCallContext) -> IsRequestProcessedCall {
        let mut f = IsRequestProcessedCall {
            func: ScView::new(HSC_NAME, HVIEW_IS_REQUEST_PROCESSED),
            params: MutableIsRequestProcessedParams { id: 0 },
            results: ImmutableIsRequestProcessedResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }
}