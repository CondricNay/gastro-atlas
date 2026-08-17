export function formatHistoricalYear(year: number): string {
  if (year < 0) {
    return `${Math.abs(year)} BCE`;
  }

  return `${year} CE`;
}