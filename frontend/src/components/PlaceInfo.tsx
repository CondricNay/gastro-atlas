import type { Place } from "../types/ingredients";
import { formatHistoricalYear } from "../utils/formatYear";

interface PlaceInfoProps {
  place: Place | null;
}

export default function PlaceInfo({ place }: PlaceInfoProps) {
  if (!place) {
    return null;
  }

  return (
    <section>
      <h2>{place.name}</h2>

      <p>{place.relationship}</p>

      <p>{formatHistoricalYear(place.startYear)}</p>

      {place.endYear !== null && (
        <p>to{" "}{formatHistoricalYear(place.endYear)}</p>
      )}

      {place.notes && <p>{place.notes}</p>}
    </section>
  );
}