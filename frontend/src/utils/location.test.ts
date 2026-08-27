import { describe, expect, it } from "vitest";
import { getLocationCoordinates } from "./location";

describe("getLocationCoordinates", () => {
  it("returns valid coordinates for a known location", () => {
    const { latitude, longitude } =
      getLocationCoordinates("Mexico");

    expect(latitude).toBeGreaterThanOrEqual(-90);
    expect(latitude).toBeLessThanOrEqual(90);

    expect(longitude).toBeGreaterThanOrEqual(-180);
    expect(longitude).toBeLessThanOrEqual(180);
  });

  it("returns valid coordinates for all supported locations", () => {
    const locations = [
      "China",
      "Korea",
      "Japan",
      "United Kingdom",
      "India",
      "Sri Lanka",
      "Mesoamerica",
      "Mexico",
      "Spain",
      "Southeast Asia",
      "Ethiopia",
      "Yemen",
      "Arabian Peninsula",
      "Ottoman Empire",
      "Europe",
    ];

    for (const location of locations) {
      const { latitude, longitude } =
        getLocationCoordinates(location);

      expect(latitude).toBeGreaterThanOrEqual(-90);
      expect(latitude).toBeLessThanOrEqual(90);

      expect(longitude).toBeGreaterThanOrEqual(-180);
      expect(longitude).toBeLessThanOrEqual(180);
    }
  });
});