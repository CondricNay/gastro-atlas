import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";

import { getIngredient } from "../api/ingredients";
import type { Ingredient } from "../types/ingredients";

import Timeline2 from "../components/Timeline";

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

  // const visiblePlaces = ingredient.places.filter(place =>
  //   place.startYear <= currentYear &&
  //   (place.endYear === null || currentYear <= place.endYear)
  // );
  const visiblePlaces = ingredient.places.filter(
    place => place.startYear <= currentYear
  );

  return (
    <main>
      <Timeline2
        minYear={minYear}
        maxYear={maxYear}
        currentYear={currentYear}
        onYearChange={setCurrentYear}
        markers={markers}
      />
      <h2>
        {currentYear < 0
          ? `${Math.abs(currentYear)} BCE`
          : `${currentYear} CE`}
      </h2>

      <h1>{ingredient.name}</h1>
      <p>{ingredient.description}</p>

      <h2>
        Historical Journey
      </h2>

      <ul>
        {
          visiblePlaces.map(place => (
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