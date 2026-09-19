from openai import OpenAI

from models import HistoricalEvent
from prompts import EXTRACTION_SYSTEM_PROMPT


MODEL = "qwen3.5:9b"
TIMEOUT = 300

client = OpenAI(
    base_url="http://localhost:1234/v1",
    api_key="lm-studio",
    timeout=TIMEOUT,
)


def extract(text: str) -> list[HistoricalEvent]:
    response = client.chat.completions.create(
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
        extra_body={
            "chat_template_kwargs": {
                "enable_thinking": False,
            }
        },
    )

    message = response.choices[0].message

    # print("\n=== THINKING ===")
    # print(getattr(message, "reasoning_content", None))

    print("\n=== CONTENT ===")
    print(message.content)

    print("\n=== FINISH ===")
    print(response.choices[0].finish_reason)

    from pydantic import TypeAdapter

    adapter = TypeAdapter(list[HistoricalEvent])

    return adapter.validate_json(message.content)