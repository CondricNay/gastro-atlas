import type { Place } from "../types/ingredients";

export interface RouteSegment {
  from: Place;
  to: Place;
}

// Places with the same startYear belong to the same historical stage.
export function generateRouteSegments(places: Place[]): RouteSegment[] {
  const groups = new Map<number, Place[]>();

  for (const place of places) {
    const group = groups.get(place.startYear);

    if (group) {
      group.push(place);
    } else {
      groups.set(place.startYear, [place]);
    }
  }

  const entries = [...groups.entries()];
  entries.sort((a, b) => a[0] - b[0]);
  const sortedGroups = entries.map(entry => entry[1]);

  const segments: RouteSegment[] = [];

  for (let i = 0; i < sortedGroups.length - 1; i++) {
    const currentGroup = sortedGroups[i];
    const nextGroup = sortedGroups[i + 1];

    // Connect every place in one historical stage
    // to every place in the next chronological stage.
    for (const from of currentGroup) {
      for (const to of nextGroup) {
        segments.push({ from, to });
      }
    }
  }

  return segments;
}