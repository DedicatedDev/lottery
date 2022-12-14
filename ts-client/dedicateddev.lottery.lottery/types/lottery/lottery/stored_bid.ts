/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "dedicateddev.lottery.lottery";

export interface StoredBid {
  index: string;
  lotteryId: number;
  betAmount: number;
  bidder: string;
}

function createBaseStoredBid(): StoredBid {
  return { index: "", lotteryId: 0, betAmount: 0, bidder: "" };
}

export const StoredBid = {
  encode(message: StoredBid, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.lotteryId !== 0) {
      writer.uint32(16).uint64(message.lotteryId);
    }
    if (message.betAmount !== 0) {
      writer.uint32(24).uint64(message.betAmount);
    }
    if (message.bidder !== "") {
      writer.uint32(34).string(message.bidder);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StoredBid {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStoredBid();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.lotteryId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.betAmount = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.bidder = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoredBid {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      lotteryId: isSet(object.lotteryId) ? Number(object.lotteryId) : 0,
      betAmount: isSet(object.betAmount) ? Number(object.betAmount) : 0,
      bidder: isSet(object.bidder) ? String(object.bidder) : "",
    };
  },

  toJSON(message: StoredBid): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.lotteryId !== undefined && (obj.lotteryId = Math.round(message.lotteryId));
    message.betAmount !== undefined && (obj.betAmount = Math.round(message.betAmount));
    message.bidder !== undefined && (obj.bidder = message.bidder);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<StoredBid>, I>>(object: I): StoredBid {
    const message = createBaseStoredBid();
    message.index = object.index ?? "";
    message.lotteryId = object.lotteryId ?? 0;
    message.betAmount = object.betAmount ?? 0;
    message.bidder = object.bidder ?? "";
    return message;
  },
};

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
