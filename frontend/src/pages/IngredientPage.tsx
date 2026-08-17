import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";

import { getIngredient } from "../api/ingredients";
import type { Ingredient, Place } from "../types/ingredients";

import Timeline from "../components/Timeline";
import WorldMap from "../components/WorldMap";
import PlaceInfo from "../components/PlaceInfo";

export default function IngredientPage() {
  const { slug } = useParams();
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);
  const [ingredient, setIngredient] = useState<Ingredient | null>(null);

  const [currentYear, setCurrentYear] = useState<number | null>(null);
  const [selectedPlace, setSelectedPlace] = useState<Place | null>(null);

  // Load ingredient
  useEffect(() => {
    if (!slug) return;

    setLoading(true);
    setError(null);
    setIngredient(null);

    getIngredient(slug)
      .then(setIngredient)
      .catch(() => {
        setError("Failed to load ingredient.");
      })
      .finally(() => {
        setLoading(false);
      });
  }, [slug]);

  // Set initial timeline year
  useEffect(() => {
    if (!ingredient) return;

    const earliestYear = Math.min(
      ...ingredient.places.map(place => place.startYear)
    );

    setCurrentYear(earliestYear);
  }, [ingredient]);

  if (loading) {
    return (
      <main className="flex min-h-screen items-center justify-center bg-slate-950 text-white">
        <p className="text-slate-400">Loading ingredient...</p>
      </main>
    );
  }

  if (error) {
    return (
      <main className="flex min-h-screen items-center justify-center bg-slate-950 text-white">
        <p className="text-red-400">{error}</p>
      </main>
    );
  }

  if (!ingredient || currentYear === null) {
    return null;
  }

  if (ingredient.places.length === 0) {
    return (
      <main className="flex min-h-screen items-center justify-center bg-slate-950 text-white">
        <div className="text-center">
          <h1 className="text-2xl font-semibold">
            {ingredient.name}
          </h1>

          <p className="mt-2 text-slate-400">
            No historical locations are available yet.
          </p>
        </div>
      </main>
    );
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
    <main className="min-h-screen bg-slate-950 text-white">
      <div className="mx-auto max-w-7xl px-6 py-8">

        {/* Ingredient header */}
        <header className="mb-8">
          <h1 className="text-4xl font-bold">
            {ingredient.name}
          </h1>

          <p className="mt-3 max-w-3xl text-slate-300">
            {ingredient.description}
          </p>
        </header>

        {/* Main visualization */}
        <section className="grid gap-6 lg:grid-cols-[1fr_320px]">

          {/* Map */}
          <div className="rounded-xl bg-slate-900 p-4">
            <WorldMap
              places={visiblePlaces}
              onPlaceClick={setSelectedPlace}
            />
          </div>

          {/* Selected place */}
          <aside className="rounded-xl bg-slate-900 p-6">
            <PlaceInfo place={selectedPlace} />
          </aside>

        </section>

        {/* Timeline */}
        <section className="mt-6 rounded-xl bg-slate-900 p-6">
          <Timeline
            minYear={minYear}
            maxYear={maxYear}
            currentYear={currentYear}
            onYearChange={setCurrentYear}
            markers={markers}
          />
        </section>
      </div>
    </main>
  );
}