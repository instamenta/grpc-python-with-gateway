import logging
from typing import TypedDict, Optional, Dict

from bson import ObjectId
from pymongo.collection import Collection
from pymongo.cursor import Cursor
from pymongo.database import Database
from pymongo.errors import PyMongoError


class MessageModel(TypedDict):
    _id: Optional[ObjectId]
    sender: ObjectId
    recipient: ObjectId
    content: str


class MessagesRepository:
    def __init__(self, db: Database) -> None:
        self.collection: Collection[MessageModel] = db["messages"]

    def get_messages_by_sender_and_recipient(self,
                                             sender: str,
                                             recipient: str,
                                             pagination: Dict[str, int]
                                             ) -> Optional[Cursor[MessageModel]]:
        try:
            filter_query = {"sender": ObjectId(sender),
                            "recipient": ObjectId(recipient)}
            return self.collection.find(filter_query).skip(pagination["skip"]).limit(pagination["limit"])
        except PyMongoError as e:
            logging.error(f"MongoDB error: {e}")
            raise

    def create_message(self, sender: str, recipient: str, content: str) -> None:
        try:
            doc = MessageModel(sender=ObjectId(sender),
                               recipient=ObjectId(recipient),
                               content=content)
            result = self.collection.insert_one(doc)
            logging.info(f"Inserted message with ID: {result.inserted_id}")
        except PyMongoError as e:
            logging.error(f"MongoDB error: {e}")
            raise

    def get_all_messages(self) -> Optional[Cursor[MessageModel]]:
        try:
            return self.collection.find()
        except PyMongoError as e:
            logging.error(f"MongoDB error: {e}")
            raise

    def delete_message(self, message_id: str) -> bool:
        try:
            message_id: ObjectId = ObjectId(message_id)
            result = self.collection.delete_one(filter={"_id": message_id})
            if result.deleted_count == 1:
                logging.info(f"Deleted message with ID: {message_id}")
                return True
            else:
                logging.warning(f"Message not found with ID: {message_id}")
                return False
        except PyMongoError as e:
            logging.error(f"MongoDB error: {e}")
            raise
