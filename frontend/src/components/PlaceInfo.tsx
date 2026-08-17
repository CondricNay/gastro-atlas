import type { Place } from "../types/ingredients";
import { formatHistoricalYear } from "../utils/formatYear";

interface PlaceInfoProps {
  place: Place | null;
}

export default function PlaceInfo({ place }: PlaceInfoProps) {
  if (!place) {
    return (
      <div className="flex h-full items-center justify-center text-center text-slate-400">
        <p>Select a location on the map to explore its history.</p>
      </div>
    );
  }

  return (
    <section className="space-y-4">
      <div>
        <h2 className="text-2xl font-semibold">
          {place.name}
        </h2>

        <p className="mt-1 text-sm text-slate-400">
          {place.type}
        </p>
      </div>

      <div>
        <p className="text-sm font-medium text-slate-400">
          <b>Historical period:</b> 
          {" "}{formatHistoricalYear(place.startYear)} - 
          {" "}{place.endYear ? formatHistoricalYear(place.endYear) : "Present"}
        </p>
      </div>

      {place.relationship && (
        <div>
          <p className="text-sm font-medium text-slate-400">
            <b>Relationship:</b> {place.relationship}
          </p>
        </div>
      )}

      {place.notes && (
        <div>
          <p className="text-sm font-medium text-slate-400">
            <b>Notes:</b> {place.notes}
          </p>
        </div>
      )}
    </section>
  );
}