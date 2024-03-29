# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chat.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nchat.proto\x12\x04\x63hat\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1cgoogle/api/annotations.proto\"?\n\x11\x43reateUserRequest\x12*\n\x04name\x18\x01 \x01(\x0b\x32\x1c.google.protobuf.StringValue\"c\n\nPagination\x12)\n\x04skip\x18\x01 \x01(\x0b\x32\x1b.google.protobuf.Int32Value\x12*\n\x05limit\x18\x02 \x01(\x0b\x32\x1b.google.protobuf.Int32Value\"3\n\x07GetById\x12(\n\x02id\x18\x01 \x01(\x0b\x32\x1c.google.protobuf.StringValue\"\xca\x01\n\x12GetMessagesRequest\x12,\n\x06sender\x18\x01 \x01(\x0b\x32\x1c.google.protobuf.StringValue\x12/\n\trecipient\x18\x02 \x01(\x0b\x32\x1c.google.protobuf.StringValue\x12)\n\x04skip\x18\x03 \x01(\x0b\x32\x1b.google.protobuf.Int32Value\x12*\n\x05limit\x18\x04 \x01(\x0b\x32\x1b.google.protobuf.Int32Value\"\\\n\x04User\x12(\n\x02id\x18\x01 \x01(\x0b\x32\x1c.google.protobuf.StringValue\x12*\n\x04name\x18\x02 \x01(\x0b\x32\x1c.google.protobuf.StringValue\"\xb8\x02\n\x07Message\x12,\n\x06sender\x18\x01 \x01(\x0b\x32\x1c.google.protobuf.StringValue\x12/\n\trecipient\x18\x02 \x01(\x0b\x32\x1c.google.protobuf.StringValue\x12-\n\x07\x63ontent\x18\x03 \x01(\x0b\x32\x1c.google.protobuf.StringValue\x12(\n\x02id\x18\x04 \x01(\x0b\x32\x1c.google.protobuf.StringValue\x12\x43\n\x0emessage_status\x18\x0c \x01(\x0e\x32\x1b.chat.Message.MessageStatusR\x0emessage_status\"0\n\rMessageStatus\x12\x08\n\x04SEEN\x10\x00\x12\x08\n\x04SEND\x10\x01\x12\x0b\n\x07PENDING\x10\x02\"\x07\n\x05\x45mpty2\x83\x04\n\x0b\x43hatService\x12N\n\nCreateUser\x12\x17.chat.CreateUserRequest\x1a\x0b.chat.Empty\"\x1a\x82\xd3\xe4\x93\x02\x14\"\x0c/api/v1/user:\x04name\x12?\n\x07GetUser\x12\r.chat.GetById\x1a\n.chat.User\"\x19\x82\xd3\xe4\x93\x02\x13\x12\x11/api/v1/user/{id}\x12\x61\n\x0bSendMessage\x12\r.chat.Message\x1a\x0b.chat.Empty\"6\x82\xd3\xe4\x93\x02\x30\"%/api/v1/messages/{sender}/{recipient}:\x07\x63ontent\x12H\n\x10GetMultipleUsers\x12\x10.chat.Pagination\x1a\n.chat.User\"\x14\x82\xd3\xe4\x93\x02\x0e\x12\x0c/api/v1/user0\x01\x12M\n\x14SendMultipleMessages\x12\r.chat.Message\x1a\x0b.chat.Empty\"\x17\x82\xd3\xe4\x93\x02\x11\"\x0f/api/deprecated(\x01\x12g\n\x0bGetMessages\x12\x18.chat.GetMessagesRequest\x1a\r.chat.Message\"-\x82\xd3\xe4\x93\x02\'\x12%/api/v1/messages/{sender}/{recipient}0\x01\x42\x07Z\x05\x63hat/b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'chat_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\005chat/'
  _globals['_CHATSERVICE'].methods_by_name['CreateUser']._options = None
  _globals['_CHATSERVICE'].methods_by_name['CreateUser']._serialized_options = b'\202\323\344\223\002\024\"\014/api/v1/user:\004name'
  _globals['_CHATSERVICE'].methods_by_name['GetUser']._options = None
  _globals['_CHATSERVICE'].methods_by_name['GetUser']._serialized_options = b'\202\323\344\223\002\023\022\021/api/v1/user/{id}'
  _globals['_CHATSERVICE'].methods_by_name['SendMessage']._options = None
  _globals['_CHATSERVICE'].methods_by_name['SendMessage']._serialized_options = b'\202\323\344\223\0020\"%/api/v1/messages/{sender}/{recipient}:\007content'
  _globals['_CHATSERVICE'].methods_by_name['GetMultipleUsers']._options = None
  _globals['_CHATSERVICE'].methods_by_name['GetMultipleUsers']._serialized_options = b'\202\323\344\223\002\016\022\014/api/v1/user'
  _globals['_CHATSERVICE'].methods_by_name['SendMultipleMessages']._options = None
  _globals['_CHATSERVICE'].methods_by_name['SendMultipleMessages']._serialized_options = b'\202\323\344\223\002\021\"\017/api/deprecated'
  _globals['_CHATSERVICE'].methods_by_name['GetMessages']._options = None
  _globals['_CHATSERVICE'].methods_by_name['GetMessages']._serialized_options = b'\202\323\344\223\002\'\022%/api/v1/messages/{sender}/{recipient}'
  _globals['_CREATEUSERREQUEST']._serialized_start=82
  _globals['_CREATEUSERREQUEST']._serialized_end=145
  _globals['_PAGINATION']._serialized_start=147
  _globals['_PAGINATION']._serialized_end=246
  _globals['_GETBYID']._serialized_start=248
  _globals['_GETBYID']._serialized_end=299
  _globals['_GETMESSAGESREQUEST']._serialized_start=302
  _globals['_GETMESSAGESREQUEST']._serialized_end=504
  _globals['_USER']._serialized_start=506
  _globals['_USER']._serialized_end=598
  _globals['_MESSAGE']._serialized_start=601
  _globals['_MESSAGE']._serialized_end=913
  _globals['_MESSAGE_MESSAGESTATUS']._serialized_start=865
  _globals['_MESSAGE_MESSAGESTATUS']._serialized_end=913
  _globals['_EMPTY']._serialized_start=915
  _globals['_EMPTY']._serialized_end=922
  _globals['_CHATSERVICE']._serialized_start=925
  _globals['_CHATSERVICE']._serialized_end=1440
# @@protoc_insertion_point(module_scope)
