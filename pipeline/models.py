from typing import Literal

from pydantic import BaseModel, Field


class HistoricalEvent(BaseModel):
    title: str
    description: str
    time_period: str
    entity: str
    location: str
    sources: list[str] = Field(default_factory=list)
    confidence: Literal["high", "medium", "low"]


class Ingredient(BaseModel):
    name: str
    slug: str
    description: str
    events: list[HistoricalEvent]
