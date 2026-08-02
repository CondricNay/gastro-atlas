import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";

import { getIngredient } from "../api/ingredients";
import type { Ingredient } from "../types/ingredients";

// import Timeline from "../components/Timeline";
import Timeline2 from "../components/Timeline2";

export default function IngredientPage() {
  const { slug } = useParams();
  const [ingredient, setIngredient] = useState<Ingredient | null>(null);
  const minYear = -500;
  const maxYear = new Date().getFullYear();
  const [currentYear, setCurrentYear] = useState(-500);

  useEffect(() => {
    if (!slug) return;

    getIngredient(slug).then(setIngredient);
  }, [slug]);

  if (!ingredient) {
    return <div>Loading...</div>;
  }

  const markers = ingredient.places.map(place => ({
    year: place.startYear,
    label: `${place.relationship}: ${place.name}`,
  }));

  return (
    <main>
      {/* <Timeline places={ingredient.places} /> */}
      <Timeline2
        minYear={minYear}
        maxYear={maxYear}
        currentYear={currentYear}
        onYearChange={setCurrentYear}
        markers={markers}
      />
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