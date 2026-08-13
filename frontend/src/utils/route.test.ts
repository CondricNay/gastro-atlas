import { describe, expect, it } from "vitest";
import { generateRouteSegments } from "../utils/route";
import type { Place } from "../types/ingredients";

const andes: Place = {
  id: 1,
  name: "Andes Region",
  type: "region",
  latitude: -13.5,
  longitude: -72,
  relationship: "origin",
  startYear: -500,
  endYear: 1500,
  notes: "",
};

const mesoamerica: Place = {
  id: 2,
  name: "Mesoamerica",
  type: "region",
  latitude: 17,
  longitude: -92,
  relationship: "cultivation",
  startYear: -500,
  endYear: 1500,
  notes: "",
};

const spain: Place = {
  id: 3,
  name: "Spanish Empire",
  type: "empire",
  latitude: 40.4168,
  longitude: -3.7038,
  relationship: "introduced",
  startYear: 1500,
  endYear: null,
  notes: "",
};

const italy: Place = {
  id: 4,
  name: "Italy",
  type: "country",
  latitude: 41.9028,
  longitude: 12.4964,
  relationship: "popularized",
  startYear: 1700,
  endYear: null,
  notes: "",
};

describe("generateRouteSegments", () => {
  it("connects consecutive historical stages", () => {
    const places = [
      andes,
      mesoamerica,
      spain,
      italy,
    ];

    const result = generateRouteSegments(places);

    expect(result).toEqual([
      { from: andes, to: spain },
      { from: mesoamerica, to: spain },
      { from: spain, to: italy },
    ]);
  });


  // same year should not connect together
  it("does not connect places within the same historical stage", () => {
    const places = [
      andes,
      mesoamerica,
    ];

    const result = generateRouteSegments(places);

    expect(result).toEqual([]);
  });
});