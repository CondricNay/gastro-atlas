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
