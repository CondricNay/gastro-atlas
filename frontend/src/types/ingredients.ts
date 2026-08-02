export interface Place {
  name: string;
  type: string;
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