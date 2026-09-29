import type { HistoricalEvent } from "../types/ingredients";
import { formatHistoricalYear } from "../utils/formatYear";

interface EventInfoProps {
  event: HistoricalEvent | null;
}

export default function EventInfo({ event }: EventInfoProps) {
  if (!event) {
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
          {event.location}
        </h2>

        <p className="mt-1 text-sm text-slate-400">
          Historical event
        </p>
      </div>

      {(event.startYear !== null ||
        event.endYear !== null) && (
        <div>
          <p className="text-sm font-medium text-slate-400">
            <b>Historical period:</b>{" "}
            {event.startYear !== null
              ? formatHistoricalYear(event.startYear)
              : "Unknown"}{" "}
            -{" "}
            {event.endYear !== null
              ? formatHistoricalYear(event.endYear)
              : "Present"}
          </p>
        </div>
      )}

      <div>
        <p className="text-sm leading-relaxed text-slate-300">
          {event.description}
        </p>
      </div>
    </section>
  );
}