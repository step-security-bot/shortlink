// source: infrastructure/rpc/cqrs/link/v1/link_command.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = (function() { return this || window || global || self || Function('return this')(); }).call(null);

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
goog.object.extend(proto, google_protobuf_empty_pb);
var domain_link_v1_link_pb = require('../../../../../domain/link/v1/link_pb.js');
goog.object.extend(proto, domain_link_v1_link_pb);
goog.exportSymbol('proto.infrastructure.rpc.cqrs.link.v1.AddRequest', null, global);
goog.exportSymbol('proto.infrastructure.rpc.cqrs.link.v1.AddResponse', null, global);
goog.exportSymbol('proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest', null, global);
goog.exportSymbol('proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest', null, global);
goog.exportSymbol('proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infrastructure.rpc.cqrs.link.v1.AddRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infrastructure.rpc.cqrs.link.v1.AddRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infrastructure.rpc.cqrs.link.v1.AddRequest.displayName = 'proto.infrastructure.rpc.cqrs.link.v1.AddRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infrastructure.rpc.cqrs.link.v1.AddResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infrastructure.rpc.cqrs.link.v1.AddResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infrastructure.rpc.cqrs.link.v1.AddResponse.displayName = 'proto.infrastructure.rpc.cqrs.link.v1.AddResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest.displayName = 'proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse.displayName = 'proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest.displayName = 'proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infrastructure.rpc.cqrs.link.v1.AddRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infrastructure.rpc.cqrs.link.v1.AddRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.AddRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.AddRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    link: (f = msg.getLink()) && domain_link_v1_link_pb.Link.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.AddRequest}
 */
proto.infrastructure.rpc.cqrs.link.v1.AddRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infrastructure.rpc.cqrs.link.v1.AddRequest;
  return proto.infrastructure.rpc.cqrs.link.v1.AddRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.AddRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.AddRequest}
 */
proto.infrastructure.rpc.cqrs.link.v1.AddRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new domain_link_v1_link_pb.Link;
      reader.readMessage(value,domain_link_v1_link_pb.Link.deserializeBinaryFromReader);
      msg.setLink(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infrastructure.rpc.cqrs.link.v1.AddRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infrastructure.rpc.cqrs.link.v1.AddRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.AddRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.AddRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getLink();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      domain_link_v1_link_pb.Link.serializeBinaryToWriter
    );
  }
};


/**
 * optional domain.link.v1.Link link = 1;
 * @return {?proto.domain.link.v1.Link}
 */
proto.infrastructure.rpc.cqrs.link.v1.AddRequest.prototype.getLink = function() {
  return /** @type{?proto.domain.link.v1.Link} */ (
    jspb.Message.getWrapperField(this, domain_link_v1_link_pb.Link, 1));
};


/**
 * @param {?proto.domain.link.v1.Link|undefined} value
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.AddRequest} returns this
*/
proto.infrastructure.rpc.cqrs.link.v1.AddRequest.prototype.setLink = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.AddRequest} returns this
 */
proto.infrastructure.rpc.cqrs.link.v1.AddRequest.prototype.clearLink = function() {
  return this.setLink(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infrastructure.rpc.cqrs.link.v1.AddRequest.prototype.hasLink = function() {
  return jspb.Message.getField(this, 1) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infrastructure.rpc.cqrs.link.v1.AddResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infrastructure.rpc.cqrs.link.v1.AddResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.AddResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.AddResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    link: (f = msg.getLink()) && domain_link_v1_link_pb.Link.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.AddResponse}
 */
proto.infrastructure.rpc.cqrs.link.v1.AddResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infrastructure.rpc.cqrs.link.v1.AddResponse;
  return proto.infrastructure.rpc.cqrs.link.v1.AddResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.AddResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.AddResponse}
 */
proto.infrastructure.rpc.cqrs.link.v1.AddResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new domain_link_v1_link_pb.Link;
      reader.readMessage(value,domain_link_v1_link_pb.Link.deserializeBinaryFromReader);
      msg.setLink(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infrastructure.rpc.cqrs.link.v1.AddResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infrastructure.rpc.cqrs.link.v1.AddResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.AddResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.AddResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getLink();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      domain_link_v1_link_pb.Link.serializeBinaryToWriter
    );
  }
};


/**
 * optional domain.link.v1.Link link = 1;
 * @return {?proto.domain.link.v1.Link}
 */
proto.infrastructure.rpc.cqrs.link.v1.AddResponse.prototype.getLink = function() {
  return /** @type{?proto.domain.link.v1.Link} */ (
    jspb.Message.getWrapperField(this, domain_link_v1_link_pb.Link, 1));
};


/**
 * @param {?proto.domain.link.v1.Link|undefined} value
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.AddResponse} returns this
*/
proto.infrastructure.rpc.cqrs.link.v1.AddResponse.prototype.setLink = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.AddResponse} returns this
 */
proto.infrastructure.rpc.cqrs.link.v1.AddResponse.prototype.clearLink = function() {
  return this.setLink(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infrastructure.rpc.cqrs.link.v1.AddResponse.prototype.hasLink = function() {
  return jspb.Message.getField(this, 1) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    link: (f = msg.getLink()) && domain_link_v1_link_pb.Link.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest}
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest;
  return proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest}
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new domain_link_v1_link_pb.Link;
      reader.readMessage(value,domain_link_v1_link_pb.Link.deserializeBinaryFromReader);
      msg.setLink(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getLink();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      domain_link_v1_link_pb.Link.serializeBinaryToWriter
    );
  }
};


/**
 * optional domain.link.v1.Link link = 1;
 * @return {?proto.domain.link.v1.Link}
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest.prototype.getLink = function() {
  return /** @type{?proto.domain.link.v1.Link} */ (
    jspb.Message.getWrapperField(this, domain_link_v1_link_pb.Link, 1));
};


/**
 * @param {?proto.domain.link.v1.Link|undefined} value
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest} returns this
*/
proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest.prototype.setLink = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest} returns this
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest.prototype.clearLink = function() {
  return this.setLink(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateRequest.prototype.hasLink = function() {
  return jspb.Message.getField(this, 1) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    link: (f = msg.getLink()) && domain_link_v1_link_pb.Link.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse}
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse;
  return proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse}
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new domain_link_v1_link_pb.Link;
      reader.readMessage(value,domain_link_v1_link_pb.Link.deserializeBinaryFromReader);
      msg.setLink(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getLink();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      domain_link_v1_link_pb.Link.serializeBinaryToWriter
    );
  }
};


/**
 * optional domain.link.v1.Link link = 1;
 * @return {?proto.domain.link.v1.Link}
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse.prototype.getLink = function() {
  return /** @type{?proto.domain.link.v1.Link} */ (
    jspb.Message.getWrapperField(this, domain_link_v1_link_pb.Link, 1));
};


/**
 * @param {?proto.domain.link.v1.Link|undefined} value
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse} returns this
*/
proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse.prototype.setLink = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse} returns this
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse.prototype.clearLink = function() {
  return this.setLink(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infrastructure.rpc.cqrs.link.v1.UpdateResponse.prototype.hasLink = function() {
  return jspb.Message.getField(this, 1) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    hash: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest}
 */
proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest;
  return proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest}
 */
proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setHash(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHash();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string hash = 1;
 * @return {string}
 */
proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest.prototype.getHash = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest} returns this
 */
proto.infrastructure.rpc.cqrs.link.v1.DeleteRequest.prototype.setHash = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


goog.object.extend(exports, proto.infrastructure.rpc.cqrs.link.v1);
