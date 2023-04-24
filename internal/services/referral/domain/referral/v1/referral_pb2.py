# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: domain/referral/v1/referral.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import field_mask_pb2 as google_dot_protobuf_dot_field__mask__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!domain/referral/v1/referral.proto\x12\x12\x64omain.referral.v1\x1a google/protobuf/field_mask.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"A\n\x12ReferralAddCommand\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12\x17\n\x07user_id\x18\x02 \x01(\tR\x06userId\"\x8f\x01\n\x15ReferralUpdateCommand\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12\x17\n\x07user_id\x18\x03 \x01(\tR\x06userId\x12\x39\n\nfield_mask\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.FieldMaskR\tfieldMask\"\'\n\x15ReferralDeleteCommand\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\"\"\n\x10ReferralGetQuery\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\",\n\x11ReferralListQuery\x12\x17\n\x07user_id\x18\x01 \x01(\tR\x06userId\"\xf8\x01\n\x08Referral\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12\x17\n\x07user_id\x18\x03 \x01(\tR\x06userId\x12\x39\n\ncreated_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x39\n\nupdated_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x39\n\nfield_mask\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.FieldMaskR\tfieldMask\"G\n\tReferrals\x12:\n\treferrals\x18\x01 \x03(\x0b\x32\x1c.domain.referral.v1.ReferralR\treferrals*\xae\x01\n\rReferralEvent\x12\x1e\n\x1aREFERRAL_EVENT_UNSPECIFIED\x10\x00\x12\x16\n\x12REFERRAL_EVENT_ADD\x10\x01\x12\x16\n\x12REFERRAL_EVENT_GET\x10\x02\x12\x17\n\x13REFERRAL_EVENT_LIST\x10\x03\x12\x19\n\x15REFERRAL_EVENT_UPDATE\x10\x04\x12\x19\n\x15REFERRAL_EVENT_DELETE\x10\x05*\x87\x01\n\x0fReferralCommand\x12 \n\x1cREFERRAL_COMMAND_UNSPECIFIED\x10\x00\x12\x18\n\x14REFERRAL_COMMAND_ADD\x10\x01\x12\x1b\n\x17REFERRAL_COMMAND_UPDATE\x10\x02\x12\x1b\n\x17REFERRAL_COMMAND_DELETE\x10\x03*`\n\rReferralQuery\x12\x1e\n\x1aREFERRAL_QUERY_UNSPECIFIED\x10\x00\x12\x16\n\x12REFERRAL_QUERY_GET\x10\x01\x12\x17\n\x13REFERRAL_QUERY_LIST\x10\x02\x42\x91\x01\n\x16\x63om.domain.referral.v1B\rReferralProtoP\x01\xa2\x02\x03\x44RX\xaa\x02\x12\x44omain.Referral.V1\xca\x02\x12\x44omain\\Referral\\V1\xe2\x02\x1e\x44omain\\Referral\\V1\\GPBMetadata\xea\x02\x14\x44omain::Referral::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'domain.referral.v1.referral_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\026com.domain.referral.v1B\rReferralProtoP\001\242\002\003DRX\252\002\022Domain.Referral.V1\312\002\022Domain\\Referral\\V1\342\002\036Domain\\Referral\\V1\\GPBMetadata\352\002\024Domain::Referral::V1'
  _globals['_REFERRALEVENT']._serialized_start=785
  _globals['_REFERRALEVENT']._serialized_end=959
  _globals['_REFERRALCOMMAND']._serialized_start=962
  _globals['_REFERRALCOMMAND']._serialized_end=1097
  _globals['_REFERRALQUERY']._serialized_start=1099
  _globals['_REFERRALQUERY']._serialized_end=1195
  _globals['_REFERRALADDCOMMAND']._serialized_start=124
  _globals['_REFERRALADDCOMMAND']._serialized_end=189
  _globals['_REFERRALUPDATECOMMAND']._serialized_start=192
  _globals['_REFERRALUPDATECOMMAND']._serialized_end=335
  _globals['_REFERRALDELETECOMMAND']._serialized_start=337
  _globals['_REFERRALDELETECOMMAND']._serialized_end=376
  _globals['_REFERRALGETQUERY']._serialized_start=378
  _globals['_REFERRALGETQUERY']._serialized_end=412
  _globals['_REFERRALLISTQUERY']._serialized_start=414
  _globals['_REFERRALLISTQUERY']._serialized_end=458
  _globals['_REFERRAL']._serialized_start=461
  _globals['_REFERRAL']._serialized_end=709
  _globals['_REFERRALS']._serialized_start=711
  _globals['_REFERRALS']._serialized_end=782
# @@protoc_insertion_point(module_scope)
