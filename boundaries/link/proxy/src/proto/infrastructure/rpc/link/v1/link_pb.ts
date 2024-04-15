// @generated by protoc-gen-es v1.8.0 with parameter "target=ts"
// @generated from file infrastructure/rpc/link/v1/link.proto (package infrastructure.rpc.link.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Link, Links } from "../../../../domain/link/v1/link_pb.js";

/**
 * GetRequest is the request message for LinkService.Get.
 *
 * @generated from message infrastructure.rpc.link.v1.GetRequest
 */
export class GetRequest extends Message<GetRequest> {
  /**
   * Hash is the hash of the link.
   *
   * @generated from field: string hash = 1;
   */
  hash = "";

  constructor(data?: PartialMessage<GetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "infrastructure.rpc.link.v1.GetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "hash", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRequest {
    return new GetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRequest {
    return new GetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRequest {
    return new GetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetRequest | PlainMessage<GetRequest> | undefined, b: GetRequest | PlainMessage<GetRequest> | undefined): boolean {
    return proto3.util.equals(GetRequest, a, b);
  }
}

/**
 * GetResponse is the response message for LinkService.Get.
 *
 * @generated from message infrastructure.rpc.link.v1.GetResponse
 */
export class GetResponse extends Message<GetResponse> {
  /**
   * Link is the link.
   *
   * @generated from field: domain.link.v1.Link link = 1;
   */
  link?: Link;

  constructor(data?: PartialMessage<GetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "infrastructure.rpc.link.v1.GetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "link", kind: "message", T: Link },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetResponse {
    return new GetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetResponse {
    return new GetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetResponse {
    return new GetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetResponse | PlainMessage<GetResponse> | undefined, b: GetResponse | PlainMessage<GetResponse> | undefined): boolean {
    return proto3.util.equals(GetResponse, a, b);
  }
}

/**
 * ListRequest is the request message for LinkService.List.
 *
 * @generated from message infrastructure.rpc.link.v1.ListRequest
 */
export class ListRequest extends Message<ListRequest> {
  /**
   * Filter is the filter of the list.
   *
   * @generated from field: string filter = 1;
   */
  filter = "";

  /**
   * Limit is the limit of the list.
   *
   * @generated from field: uint32 limit = 2;
   */
  limit = 0;

  /**
   * Cursor is the token of the list.
   *
   * @generated from field: string cursor = 3;
   */
  cursor = "";

  constructor(data?: PartialMessage<ListRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "infrastructure.rpc.link.v1.ListRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "filter", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "limit", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 3, name: "cursor", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListRequest {
    return new ListRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListRequest {
    return new ListRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListRequest {
    return new ListRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListRequest | PlainMessage<ListRequest> | undefined, b: ListRequest | PlainMessage<ListRequest> | undefined): boolean {
    return proto3.util.equals(ListRequest, a, b);
  }
}

/**
 * ListResponse is the response message for LinkService.List.
 *
 * @generated from message infrastructure.rpc.link.v1.ListResponse
 */
export class ListResponse extends Message<ListResponse> {
  /**
   * Links is the list of links.
   *
   * @generated from field: domain.link.v1.Links links = 1;
   */
  links?: Links;

  /**
   * Cursor is the next token of the list.
   *
   * @generated from field: string cursor = 2;
   */
  cursor = "";

  constructor(data?: PartialMessage<ListResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "infrastructure.rpc.link.v1.ListResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "links", kind: "message", T: Links },
    { no: 2, name: "cursor", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListResponse {
    return new ListResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListResponse {
    return new ListResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListResponse {
    return new ListResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListResponse | PlainMessage<ListResponse> | undefined, b: ListResponse | PlainMessage<ListResponse> | undefined): boolean {
    return proto3.util.equals(ListResponse, a, b);
  }
}

/**
 * AddRequest is the request message for LinkService.Add.
 *
 * @generated from message infrastructure.rpc.link.v1.AddRequest
 */
export class AddRequest extends Message<AddRequest> {
  /**
   * Link is the link.
   *
   * @generated from field: domain.link.v1.Link link = 1;
   */
  link?: Link;

  constructor(data?: PartialMessage<AddRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "infrastructure.rpc.link.v1.AddRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "link", kind: "message", T: Link },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddRequest {
    return new AddRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddRequest {
    return new AddRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddRequest {
    return new AddRequest().fromJsonString(jsonString, options);
  }

  static equals(a: AddRequest | PlainMessage<AddRequest> | undefined, b: AddRequest | PlainMessage<AddRequest> | undefined): boolean {
    return proto3.util.equals(AddRequest, a, b);
  }
}

/**
 * AddResponse is the response message for LinkService.Add.
 *
 * @generated from message infrastructure.rpc.link.v1.AddResponse
 */
export class AddResponse extends Message<AddResponse> {
  /**
   * Link is the link.
   *
   * @generated from field: domain.link.v1.Link link = 1;
   */
  link?: Link;

  constructor(data?: PartialMessage<AddResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "infrastructure.rpc.link.v1.AddResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "link", kind: "message", T: Link },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddResponse {
    return new AddResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddResponse {
    return new AddResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddResponse {
    return new AddResponse().fromJsonString(jsonString, options);
  }

  static equals(a: AddResponse | PlainMessage<AddResponse> | undefined, b: AddResponse | PlainMessage<AddResponse> | undefined): boolean {
    return proto3.util.equals(AddResponse, a, b);
  }
}

/**
 * UpdateRequest is the request message for LinkService.Update.
 *
 * @generated from message infrastructure.rpc.link.v1.UpdateRequest
 */
export class UpdateRequest extends Message<UpdateRequest> {
  /**
   * Link is the link.
   *
   * @generated from field: domain.link.v1.Link link = 1;
   */
  link?: Link;

  constructor(data?: PartialMessage<UpdateRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "infrastructure.rpc.link.v1.UpdateRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "link", kind: "message", T: Link },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateRequest {
    return new UpdateRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateRequest {
    return new UpdateRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateRequest {
    return new UpdateRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateRequest | PlainMessage<UpdateRequest> | undefined, b: UpdateRequest | PlainMessage<UpdateRequest> | undefined): boolean {
    return proto3.util.equals(UpdateRequest, a, b);
  }
}

/**
 * UpdateResponse is the response message for LinkService.Update.
 *
 * @generated from message infrastructure.rpc.link.v1.UpdateResponse
 */
export class UpdateResponse extends Message<UpdateResponse> {
  /**
   * Link is the link.
   *
   * @generated from field: domain.link.v1.Link link = 1;
   */
  link?: Link;

  constructor(data?: PartialMessage<UpdateResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "infrastructure.rpc.link.v1.UpdateResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "link", kind: "message", T: Link },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateResponse {
    return new UpdateResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateResponse {
    return new UpdateResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateResponse {
    return new UpdateResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateResponse | PlainMessage<UpdateResponse> | undefined, b: UpdateResponse | PlainMessage<UpdateResponse> | undefined): boolean {
    return proto3.util.equals(UpdateResponse, a, b);
  }
}

/**
 * DeleteRequest is the request message for LinkService.Delete.
 *
 * @generated from message infrastructure.rpc.link.v1.DeleteRequest
 */
export class DeleteRequest extends Message<DeleteRequest> {
  /**
   * Hash is the hash of the link.
   *
   * @generated from field: string hash = 1;
   */
  hash = "";

  constructor(data?: PartialMessage<DeleteRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "infrastructure.rpc.link.v1.DeleteRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "hash", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteRequest {
    return new DeleteRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteRequest {
    return new DeleteRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteRequest {
    return new DeleteRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteRequest | PlainMessage<DeleteRequest> | undefined, b: DeleteRequest | PlainMessage<DeleteRequest> | undefined): boolean {
    return proto3.util.equals(DeleteRequest, a, b);
  }
}

