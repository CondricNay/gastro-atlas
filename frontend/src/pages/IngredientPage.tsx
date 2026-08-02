import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";

import { getIngredient } from "../api/ingredients";
import type { Ingredient } from "../types/ingredients";

import Timeline from "../components/Timeline";

export default function IngredientPage() {

  const { slug } = useParams();

  const [ingredient, setIngredient] =
    useState<Ingredient | null>(null);

  useEffect(() => {

    if (!slug) return;

    getIngredient(slug)
      .then(setIngredient);

  }, [slug]);


  if (!ingredient) {
    return <div>Loading...</div>;
  }


  return (
    <main>

      <Timeline places={ingredient.places} />
      <h1>
        {ingredient.name}
      </h1>

      <p>
        {ingredient.description}
      </p>


      <h2>
        Historical Journey
      </h2>

      <ul>
        {
          ingredient.places.map(place => (
            <li key={place.name}>
              {place.name}
              {" - "}
              {place.relationship}
              {" ("}{place.startYear}{")"}
            </li>
          ))
        }
      </ul>

    </main>
  );
}