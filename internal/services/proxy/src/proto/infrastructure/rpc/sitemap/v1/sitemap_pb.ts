// @generated by protoc-gen-es v1.3.0 with parameter "target=ts"
// @generated from file infrastructure/rpc/sitemap/v1/sitemap.proto (package infrastructure.rpc.sitemap.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message infrastructure.rpc.sitemap.v1.ParseRequest
 */
export class ParseRequest extends Message<ParseRequest> {
  /**
   * @generated from field: string url = 1;
   */
  url = "";

  constructor(data?: PartialMessage<ParseRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "infrastructure.rpc.sitemap.v1.ParseRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ParseRequest {
    return new ParseRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ParseRequest {
    return new ParseRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ParseRequest {
    return new ParseRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ParseRequest | PlainMessage<ParseRequest> | undefined, b: ParseRequest | PlainMessage<ParseRequest> | undefined): boolean {
    return proto3.util.equals(ParseRequest, a, b);
  }
}

