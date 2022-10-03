// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

#![allow(dead_code)]
#![allow(unused_imports)]

use crate::coreblocklog::*;
use crate::*;

#[derive(Clone)]
pub struct ImmutableControlAddressesResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableControlAddressesResults {
    // the addresses have been set as state controller address or governing address since the following block index
    pub fn block_index(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(RESULT_BLOCK_INDEX))
	}

    pub fn governing_address(&self) -> ScImmutableAddress {
		ScImmutableAddress::new(self.proxy.root(RESULT_GOVERNING_ADDRESS))
	}

    pub fn state_controller_address(&self) -> ScImmutableAddress {
		ScImmutableAddress::new(self.proxy.root(RESULT_STATE_CONTROLLER_ADDRESS))
	}
}

#[derive(Clone)]
pub struct MutableControlAddressesResults {
	pub(crate) proxy: Proxy,
}

impl MutableControlAddressesResults {
    // the addresses have been set as state controller address or governing address since the following block index
    pub fn block_index(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(RESULT_BLOCK_INDEX))
	}

    pub fn governing_address(&self) -> ScMutableAddress {
		ScMutableAddress::new(self.proxy.root(RESULT_GOVERNING_ADDRESS))
	}

    pub fn state_controller_address(&self) -> ScMutableAddress {
		ScMutableAddress::new(self.proxy.root(RESULT_STATE_CONTROLLER_ADDRESS))
	}
}

#[derive(Clone)]
pub struct ImmutableGetBlockInfoResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetBlockInfoResults {
    pub fn block_index(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(RESULT_BLOCK_INDEX))
	}

    pub fn block_info(&self) -> ScImmutableBytes {
		ScImmutableBytes::new(self.proxy.root(RESULT_BLOCK_INFO))
	}
}

#[derive(Clone)]
pub struct MutableGetBlockInfoResults {
	pub(crate) proxy: Proxy,
}

impl MutableGetBlockInfoResults {
    pub fn block_index(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(RESULT_BLOCK_INDEX))
	}

    pub fn block_info(&self) -> ScMutableBytes {
		ScMutableBytes::new(self.proxy.root(RESULT_BLOCK_INFO))
	}
}

#[derive(Clone)]
pub struct ArrayOfImmutableBytes {
	pub(crate) proxy: Proxy,
}

impl ArrayOfImmutableBytes {
    pub fn length(&self) -> u32 {
        self.proxy.length()
    }

    pub fn get_bytes(&self, index: u32) -> ScImmutableBytes {
        ScImmutableBytes::new(self.proxy.index(index))
    }
}

#[derive(Clone)]
pub struct ImmutableGetEventsForBlockResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetEventsForBlockResults {
    // native contract, so this is an Array16
    pub fn event(&self) -> ArrayOfImmutableBytes {
		ArrayOfImmutableBytes { proxy: self.proxy.root(RESULT_EVENT) }
	}
}

#[derive(Clone)]
pub struct ArrayOfMutableBytes {
	pub(crate) proxy: Proxy,
}

impl ArrayOfMutableBytes {
	pub fn append_bytes(&self) -> ScMutableBytes {
		ScMutableBytes::new(self.proxy.append())
	}

	pub fn clear(&self) {
        self.proxy.clear_array();
    }

    pub fn length(&self) -> u32 {
        self.proxy.length()
    }

    pub fn get_bytes(&self, index: u32) -> ScMutableBytes {
        ScMutableBytes::new(self.proxy.index(index))
    }
}

#[derive(Clone)]
pub struct MutableGetEventsForBlockResults {
	pub(crate) proxy: Proxy,
}

impl MutableGetEventsForBlockResults {
    // native contract, so this is an Array16
    pub fn event(&self) -> ArrayOfMutableBytes {
		ArrayOfMutableBytes { proxy: self.proxy.root(RESULT_EVENT) }
	}
}

#[derive(Clone)]
pub struct ImmutableGetEventsForContractResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetEventsForContractResults {
    // native contract, so this is an Array16
    pub fn event(&self) -> ArrayOfImmutableBytes {
		ArrayOfImmutableBytes { proxy: self.proxy.root(RESULT_EVENT) }
	}
}

#[derive(Clone)]
pub struct MutableGetEventsForContractResults {
	pub(crate) proxy: Proxy,
}

impl MutableGetEventsForContractResults {
    // native contract, so this is an Array16
    pub fn event(&self) -> ArrayOfMutableBytes {
		ArrayOfMutableBytes { proxy: self.proxy.root(RESULT_EVENT) }
	}
}

#[derive(Clone)]
pub struct ImmutableGetEventsForRequestResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetEventsForRequestResults {
    // native contract, so this is an Array16
    pub fn event(&self) -> ArrayOfImmutableBytes {
		ArrayOfImmutableBytes { proxy: self.proxy.root(RESULT_EVENT) }
	}
}

