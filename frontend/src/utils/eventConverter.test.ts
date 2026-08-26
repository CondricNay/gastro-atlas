import { describe, expect, it } from "vitest";

import { convertEventToPlace } from "./eventConverter";
import type { HistoricalEvent } from "../types/ingredients";

describe("convertEventToPlace", () => {
  it("converts a historical event into a Place", () => {
    const event: HistoricalEvent = {
      id: 1,
      title: "Tomatoes were introduced to Spain",
      description:
        "Tomato seeds were brought from Mexico to Spain by early Spanish explorers.",
      time: "Early 16th century",
      entity: "Spanish explorers",
      location: "Spain",
      sources: ["uvm-tomato-history"],
      confidence: "high",
    };

    const place = convertEventToPlace(event);

    expect(place).toEqual({
      id: 1,
      name: "Spain",
      type: "historical",
      latitude: 40.4,
      longitude: -3.7,
      relationship: "Spanish explorers",
      startYear: 1500,
      endYear: 1530,
      notes:
        "Tomato seeds were brought from Mexico to Spain by early Spanish explorers.",
    });
  });
});