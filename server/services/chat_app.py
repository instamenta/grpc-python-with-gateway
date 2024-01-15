import logging

import google.protobuf.wrappers_pb2 as wp
import grpc

import chat_pb2 as proto_messages
import chat_pb2_grpc as proto
from server.repositories.messages import MessagesRepository
from server.repositories.users import UsersRepository


class ChatAppServicer(proto.ChatServiceServicer):
    def __init__(self, users_repo: UsersRepository, messages_repo: MessagesRepository):
        self.users_repo = users_repo
        self.messages_repo = messages_repo

    def CreateUser(self, request, context):
        if not request.name:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details("Name is required")
            return proto_messages.Empty()
        try:
            self.users_repo.create_user(request.name)
            return proto_messages.Empty()
        except Exception as e:
            logging.error(f"Error creating user: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(f"Error creating user: {e}")
            return proto_messages.Empty()

    def GetUser(self, request, context):
        if not request.HasField("id"):
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details("ID is required")
            return proto_messages.Empty()
        try:
            user = self.users_repo.get_user(request.id.value)
            if user:
                return proto_messages.User(
                    id=wp.StringValue(value=str(user['_id'])),
                    name=wp.StringValue(value=user['name']))
            else:
                context.set_code(grpc.StatusCode.NOT_FOUND)
                context.set_details(f"User not found with ID: {request.id}")
                return proto_messages.Empty()
        except Exception as e:
            logging.error(f"Error getting user: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(f"Error getting user: {e}")
            return proto_messages.Empty()

    def SendMessage(self, request, context):
        print(request)
        if not request.HasField("sender") or not request.HasField("recipient") or not request.HasField("content"):
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details("Sender, recipient, and content are required")
            return proto_messages.Empty()
        try:
            self.messages_repo.create_message(sender=request.sender.value,
                                              recipient=request.recipient.value,
                                              content=request.content.value)
            return proto_messages.Empty()
        except Exception as e:
            logging.error(f"Error sending message: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(f"Error sending message: {e}")
            return proto_messages.Empty()

    def GetMultipleUsers(self, request, context):
        try:
            users = self.users_repo.get_users(
                {"skip": request.skip.value if request.HasField('skip') else 0,
                 "limit": request.limit.value if request.HasField('limit') else 10})
            for user in users:
                yield proto_messages.User(
                    id=wp.StringValue(value=str(user["_id"])),
                    name=wp.StringValue(value=user["name"]))
        except Exception as e:
            logging.error(f"Error getting multiple users: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(f"Error getting multiple users: {e}")
            return proto_messages.User()

    def SendMultipleMessages(self, request_iterator, context):
        try:
            for request in request_iterator:
                if not request.sender or not request.recipient or not request.content:
                    context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
                    context.set_details("Sender, recipient, and content are required for each message")
                    return proto_messages.Empty()

                self.messages_repo.create_message(sender=request.sender,
                                                  recipient=request.recipient,
                                                  content=request.content)

            return proto_messages.Empty()
        except Exception as e:
            logging.error(f"Error sending multiple messages: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(f"Error sending multiple messages: {e}")
            return proto_messages.Empty()

    def GetMessages(self, request, context):
        if not request.HasField("sender") or not request.HasField("recipient"):
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details("Sender and recipient are required")
            return proto_messages.Empty()

        pagination = {"skip": request.skip.value if request.HasField('skip') else 0,
                      "limit": request.limit.value if request.HasField('limit') else 10}
        try:
            messages = self.messages_repo.get_messages_by_sender_and_recipient(
                sender=request.sender.value,
                recipient=request.recipient.value,
                pagination=pagination)
            for message in messages:
                yield proto_messages.Message(id=wp.StringValue(value=str(message["_id"])),
                                             sender=wp.StringValue(value=str(message['sender'])),
                                             recipient=wp.StringValue(value=str(message['recipient'])),
                                             content=wp.StringValue(value=message['content']))
        except Exception as e:
            logging.error(f"Error getting messages: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(f"Error getting messages: {e}")
            return proto_messages.Empty()