#[derive(Clone)]
pub struct MutableGetEventsForRequestResults {
	pub(crate) proxy: Proxy,
}

impl MutableGetEventsForRequestResults {
    // native contract, so this is an Array16
    pub fn event(&self) -> ArrayOfMutableBytes {
		ArrayOfMutableBytes { proxy: self.proxy.root(RESULT_EVENT) }
	}
}

#[derive(Clone)]
pub struct ArrayOfImmutableRequestID {
	pub(crate) proxy: Proxy,
}

impl ArrayOfImmutableRequestID {
    pub fn length(&self) -> u32 {
        self.proxy.length()
    }

    pub fn get_request_id(&self, index: u32) -> ScImmutableRequestID {
        ScImmutableRequestID::new(self.proxy.index(index))
    }
}

#[derive(Clone)]
pub struct ImmutableGetRequestIDsForBlockResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetRequestIDsForBlockResults {
    // native contract, so this is an Array16
    pub fn request_id(&self) -> ArrayOfImmutableRequestID {
		ArrayOfImmutableRequestID { proxy: self.proxy.root(RESULT_REQUEST_ID) }
	}
}

#[derive(Clone)]
pub struct ArrayOfMutableRequestID {
	pub(crate) proxy: Proxy,
}

impl ArrayOfMutableRequestID {
	pub fn append_request_id(&self) -> ScMutableRequestID {
		ScMutableRequestID::new(self.proxy.append())
	}

	pub fn clear(&self) {
        self.proxy.clear_array();
    }

    pub fn length(&self) -> u32 {
        self.proxy.length()
    }

    pub fn get_request_id(&self, index: u32) -> ScMutableRequestID {
        ScMutableRequestID::new(self.proxy.index(index))
    }
}

#[derive(Clone)]
pub struct MutableGetRequestIDsForBlockResults {
	pub(crate) proxy: Proxy,
}

impl MutableGetRequestIDsForBlockResults {
    // native contract, so this is an Array16
    pub fn request_id(&self) -> ArrayOfMutableRequestID {
		ArrayOfMutableRequestID { proxy: self.proxy.root(RESULT_REQUEST_ID) }
	}
}

#[derive(Clone)]
pub struct ImmutableGetRequestReceiptResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetRequestReceiptResults {
    pub fn block_index(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(RESULT_BLOCK_INDEX))
	}

    pub fn request_index(&self) -> ScImmutableUint16 {
		ScImmutableUint16::new(self.proxy.root(RESULT_REQUEST_INDEX))
	}

    pub fn request_record(&self) -> ScImmutableBytes {
		ScImmutableBytes::new(self.proxy.root(RESULT_REQUEST_RECORD))
	}
}

#[derive(Clone)]
pub struct MutableGetRequestReceiptResults {
	pub(crate) proxy: Proxy,
}

impl MutableGetRequestReceiptResults {
    pub fn block_index(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(RESULT_BLOCK_INDEX))
	}

    pub fn request_index(&self) -> ScMutableUint16 {
		ScMutableUint16::new(self.proxy.root(RESULT_REQUEST_INDEX))
	}

    pub fn request_record(&self) -> ScMutableBytes {
		ScMutableBytes::new(self.proxy.root(RESULT_REQUEST_RECORD))
	}
}

#[derive(Clone)]
pub struct ImmutableGetRequestReceiptsForBlockResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetRequestReceiptsForBlockResults {
    // native contract, so this is an Array16
    pub fn request_record(&self) -> ArrayOfImmutableBytes {
		ArrayOfImmutableBytes { proxy: self.proxy.root(RESULT_REQUEST_RECORD) }
	}
}

#[derive(Clone)]
pub struct MutableGetRequestReceiptsForBlockResults {
	pub(crate) proxy: Proxy,
}

impl MutableGetRequestReceiptsForBlockResults {
    // native contract, so this is an Array16
    pub fn request_record(&self) -> ArrayOfMutableBytes {
		ArrayOfMutableBytes { proxy: self.proxy.root(RESULT_REQUEST_RECORD) }
	}
}

#[derive(Clone)]
pub struct ImmutableIsRequestProcessedResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableIsRequestProcessedResults {
    pub fn request_processed(&self) -> ScImmutableBool {
		ScImmutableBool::new(self.proxy.root(RESULT_REQUEST_PROCESSED))
	}
}

#[derive(Clone)]
pub struct MutableIsRequestProcessedResults {
	pub(crate) proxy: Proxy,
}

impl MutableIsRequestProcessedResults {
    pub fn request_processed(&self) -> ScMutableBool {
		ScMutableBool::new(self.proxy.root(RESULT_REQUEST_PROCESSED))
	}
}
