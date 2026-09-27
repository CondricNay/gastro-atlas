# import os

# from dotenv import load_dotenv
# from geopy.geocoders import GeoNames


# load_dotenv()

# geolocator = GeoNames(
#     username=os.environ["GEONAMES_USERNAME"],
#     timeout=10,
# )


# def resolve_location(
#     location: str,
# ) -> tuple[float, float] | None:
#     result = geolocator.geocode(
#         location,
#         exactly_one=True,
#     )

#     if result is None:
#         return None

#     return result.latitude, result.longitude


import os

from dotenv import load_dotenv
from geopy.geocoders import GeoNames


load_dotenv()

geolocator = GeoNames(
    username=os.environ["GEONAMES_USERNAME"],
    timeout=10,
)


def resolve_location(
    location: str,
) -> tuple[float, float] | None:
    results = geolocator.geocode(
        location,
        exactly_one=False,
    )

    if not results:
        return None

    location_lower = location.lower().strip()

    # Prefer an exact name match.
    exact_matches = [
        result
        for result in results
        if result.raw.get("name", "").lower().strip() == location_lower
        or result.raw.get("toponymName", "").lower().strip() == location_lower
    ]

    if exact_matches:
        results = exact_matches

    # Prefer populated places and administrative areas.
    # P = populated place
    # A = administrative feature
    results.sort(
        key=lambda result: (
            result.raw.get("fcl") not in ("P", "A"),
            -int(result.raw.get("population", 0)),
        )
    )

    result = results[0]

    return result.latitude, result.longitude
