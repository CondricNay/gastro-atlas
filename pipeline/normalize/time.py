import re


def resolve_time(
    time_period: str | None,
) -> tuple[int | None, int | None]:
    if not time_period:
        return None, None

    text = time_period.lower().strip()

    # Exact year: "1654", "in 1654", "around 1652"
    year_match = re.search(r"\b(\d{3,4})\b", text)
    if year_match:
        year = int(year_match.group(1))
        return year, year

    # Century expressions:
    # "17th century"
    # "17th and 18th centuries"
    century_matches = re.findall(
        r"\b(\d+)(?:st|nd|rd|th)(?=\s+(?:centur(?:y|ies)|and))",
        text,
    )

    if not century_matches:
        return None, None

    centuries = [int(value) for value in century_matches]

    # Multiple centuries
    if len(centuries) > 1:
        start = (min(centuries) - 1) * 100 + 1
        end = max(centuries) * 100
        return start, end

    century = centuries[0]

    start = (century - 1) * 100 + 1
    end = century * 100

    # "By the 16th century"
    if re.search(r"\bby\b", text):
        return None, end

    # Split the 100-year century into three approximately equal parts.
    # Example: 17th century = 1601–1700
    # Early  = 1601–1633
    # Middle = 1634–1667
    # Late   = 1668–1700

    if re.search(r"\b(early|beginning)\b", text):
        return start, start + 32

    if re.search(r"\b(mid|middle)\b", text):
        return start + 33, start + 66

    if re.search(r"\blate\b", text):
        return start + 67, end

    return start, end
