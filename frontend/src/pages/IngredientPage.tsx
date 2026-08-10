import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";

import { getIngredient } from "../api/ingredients";
import type { Ingredient, Place } from "../types/ingredients";

import Timeline from "../components/Timeline";
import WorldMap from "../components/WorldMap";

export default function IngredientPage() {
  const { slug } = useParams();
  const [ingredient, setIngredient] = useState<Ingredient | null>(null);
  const minYear = -500;
  const maxYear = new Date().getFullYear();
  const [currentYear, setCurrentYear] = useState(-500);
  const [selectedPlace, setSelectedPlace] = useState<Place | null>(null);

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
      <WorldMap
        places={visiblePlaces}
        onPlaceClick={setSelectedPlace}
      />
      <Timeline
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

      {selectedPlace && (
        <div>
          <h2>{selectedPlace.name}</h2>
          <p>{selectedPlace.relationship}</p>
          <p>
            {selectedPlace.startYear < 0
              ? `${Math.abs(selectedPlace.startYear)} BCE`
              : `${selectedPlace.startYear} CE`}
          </p>
          <p>{selectedPlace.notes}</p>
        </div>
      )}

      {/* <ul>
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
      </ul> */}
    </main>
  );
}