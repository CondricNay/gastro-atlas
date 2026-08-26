import type { HistoricalEvent, Place } from "../types/ingredients";
import { getLocationCoordinates } from "./location";
import { getTimeRange } from "./timeMap";

export function convertEventToPlace(
  event: HistoricalEvent
): Place {
  const { latitude, longitude } =
    getLocationCoordinates(event.location);

  const { startYear, endYear } =
    getTimeRange(event.timePeriod);

  return {
    id: event.id,
    name: event.location,
    type: "historical",
    latitude,
    longitude,
    relationship: event.entity,
    startYear,
    endYear,
    notes: event.description,
  };
}