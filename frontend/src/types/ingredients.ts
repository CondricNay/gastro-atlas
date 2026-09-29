export interface HistoricalEvent {
  id: number;
  ingredientId: number;
  description: string;
  timePeriod: string;
  location: string;
  startYear: number | null;
  endYear: number | null;
  latitude: number | null;
  longitude: number | null;
}

export interface Ingredient {
  id: number;
  slug: string;
  name: string;
  description: string;
  events: HistoricalEvent[];
}