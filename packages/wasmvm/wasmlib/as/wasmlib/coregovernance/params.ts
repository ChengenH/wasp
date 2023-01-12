// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

import * as wasmtypes from "../wasmtypes";
import * as sc from "./index";

export class ImmutableAddAllowedStateControllerAddressParams extends wasmtypes.ScProxy {
    stateControllerAddress(): wasmtypes.ScImmutableAddress {
        return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamStateControllerAddress));
    }
}

export class MutableAddAllowedStateControllerAddressParams extends wasmtypes.ScProxy {
    stateControllerAddress(): wasmtypes.ScMutableAddress {
        return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamStateControllerAddress));
    }
}

export class ImmutableAddCandidateNodeParams extends wasmtypes.ScProxy {
    accessNodeInfoAccessAPI(): wasmtypes.ScImmutableString {
        return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamAccessNodeInfoAccessAPI));
    }

    accessNodeInfoCertificate(): wasmtypes.ScImmutableBytes {
        return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ParamAccessNodeInfoCertificate));
    }

    accessNodeInfoForCommittee(): wasmtypes.ScImmutableBool {
        return new wasmtypes.ScImmutableBool(this.proxy.root(sc.ParamAccessNodeInfoForCommittee));
    }

    accessNodeInfoPubKey(): wasmtypes.ScImmutableBytes {
        return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ParamAccessNodeInfoPubKey));
    }
}

export class MutableAddCandidateNodeParams extends wasmtypes.ScProxy {
    accessNodeInfoAccessAPI(): wasmtypes.ScMutableString {
        return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamAccessNodeInfoAccessAPI));
    }

    accessNodeInfoCertificate(): wasmtypes.ScMutableBytes {
        return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ParamAccessNodeInfoCertificate));
    }

    accessNodeInfoForCommittee(): wasmtypes.ScMutableBool {
        return new wasmtypes.ScMutableBool(this.proxy.root(sc.ParamAccessNodeInfoForCommittee));
    }

    accessNodeInfoPubKey(): wasmtypes.ScMutableBytes {
        return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ParamAccessNodeInfoPubKey));
    }
}

export class MapBytesToImmutableUint8 extends wasmtypes.ScProxy {

    getUint8(key: Uint8Array): wasmtypes.ScImmutableUint8 {
        return new wasmtypes.ScImmutableUint8(this.proxy.key(wasmtypes.bytesToBytes(key)));
    }
}

export class ImmutableChangeAccessNodesParams extends wasmtypes.ScProxy {
    changeAccessNodesActions(): sc.MapBytesToImmutableUint8 {
        return new sc.MapBytesToImmutableUint8(this.proxy.root(sc.ParamChangeAccessNodesActions));
    }
}

export class MapBytesToMutableUint8 extends wasmtypes.ScProxy {

    clear(): void {
        this.proxy.clearMap();
    }

    getUint8(key: Uint8Array): wasmtypes.ScMutableUint8 {
        return new wasmtypes.ScMutableUint8(this.proxy.key(wasmtypes.bytesToBytes(key)));
    }
}

export class MutableChangeAccessNodesParams extends wasmtypes.ScProxy {
    changeAccessNodesActions(): sc.MapBytesToMutableUint8 {
        return new sc.MapBytesToMutableUint8(this.proxy.root(sc.ParamChangeAccessNodesActions));
    }
}

export class ImmutableDelegateChainOwnershipParams extends wasmtypes.ScProxy {
    chainOwner(): wasmtypes.ScImmutableAgentID {
        return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamChainOwner));
    }
}

export class MutableDelegateChainOwnershipParams extends wasmtypes.ScProxy {
    chainOwner(): wasmtypes.ScMutableAgentID {
        return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamChainOwner));
    }
}

export class ImmutableRemoveAllowedStateControllerAddressParams extends wasmtypes.ScProxy {
    stateControllerAddress(): wasmtypes.ScImmutableAddress {
        return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamStateControllerAddress));
    }
}

export class MutableRemoveAllowedStateControllerAddressParams extends wasmtypes.ScProxy {
    stateControllerAddress(): wasmtypes.ScMutableAddress {
        return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamStateControllerAddress));
    }
}

export class ImmutableRevokeAccessNodeParams extends wasmtypes.ScProxy {
    accessNodeInfoCertificate(): wasmtypes.ScImmutableBytes {
        return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ParamAccessNodeInfoCertificate));
    }

    accessNodeInfoPubKey(): wasmtypes.ScImmutableBytes {
        return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ParamAccessNodeInfoPubKey));
    }
}

export class MutableRevokeAccessNodeParams extends wasmtypes.ScProxy {
    accessNodeInfoCertificate(): wasmtypes.ScMutableBytes {
        return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ParamAccessNodeInfoCertificate));
    }

    accessNodeInfoPubKey(): wasmtypes.ScMutableBytes {
        return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ParamAccessNodeInfoPubKey));
    }
}

export class ImmutableRotateStateControllerParams extends wasmtypes.ScProxy {
    stateControllerAddress(): wasmtypes.ScImmutableAddress {
        return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamStateControllerAddress));
    }
}

export class MutableRotateStateControllerParams extends wasmtypes.ScProxy {
    stateControllerAddress(): wasmtypes.ScMutableAddress {
        return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamStateControllerAddress));
    }
}

export class ImmutableSetChainInfoParams extends wasmtypes.ScProxy {
    // default maximum size of a blob
    maxBlobSize(): wasmtypes.ScImmutableUint32 {
        return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamMaxBlobSize));
    }

    // default maximum size of a single event
    maxEventSize(): wasmtypes.ScImmutableUint16 {
        return new wasmtypes.ScImmutableUint16(this.proxy.root(sc.ParamMaxEventSize));
    }

    // default maximum number of events per request
    maxEventsPerReq(): wasmtypes.ScImmutableUint16 {
        return new wasmtypes.ScImmutableUint16(this.proxy.root(sc.ParamMaxEventsPerReq));
    }
}

export class MutableSetChainInfoParams extends wasmtypes.ScProxy {
    // default maximum size of a blob
    maxBlobSize(): wasmtypes.ScMutableUint32 {
        return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamMaxBlobSize));
    }

    // default maximum size of a single event
    maxEventSize(): wasmtypes.ScMutableUint16 {
        return new wasmtypes.ScMutableUint16(this.proxy.root(sc.ParamMaxEventSize));
    }

    // default maximum number of events per request
    maxEventsPerReq(): wasmtypes.ScMutableUint16 {
        return new wasmtypes.ScMutableUint16(this.proxy.root(sc.ParamMaxEventsPerReq));
    }
}

export class ImmutableSetFeePolicyParams extends wasmtypes.ScProxy {
    feePolicyBytes(): wasmtypes.ScImmutableBytes {
        return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ParamFeePolicyBytes));
    }
}

export class MutableSetFeePolicyParams extends wasmtypes.ScProxy {
    feePolicyBytes(): wasmtypes.ScMutableBytes {
        return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ParamFeePolicyBytes));
    }
}