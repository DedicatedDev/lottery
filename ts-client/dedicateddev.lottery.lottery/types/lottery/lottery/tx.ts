/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "dedicateddev.lottery.lottery";

export interface MsgCreateLottery {
  creator: string;
  mintBetAmount: number;
  fee: number;
}

export interface MsgCreateLotteryResponse {
  lotteryIndex: string;
}

export interface MsgBidToLottery {
  creator: string;
  lotteryId: number;
  bidAmount: number;
  bidCount: number;
}

export interface MsgBidToLotteryResponse {
  bidId: number;
}

function createBaseMsgCreateLottery(): MsgCreateLottery {
  return { creator: "", mintBetAmount: 0, fee: 0 };
}

export const MsgCreateLottery = {
  encode(message: MsgCreateLottery, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.mintBetAmount !== 0) {
      writer.uint32(16).uint64(message.mintBetAmount);
    }
    if (message.fee !== 0) {
      writer.uint32(24).uint64(message.fee);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateLottery {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateLottery();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.mintBetAmount = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.fee = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateLottery {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      mintBetAmount: isSet(object.mintBetAmount) ? Number(object.mintBetAmount) : 0,
      fee: isSet(object.fee) ? Number(object.fee) : 0,
    };
  },

  toJSON(message: MsgCreateLottery): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.mintBetAmount !== undefined && (obj.mintBetAmount = Math.round(message.mintBetAmount));
    message.fee !== undefined && (obj.fee = Math.round(message.fee));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateLottery>, I>>(object: I): MsgCreateLottery {
    const message = createBaseMsgCreateLottery();
    message.creator = object.creator ?? "";
    message.mintBetAmount = object.mintBetAmount ?? 0;
    message.fee = object.fee ?? 0;
    return message;
  },
};

function createBaseMsgCreateLotteryResponse(): MsgCreateLotteryResponse {
  return { lotteryIndex: "" };
}

export const MsgCreateLotteryResponse = {
  encode(message: MsgCreateLotteryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.lotteryIndex !== "") {
      writer.uint32(10).string(message.lotteryIndex);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateLotteryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateLotteryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.lotteryIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateLotteryResponse {
    return { lotteryIndex: isSet(object.lotteryIndex) ? String(object.lotteryIndex) : "" };
  },

  toJSON(message: MsgCreateLotteryResponse): unknown {
    const obj: any = {};
    message.lotteryIndex !== undefined && (obj.lotteryIndex = message.lotteryIndex);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateLotteryResponse>, I>>(object: I): MsgCreateLotteryResponse {
    const message = createBaseMsgCreateLotteryResponse();
    message.lotteryIndex = object.lotteryIndex ?? "";
    return message;
  },
};

function createBaseMsgBidToLottery(): MsgBidToLottery {
  return { creator: "", lotteryId: 0, bidAmount: 0, bidCount: 0 };
}

export const MsgBidToLottery = {
  encode(message: MsgBidToLottery, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.lotteryId !== 0) {
      writer.uint32(16).uint64(message.lotteryId);
    }
    if (message.bidAmount !== 0) {
      writer.uint32(24).uint64(message.bidAmount);
    }
    if (message.bidCount !== 0) {
      writer.uint32(32).uint64(message.bidCount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBidToLottery {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBidToLottery();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.lotteryId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.bidAmount = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.bidCount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBidToLottery {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      lotteryId: isSet(object.lotteryId) ? Number(object.lotteryId) : 0,
      bidAmount: isSet(object.bidAmount) ? Number(object.bidAmount) : 0,
      bidCount: isSet(object.bidCount) ? Number(object.bidCount) : 0,
    };
  },

  toJSON(message: MsgBidToLottery): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.lotteryId !== undefined && (obj.lotteryId = Math.round(message.lotteryId));
    message.bidAmount !== undefined && (obj.bidAmount = Math.round(message.bidAmount));
    message.bidCount !== undefined && (obj.bidCount = Math.round(message.bidCount));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBidToLottery>, I>>(object: I): MsgBidToLottery {
    const message = createBaseMsgBidToLottery();
    message.creator = object.creator ?? "";
    message.lotteryId = object.lotteryId ?? 0;
    message.bidAmount = object.bidAmount ?? 0;
    message.bidCount = object.bidCount ?? 0;
    return message;
  },
};

function createBaseMsgBidToLotteryResponse(): MsgBidToLotteryResponse {
  return { bidId: 0 };
}

export const MsgBidToLotteryResponse = {
  encode(message: MsgBidToLotteryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.bidId !== 0) {
      writer.uint32(8).uint64(message.bidId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBidToLotteryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBidToLotteryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.bidId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBidToLotteryResponse {
    return { bidId: isSet(object.bidId) ? Number(object.bidId) : 0 };
  },

  toJSON(message: MsgBidToLotteryResponse): unknown {
    const obj: any = {};
    message.bidId !== undefined && (obj.bidId = Math.round(message.bidId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBidToLotteryResponse>, I>>(object: I): MsgBidToLotteryResponse {
    const message = createBaseMsgBidToLotteryResponse();
    message.bidId = object.bidId ?? 0;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateLottery(request: MsgCreateLottery): Promise<MsgCreateLotteryResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  BidToLottery(request: MsgBidToLottery): Promise<MsgBidToLotteryResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateLottery = this.CreateLottery.bind(this);
    this.BidToLottery = this.BidToLottery.bind(this);
  }
  CreateLottery(request: MsgCreateLottery): Promise<MsgCreateLotteryResponse> {
    const data = MsgCreateLottery.encode(request).finish();
    const promise = this.rpc.request("dedicateddev.lottery.lottery.Msg", "CreateLottery", data);
    return promise.then((data) => MsgCreateLotteryResponse.decode(new _m0.Reader(data)));
  }

  BidToLottery(request: MsgBidToLottery): Promise<MsgBidToLotteryResponse> {
    const data = MsgBidToLottery.encode(request).finish();
    const promise = this.rpc.request("dedicateddev.lottery.lottery.Msg", "BidToLottery", data);
    return promise.then((data) => MsgBidToLotteryResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
