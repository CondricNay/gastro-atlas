from pathlib import Path

from models import Ingredient


def normalize_string(value: str | None) -> str | None:
    if value is None:
        return None

    value = value.strip()

    if not value:
        return None

    return value


def normalize_event(event):
    event.title = normalize_string(event.title)
    event.description = normalize_string(event.description)
    event.time_period = normalize_string(event.time_period)
    event.entity = normalize_string(event.entity)
    event.location = normalize_string(event.location)

    return event


def validate_ingredient(ingredient: Ingredient) -> Ingredient:
    ingredient.events = [
        normalize_event(event)
        for event in ingredient.events
    ]

    return ingredient


def main():
    input_file = Path("output/coffee_raw.json")
    output_file = Path("output/coffee_processed.json")

    ingredient = Ingredient.model_validate_json(
        input_file.read_text(encoding="utf-8")
    )

    ingredient = validate_ingredient(ingredient)

    output_file.write_text(
        ingredient.model_dump_json(indent=2),
        encoding="utf-8",
    )

    print(f"Saved to {output_file}")


if __name__ == "__main__":
    main()