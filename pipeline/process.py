import json

from models import HistoricalEvent
from normalize.time import resolve_time
from normalize.location import resolve_location


INPUT_PATH = "data/extracted/coffee.json"
OUTPUT_PATH = "data/processed/coffee.json"


def main():
    with open(INPUT_PATH, "r", encoding="utf-8") as file:
        data = json.load(file)

    events = [HistoricalEvent.model_validate(item) for item in data]

    processed = []

    for event in events:
        start_year, end_year = resolve_time(event.time_period)
        coordinates = resolve_location(event.location)

        processed.append({
            "description": event.description,
            "time_period": event.time_period,
            "location": event.location,
            "start_year": start_year,
            "end_year": end_year,
            "latitude": coordinates[0] if coordinates else None,
            "longitude": coordinates[1] if coordinates else None,
        })

    with open(OUTPUT_PATH, "w", encoding="utf-8") as file:
        json.dump(processed, file, ensure_ascii=False, indent=2)


if __name__ == "__main__":
    main()
