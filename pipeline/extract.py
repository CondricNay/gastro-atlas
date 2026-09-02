from pathlib import Path

from ollama import chat
from models import Ingredient
from prompts import EXTRACTION_SYSTEM_PROMPT


MODEL = "qwen3:8b"


def extract(text: str) -> Ingredient:
    response = chat(
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


def chunk_text(text: str, paragraphs_per_chunk: int = 3) -> list[str]:
    paragraphs = [
        p.strip()
        for p in text.split("\n\n")
        if p.strip()
    ]

    return [
        "\n\n".join(paragraphs[i:i + paragraphs_per_chunk])
        for i in range(0, len(paragraphs), paragraphs_per_chunk)
    ]

def main():
    input_file = Path("input/tomato.txt")
    output_file = Path("output/tomato_chunked.json")

    text = input_file.read_text(encoding="utf-8")

    chunks = chunk_text(text, paragraphs_per_chunk=3)

    print(f"Found {len(chunks)} chunks")

    all_events = []

    for i, chunk in enumerate(chunks, start=1):
        print(f"Extracting chunk {i}/{len(chunks)}...")

        ingredient = extract(chunk)

        all_events.extend(ingredient.events)

    # For now, just use the first ingredient's metadata
    result = Ingredient(
        name=ingredient.name,
        slug=ingredient.slug,
        description=ingredient.description,
        events=all_events,
    )

    output_file.parent.mkdir(parents=True, exist_ok=True)

    output_file.write_text(
        result.model_dump_json(indent=2),
        encoding="utf-8",
    )

    print(f"Saved to {output_file}")

if __name__ == "__main__":
    main()
