from ollama import Client
from models import Ingredient
from prompts import EXTRACTION_SYSTEM_PROMPT


MODEL = "qwen3:8b"
MAX_RETRIES = 1
TIMEOUT = 90


client = Client(
    host="http://localhost:11434",
    timeout=TIMEOUT,
)


def extract(text: str) -> Ingredient:
    for attempt in range(1, MAX_RETRIES + 1):
        try:
            response = client.chat(
                model=MODEL,
                messages=[
                    {
                        "role": "system",
                        "content": EXTRACTION_SYSTEM_PROMPT,
                    },
                    {
                        "role": "user",
                        "content": text,
                    },
                ],
                think=False,
                format=Ingredient.model_json_schema(),
            )

            return Ingredient.model_validate_json(
                response.message.content
            )

        except Exception as error:
            print(
                f"  Attempt {attempt}/{MAX_RETRIES} failed: {error}"
            )

            if attempt == MAX_RETRIES:
                raise
