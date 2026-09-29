from typing import Literal

from pydantic import BaseModel, Field


class HistoricalEvent(BaseModel):
    description: str
    time_period: str | None = Field(default=None)
    location: str
    # sources: list[str] = Field(default_factory=list)
    # confidence: Literal["high", "medium", "low"]

class ProcessedHistoricalEvent(BaseModel):
    description: str
    time_period: str | None = Field(default=None)
    location: str
    start_year: int | None = None
    end_year: int | None = None
    latitude: float | None = None
    longitude: float | None = None

class Ingredient(BaseModel):
    name: str
    slug: str
    description: str
    events: list[HistoricalEvent]
