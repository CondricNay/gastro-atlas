import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";

import { getIngredient } from "../api/ingredients";
import type { Ingredient, Place } from "../types/ingredients";

import Timeline from "../components/Timeline";
import WorldMap from "../components/WorldMap";
import { formatHistoricalYear } from "../utils/formatYear";
import PlaceInfo from "../components/PlaceInfo";

export default function IngredientPage() {
  const { slug } = useParams();
  const [ingredient, setIngredient] = useState<Ingredient | null>(null);

  const [currentYear, setCurrentYear] = useState<number | null>(null);
  const [selectedPlace, setSelectedPlace] = useState<Place | null>(null);

  // Load ingredient
  useEffect(() => {
    if (!slug) return;

    getIngredient(slug).then(setIngredient);
  }, [slug]);

  // Set initial timeline year
  useEffect(() => {
    if (!ingredient) return;

    const earliestYear = Math.min(
      ...ingredient.places.map(place => place.startYear)
    );

    setCurrentYear(earliestYear);
  }, [ingredient]);

  if (!ingredient || currentYear === null) {
    return <div>Loading...</div>;
  }

  const minYear = Math.min(
    ...ingredient.places.map(place => place.startYear)
  );

  const maxYear = new Date().getFullYear();

  const visiblePlaces = ingredient.places.filter(
    place => place.startYear <= currentYear
  );

  const markers = ingredient.places.map(place => ({
    year: place.startYear,
    label: `${place.relationship}: ${place.name}`,
  }));

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
      <h2>{formatHistoricalYear(currentYear)}</h2>

      <h1>{ingredient.name}</h1>
      <p>{ingredient.description}</p>

      <h2>
        Historical Journey
      </h2>

      <PlaceInfo place={selectedPlace} />
    </main>
  );
}