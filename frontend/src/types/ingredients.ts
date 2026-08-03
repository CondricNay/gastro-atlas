export interface Place {
  id: number;
  name: string;
  type: string;
  latitude: number;
  longitude: number;
  relationship: string;
  startYear: number;
  endYear: number | null;
  notes: string;
}

export interface Ingredient {
  id: number;
  slug: string;
  name: string;
  description: string;
  places: Place[];
}