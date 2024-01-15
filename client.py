import google.protobuf.wrappers_pb2 as wp
import grpc

import chat_pb2 as proto_messages
import chat_pb2_grpc as proto


def run():
    client = proto.ChatServiceStub(grpc.insecure_channel('localhost:50052'))

    response_iterator = client.GetMultipleUsers(
        proto_messages.Pagination(skip=wp.Int32Value(value=0),
                                  limit=wp.Int32Value(value=10)))

    for response in response_iterator:
        print(f"Received user: {response.id}, {response.name}")

    response = client.SendMessage(
        proto_messages.Message(
            sender=wp.StringValue(value="65a1b9b6a979dc92b9c15dd1"),
            content=wp.StringValue(value="content"),
            recipient=wp.StringValue(value="65a2b6672628c003caaa4cbb")
        )
    )

    print(response)


if __name__ == '__main__':
    run()
