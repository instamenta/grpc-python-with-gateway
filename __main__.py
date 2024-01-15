import logging
from concurrent import futures

import grpc
from pymongo.mongo_client import MongoClient

import chat_pb2_grpc as proto
from repositories.messages import MessagesRepository
from repositories.users import UsersRepository
from services.chat_app import ChatAppServicer

uri = "mongodb+srv://janoopsi:janoopsi9999@clickercluster.ltycehn.mongodb.net/?retryWrites=true&w=majority"


def serve():
    logging.basicConfig(level=logging.DEBUG)
    try:
        client = MongoClient(uri)
        client.admin.command('ping')
        print("Pinged your deployment. You successfully connected to MongoDB!")
    except Exception as e:
        print(f"Failed to connect to mongodb Error: {e}")
        raise
    try:
        users_repo = UsersRepository(client.get_database("chat"))
        messages_repo = MessagesRepository(client.get_database("chat"))
        service = ChatAppServicer(users_repo=users_repo, messages_repo=messages_repo)

        server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        proto.add_ChatServiceServicer_to_server(service, server)
        server.add_insecure_port('[::]:50051')
        server.start()
        logging.info("Server started. Listening on port 50051.")
        server.wait_for_termination()
    except Exception as error:
        logging.error(f"Server failed with exception: {error}")


if __name__ == "__main__":
    serve()
