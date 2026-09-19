# from pathlib import Path

# from models import Ingredient
# from chunker import chunk_text
# from extractor import extract


# def main():
#     input_file = Path("input/coffee.txt")
#     output_file = Path("output/coffee_raw.json")

#     text = input_file.read_text(encoding="utf-8")

#     chunks = chunk_text(
#         text,
#         paragraphs_per_chunk=3,
#     )

#     print(f"Found {len(chunks)} chunks")

#     all_events = []
#     first_ingredient = None

#     for i, chunk in enumerate(chunks, start=1):
#         print(f"Extracting chunk {i}/{len(chunks)}...")

#         ingredient = extract(chunk)

#         if first_ingredient is None:
#             first_ingredient = ingredient

#         for event in ingredient.events:
#             print(f"  → {event.title}")

#         all_events.extend(ingredient.events)

#     result = Ingredient(
#         name=first_ingredient.name,
#         slug=first_ingredient.slug,
#         description=first_ingredient.description,
#         events=all_events,
#     )

#     output_file.parent.mkdir(parents=True, exist_ok=True)

#     output_file.write_text(
#         result.model_dump_json(indent=2),
#         encoding="utf-8",
#     )

#     print(f"Saved to {output_file}")


# if __name__ == "__main__":
#     main()

from pathlib import Path
import json

from models import HistoricalEvent
from chunker import chunk_text
from extractor import extract


def main():
    input_file = Path("input/coffee.txt")
    output_file = Path("output/coffee_raw.json")

    text = input_file.read_text(encoding="utf-8")

    chunks = chunk_text(
        text,
        paragraphs_per_chunk=3,
    )

    print(f"Found {len(chunks)} chunks")

    all_events: list[HistoricalEvent] = []

    for i, chunk in enumerate(chunks, start=1):
        print(f"Extracting chunk {i}/{len(chunks)}...")

        events = extract(chunk)

        # for event in events:
        #     print(f"  → {event.title}")

        all_events.extend(events)

    output_file.parent.mkdir(parents=True, exist_ok=True)

    output_file.write_text(
        json.dumps(
            [event.model_dump() for event in all_events],
            indent=2,
            ensure_ascii=False,
        ),
        encoding="utf-8",
    )

    print(f"Saved to {output_file}")


if __name__ == "__main__":
    main()