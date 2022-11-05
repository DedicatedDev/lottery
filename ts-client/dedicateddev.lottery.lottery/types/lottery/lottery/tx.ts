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

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateLottery(request: MsgCreateLottery): Promise<MsgCreateLotteryResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateLottery = this.CreateLottery.bind(this);
  }
  CreateLottery(request: MsgCreateLottery): Promise<MsgCreateLotteryResponse> {
    const data = MsgCreateLottery.encode(request).finish();
    const promise = this.rpc.request("dedicateddev.lottery.lottery.Msg", "CreateLottery", data);
    return promise.then((data) => MsgCreateLotteryResponse.decode(new _m0.Reader(data)));
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
