/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Params } from "./params";
import { StoredBid } from "./stored_bid";
import { StoredLottery } from "./stored_lottery";
import { SystemInfo } from "./system_info";

export const protobufPackage = "dedicateddev.lottery.lottery";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetSystemInfoRequest {
}

export interface QueryGetSystemInfoResponse {
  SystemInfo: SystemInfo | undefined;
}

export interface QueryGetStoredLotteryRequest {
  index: string;
}

export interface QueryGetStoredLotteryResponse {
  storedLottery: StoredLottery | undefined;
}

export interface QueryAllStoredLotteryRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllStoredLotteryResponse {
  storedLottery: StoredLottery[];
  pagination: PageResponse | undefined;
}

export interface QueryGetStoredBidRequest {
  index: string;
}

export interface QueryGetStoredBidResponse {
  storedBid: StoredBid | undefined;
}

export interface QueryAllStoredBidRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllStoredBidResponse {
  storedBid: StoredBid[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetSystemInfoRequest(): QueryGetSystemInfoRequest {
  return {};
}

export const QueryGetSystemInfoRequest = {
  encode(_: QueryGetSystemInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetSystemInfoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetSystemInfoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetSystemInfoRequest {
    return {};
  },

  toJSON(_: QueryGetSystemInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetSystemInfoRequest>, I>>(_: I): QueryGetSystemInfoRequest {
    const message = createBaseQueryGetSystemInfoRequest();
    return message;
  },
};

function createBaseQueryGetSystemInfoResponse(): QueryGetSystemInfoResponse {
  return { SystemInfo: undefined };
}

export const QueryGetSystemInfoResponse = {
  encode(message: QueryGetSystemInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.SystemInfo !== undefined) {
      SystemInfo.encode(message.SystemInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetSystemInfoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetSystemInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.SystemInfo = SystemInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSystemInfoResponse {
    return { SystemInfo: isSet(object.SystemInfo) ? SystemInfo.fromJSON(object.SystemInfo) : undefined };
  },

  toJSON(message: QueryGetSystemInfoResponse): unknown {
    const obj: any = {};
    message.SystemInfo !== undefined
      && (obj.SystemInfo = message.SystemInfo ? SystemInfo.toJSON(message.SystemInfo) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetSystemInfoResponse>, I>>(object: I): QueryGetSystemInfoResponse {
    const message = createBaseQueryGetSystemInfoResponse();
    message.SystemInfo = (object.SystemInfo !== undefined && object.SystemInfo !== null)
      ? SystemInfo.fromPartial(object.SystemInfo)
      : undefined;
    return message;
  },
};

function createBaseQueryGetStoredLotteryRequest(): QueryGetStoredLotteryRequest {
  return { index: "" };
}

export const QueryGetStoredLotteryRequest = {
  encode(message: QueryGetStoredLotteryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetStoredLotteryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetStoredLotteryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredLotteryRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetStoredLotteryRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetStoredLotteryRequest>, I>>(object: I): QueryGetStoredLotteryRequest {
    const message = createBaseQueryGetStoredLotteryRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetStoredLotteryResponse(): QueryGetStoredLotteryResponse {
  return { storedLottery: undefined };
}

export const QueryGetStoredLotteryResponse = {
  encode(message: QueryGetStoredLotteryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.storedLottery !== undefined) {
      StoredLottery.encode(message.storedLottery, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetStoredLotteryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetStoredLotteryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedLottery = StoredLottery.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredLotteryResponse {
    return { storedLottery: isSet(object.storedLottery) ? StoredLottery.fromJSON(object.storedLottery) : undefined };
  },

  toJSON(message: QueryGetStoredLotteryResponse): unknown {
    const obj: any = {};
    message.storedLottery !== undefined
      && (obj.storedLottery = message.storedLottery ? StoredLottery.toJSON(message.storedLottery) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetStoredLotteryResponse>, I>>(
    object: I,
  ): QueryGetStoredLotteryResponse {
    const message = createBaseQueryGetStoredLotteryResponse();
    message.storedLottery = (object.storedLottery !== undefined && object.storedLottery !== null)
      ? StoredLottery.fromPartial(object.storedLottery)
      : undefined;
    return message;
  },
};

function createBaseQueryAllStoredLotteryRequest(): QueryAllStoredLotteryRequest {
  return { pagination: undefined };
}

export const QueryAllStoredLotteryRequest = {
  encode(message: QueryAllStoredLotteryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllStoredLotteryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllStoredLotteryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredLotteryRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllStoredLotteryRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllStoredLotteryRequest>, I>>(object: I): QueryAllStoredLotteryRequest {
    const message = createBaseQueryAllStoredLotteryRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllStoredLotteryResponse(): QueryAllStoredLotteryResponse {
  return { storedLottery: [], pagination: undefined };
}

export const QueryAllStoredLotteryResponse = {
  encode(message: QueryAllStoredLotteryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.storedLottery) {
      StoredLottery.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllStoredLotteryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllStoredLotteryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedLottery.push(StoredLottery.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredLotteryResponse {
    return {
      storedLottery: Array.isArray(object?.storedLottery)
        ? object.storedLottery.map((e: any) => StoredLottery.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllStoredLotteryResponse): unknown {
    const obj: any = {};
    if (message.storedLottery) {
      obj.storedLottery = message.storedLottery.map((e) => e ? StoredLottery.toJSON(e) : undefined);
    } else {
      obj.storedLottery = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllStoredLotteryResponse>, I>>(
    object: I,
  ): QueryAllStoredLotteryResponse {
    const message = createBaseQueryAllStoredLotteryResponse();
    message.storedLottery = object.storedLottery?.map((e) => StoredLottery.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetStoredBidRequest(): QueryGetStoredBidRequest {
  return { index: "" };
}

export const QueryGetStoredBidRequest = {
  encode(message: QueryGetStoredBidRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetStoredBidRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetStoredBidRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredBidRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetStoredBidRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetStoredBidRequest>, I>>(object: I): QueryGetStoredBidRequest {
    const message = createBaseQueryGetStoredBidRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetStoredBidResponse(): QueryGetStoredBidResponse {
  return { storedBid: undefined };
}

export const QueryGetStoredBidResponse = {
  encode(message: QueryGetStoredBidResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.storedBid !== undefined) {
      StoredBid.encode(message.storedBid, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetStoredBidResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetStoredBidResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedBid = StoredBid.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredBidResponse {
    return { storedBid: isSet(object.storedBid) ? StoredBid.fromJSON(object.storedBid) : undefined };
  },

  toJSON(message: QueryGetStoredBidResponse): unknown {
    const obj: any = {};
    message.storedBid !== undefined
      && (obj.storedBid = message.storedBid ? StoredBid.toJSON(message.storedBid) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetStoredBidResponse>, I>>(object: I): QueryGetStoredBidResponse {
    const message = createBaseQueryGetStoredBidResponse();
    message.storedBid = (object.storedBid !== undefined && object.storedBid !== null)
      ? StoredBid.fromPartial(object.storedBid)
      : undefined;
    return message;
  },
};

function createBaseQueryAllStoredBidRequest(): QueryAllStoredBidRequest {
  return { pagination: undefined };
}

export const QueryAllStoredBidRequest = {
  encode(message: QueryAllStoredBidRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllStoredBidRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllStoredBidRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredBidRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllStoredBidRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllStoredBidRequest>, I>>(object: I): QueryAllStoredBidRequest {
    const message = createBaseQueryAllStoredBidRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllStoredBidResponse(): QueryAllStoredBidResponse {
  return { storedBid: [], pagination: undefined };
}

export const QueryAllStoredBidResponse = {
  encode(message: QueryAllStoredBidResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.storedBid) {
      StoredBid.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllStoredBidResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllStoredBidResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedBid.push(StoredBid.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredBidResponse {
    return {
      storedBid: Array.isArray(object?.storedBid) ? object.storedBid.map((e: any) => StoredBid.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllStoredBidResponse): unknown {
    const obj: any = {};
    if (message.storedBid) {
      obj.storedBid = message.storedBid.map((e) => e ? StoredBid.toJSON(e) : undefined);
    } else {
      obj.storedBid = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllStoredBidResponse>, I>>(object: I): QueryAllStoredBidResponse {
    const message = createBaseQueryAllStoredBidResponse();
    message.storedBid = object.storedBid?.map((e) => StoredBid.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a SystemInfo by index. */
  SystemInfo(request: QueryGetSystemInfoRequest): Promise<QueryGetSystemInfoResponse>;
  /** Queries a StoredLottery by index. */
  StoredLottery(request: QueryGetStoredLotteryRequest): Promise<QueryGetStoredLotteryResponse>;
  /** Queries a list of StoredLottery items. */
  StoredLotteryAll(request: QueryAllStoredLotteryRequest): Promise<QueryAllStoredLotteryResponse>;
  /** Queries a StoredBid by index. */
  StoredBid(request: QueryGetStoredBidRequest): Promise<QueryGetStoredBidResponse>;
  /** Queries a list of StoredBid items. */
  StoredBidAll(request: QueryAllStoredBidRequest): Promise<QueryAllStoredBidResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.SystemInfo = this.SystemInfo.bind(this);
    this.StoredLottery = this.StoredLottery.bind(this);
    this.StoredLotteryAll = this.StoredLotteryAll.bind(this);
    this.StoredBid = this.StoredBid.bind(this);
    this.StoredBidAll = this.StoredBidAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("dedicateddev.lottery.lottery.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  SystemInfo(request: QueryGetSystemInfoRequest): Promise<QueryGetSystemInfoResponse> {
    const data = QueryGetSystemInfoRequest.encode(request).finish();
    const promise = this.rpc.request("dedicateddev.lottery.lottery.Query", "SystemInfo", data);
    return promise.then((data) => QueryGetSystemInfoResponse.decode(new _m0.Reader(data)));
  }

  StoredLottery(request: QueryGetStoredLotteryRequest): Promise<QueryGetStoredLotteryResponse> {
    const data = QueryGetStoredLotteryRequest.encode(request).finish();
    const promise = this.rpc.request("dedicateddev.lottery.lottery.Query", "StoredLottery", data);
    return promise.then((data) => QueryGetStoredLotteryResponse.decode(new _m0.Reader(data)));
  }

  StoredLotteryAll(request: QueryAllStoredLotteryRequest): Promise<QueryAllStoredLotteryResponse> {
    const data = QueryAllStoredLotteryRequest.encode(request).finish();
    const promise = this.rpc.request("dedicateddev.lottery.lottery.Query", "StoredLotteryAll", data);
    return promise.then((data) => QueryAllStoredLotteryResponse.decode(new _m0.Reader(data)));
  }

  StoredBid(request: QueryGetStoredBidRequest): Promise<QueryGetStoredBidResponse> {
    const data = QueryGetStoredBidRequest.encode(request).finish();
    const promise = this.rpc.request("dedicateddev.lottery.lottery.Query", "StoredBid", data);
    return promise.then((data) => QueryGetStoredBidResponse.decode(new _m0.Reader(data)));
  }

  StoredBidAll(request: QueryAllStoredBidRequest): Promise<QueryAllStoredBidResponse> {
    const data = QueryAllStoredBidRequest.encode(request).finish();
    const promise = this.rpc.request("dedicateddev.lottery.lottery.Query", "StoredBidAll", data);
    return promise.then((data) => QueryAllStoredBidResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
