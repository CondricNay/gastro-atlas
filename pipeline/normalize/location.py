import os

from dotenv import load_dotenv
from geopy.geocoders import GeoNames


load_dotenv()

geolocator = GeoNames(
    username=os.environ["GEONAMES_USERNAME"],
    timeout=10,
)


def is_valid_result(result) -> bool:
    feature_class = result.raw.get("fcl")

    if feature_class not in ("A", "P", "T", "L"):
        return False

    return True


def resolve_location(
    location: str,
) -> tuple[float, float] | None:
    results = geolocator.geocode(
        location,
        exactly_one=False,
    )

    if not results:
        return None

    results = [
        result
        for result in results
        if is_valid_result(result)
    ]

    if not results:
        return None

    location_lower = location.lower().strip()

    exact_matches = [
        result
        for result in results
        if result.raw.get("name", "").lower().strip() == location_lower
        or result.raw.get("toponymName", "").lower().strip() == location_lower
    ]

    if exact_matches:
        results = exact_matches

    results.sort(
        key=lambda result: (
            result.raw.get("fcl") not in ("P", "A"),
            -int(result.raw.get("population", 0)),
        )
    )

    result = results[0]

    return result.latitude, result.longitude
