# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: src/domain/referral/v1/referral.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import field_mask_pb2 as google_dot_protobuf_dot_field__mask__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from src.domain.referral.v1 import commands_pb2 as src_dot_domain_dot_referral_dot_v1_dot_commands__pb2
from src.domain.referral.v1 import events_pb2 as src_dot_domain_dot_referral_dot_v1_dot_events__pb2
from src.domain.referral.v1 import queries_pb2 as src_dot_domain_dot_referral_dot_v1_dot_queries__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n%src/domain/referral/v1/referral.proto\x12\x12\x64omain.referral.v1\x1a google/protobuf/field_mask.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a%src/domain/referral/v1/commands.proto\x1a#src/domain/referral/v1/events.proto\x1a$src/domain/referral/v1/queries.proto\"\xf8\x01\n\x08Referral\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12\x17\n\x07user_id\x18\x03 \x01(\tR\x06userId\x12\x39\n\ncreated_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x39\n\nupdated_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x39\n\nfield_mask\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.FieldMaskR\tfieldMask\"G\n\tReferrals\x12:\n\treferrals\x18\x01 \x03(\x0b\x32\x1c.domain.referral.v1.ReferralR\treferralsB\x91\x01\n\x16\x63om.domain.referral.v1B\rReferralProtoP\x01\xa2\x02\x03\x44RX\xaa\x02\x12\x44omain.Referral.V1\xca\x02\x12\x44omain\\Referral\\V1\xe2\x02\x1e\x44omain\\Referral\\V1\\GPBMetadata\xea\x02\x14\x44omain::Referral::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'src.domain.referral.v1.referral_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\026com.domain.referral.v1B\rReferralProtoP\001\242\002\003DRX\252\002\022Domain.Referral.V1\312\002\022Domain\\Referral\\V1\342\002\036Domain\\Referral\\V1\\GPBMetadata\352\002\024Domain::Referral::V1'
  _globals['_REFERRAL']._serialized_start=243
  _globals['_REFERRAL']._serialized_end=491
  _globals['_REFERRALS']._serialized_start=493
  _globals['_REFERRALS']._serialized_end=564
# @@protoc_insertion_point(module_scope)
