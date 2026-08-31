from pathlib import Path

from ollama import chat
from models import Ingredient
from prompts import EXTRACTION_PROMPT

MODEL = "qwen3:4b"

def extract(text: str) -> Ingredient:
    response = chat(
        model=MODEL,
        messages=[
            {
                "role": "system",
                "content": EXTRACTION_PROMPT,
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

def main():
    input_file = Path("input/tomato.txt")
    output_file = Path("output/tomato.json")

    text = input_file.read_text(encoding="utf-8")

    print(f"Extracting {input_file}...")

    ingredient = extract(text)

    output_file.parent.mkdir(parents=True, exist_ok=True)

    output_file.write_text(
        ingredient.model_dump_json(indent=2),
        encoding="utf-8",
    )

    print(f"Saved to {output_file}")


if __name__ == "__main__":
    main()