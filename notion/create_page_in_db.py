import requests
import os
from dotenv import load_dotenv


def generate_current_date() -> str:
    """
    Generate current date in YYYY-MM-DD format.

    :return: Current date.
    """
    from datetime import datetime

    now = datetime.now()
    return now.strftime("%Y-%m-%d")


def load_environment_variables() -> str:
    """
    Load environment variables from .env file.

    :return: Integration token.
    """

    load_dotenv()

    token = os.getenv("INTEGRATION_TOKEN")
    return token


def filter_database_return_len(
        database_id: str, filter_property: str, filter_value: str
) -> int:
    """
    Filter a database by a property.

    :param database_id: ID of the database to filter
    :param filter_property: Property to filter by
    :param filter_value: Value to filter by
    :return: Number of results
    """
    headers = generate_headers(token=load_environment_variables())
    res = requests.post(
        f"https://api.notion.com/v1/databases/{database_id}/query",
        headers=headers,
        json={
            "filter": {
                "property": filter_property,
                "title": {
                    "equals": filter_value,
                },
            },
        },
    )

    parsed_res = res.json()
    return len(parsed_res["results"])


def generate_headers(token: str, api_version: str = "2022-06-28") -> dict:
    return {
        "Authorization": f"Bearer {token}",
        "Notion-Version": api_version,
        "Content-Type": "application/json",
    }


def create_page_in_db(title: str, database_id: str) -> str:
    """
    Create a new page in a database.

    :param title: Title of the page to create
    :param database_id: ID of the database to create the page in
    :return: String with the response from the API
    """
    headers = generate_headers(token=load_environment_variables())
    # requests.post("https://api.notion.com/v1/pages", headers=headers, json={})
    res = requests.post(
        "https://api.notion.com/v1/pages",
        headers=headers,
        json={
            "parent": {
                "database_id": database_id,
            },
            "properties": {
                "Name": {"title": [{"text": {"content": title}}]},
                "Tags": {
                    "multi_select": [{"id": "3b6f5ceb-099f-4521-a306-a9a385556f80"}]
                },
            },
        },
    )

    parsed_res = res.json()
    return parsed_res


if __name__ == "__main__":
    DATABASE_ID = "e834563bdcbf4dbb9643a54e3111b03d"
    current_date = generate_current_date()

    length = filter_database_return_len(
        database_id=DATABASE_ID, filter_property="Name", filter_value=current_date
    )
    if length != 0:
        raise NameError("Page already exists")

    create_page_in_db(title=current_date, database_id=DATABASE_ID)
