import type { HistoricalEvent } from "../types/ingredients";

export interface RouteSegment {
  from: HistoricalEvent;
  to: HistoricalEvent;
}

// Events with the same startYear belong to the same historical stage.
export function generateRouteSegments(
  events: HistoricalEvent[]
): RouteSegment[] {
  // Only events with valid years and coordinates
  // can participate in map routes.
  const plottableEvents = events.filter(
    (event) =>
      event.startYear !== null &&
      event.latitude !== null &&
      event.longitude !== null
  );

  const groups = new Map<number, HistoricalEvent[]>();

  for (const event of plottableEvents) {
    const group = groups.get(event.startYear);

    if (group) {
      group.push(event);
    } else {
      groups.set(event.startYear, [event]);
    }
  }

  const entries = [...groups.entries()];

  entries.sort((a, b) => a[0] - b[0]);

  const sortedGroups = entries.map(
    (entry) => entry[1]
  );

  const segments: RouteSegment[] = [];

  for (let i = 0; i < sortedGroups.length - 1; i++) {
    const currentGroup = sortedGroups[i];
    const nextGroup = sortedGroups[i + 1];

    // Connect every event in one historical stage
    // to every event in the next chronological stage.
    for (const from of currentGroup) {
      for (const to of nextGroup) {
        segments.push({from, to});
      }
    }
  }

  return segments;
}