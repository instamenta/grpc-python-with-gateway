from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateUserRequest(_message.Message):
    __slots__ = ("name",)
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class Pagination(_message.Message):
    __slots__ = ("skip", "limit")
    SKIP_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    skip: _wrappers_pb2.Int32Value
    limit: _wrappers_pb2.Int32Value
    def __init__(self, skip: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class GetById(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class GetMessagesRequest(_message.Message):
    __slots__ = ("sender", "recipient", "skip", "limit")
    SENDER_FIELD_NUMBER: _ClassVar[int]
    RECIPIENT_FIELD_NUMBER: _ClassVar[int]
    SKIP_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    sender: _wrappers_pb2.StringValue
    recipient: _wrappers_pb2.StringValue
    skip: _wrappers_pb2.Int32Value
    limit: _wrappers_pb2.Int32Value
    def __init__(self, sender: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., recipient: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., skip: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class User(_message.Message):
    __slots__ = ("id", "name")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class Message(_message.Message):
    __slots__ = ("sender", "recipient", "content", "id", "message_status")
    class MessageStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        SEEN: _ClassVar[Message.MessageStatus]
        SEND: _ClassVar[Message.MessageStatus]
        PENDING: _ClassVar[Message.MessageStatus]
    SEEN: Message.MessageStatus
    SEND: Message.MessageStatus
    PENDING: Message.MessageStatus
    SENDER_FIELD_NUMBER: _ClassVar[int]
    RECIPIENT_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_STATUS_FIELD_NUMBER: _ClassVar[int]
    sender: _wrappers_pb2.StringValue
    recipient: _wrappers_pb2.StringValue
    content: _wrappers_pb2.StringValue
    id: _wrappers_pb2.StringValue
    message_status: Message.MessageStatus
    def __init__(self, sender: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., recipient: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., content: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., message_status: _Optional[_Union[Message.MessageStatus, str]] = ...) -> None: ...

class Empty(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
