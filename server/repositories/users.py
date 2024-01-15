import logging
from typing import Optional, Dict, TypedDict

from bson import ObjectId
from pymongo.collection import Collection
from pymongo.cursor import Cursor
from pymongo.database import Database
from pymongo.errors import PyMongoError


class UsersModel(TypedDict):
    _id: Optional[ObjectId]
    name: str


class UsersRepository:
    def __init__(self, db: Database) -> None:
        self.col: Collection[UsersModel] = db["user"]

    def create_user(self, name: str) -> None:
        try:
            result = self.col.insert_one(UsersModel(name=name))
            logging.info(f"Inserted user with ID: {result.inserted_id}")
        except PyMongoError as error:
            logging.error(f"MongoDB error: {error}")
            raise

    def get_users(self, pagination: Dict[str, int]) -> Optional[Cursor[UsersModel]]:
        try:
            return self.col.find().skip(int(pagination["skip"])).limit(int(pagination["limit"]))
        except PyMongoError as error:
            logging.error(f"MongoDB error: {error}")
            raise

    def get_user(self, id: str) -> Optional[UsersModel]:
        try:
            result: UsersModel | None = self.col.find_one(filter={"_id": ObjectId(id)})
            if result:
                logging.info(f"User ID: {id}, Name: {result['name']}")
                return result
            else:
                logging.warning(f"User not found with ID: {id}")
                return None
        except PyMongoError as error:
            logging.error(f"MongoDB error: {error}")
            raise
