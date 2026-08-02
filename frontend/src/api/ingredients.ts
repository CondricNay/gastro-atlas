import type { Ingredient } from "../types/ingredients";


const API_URL = "http://localhost:8080";

export async function getIngredient(slug: string): Promise<Ingredient> {

  const response = await fetch(
    `${API_URL}/ingredients/${slug}`
  );

  if (!response.ok) {
    throw new Error("Failed to fetch ingredient");
  }

  return response.json();
}