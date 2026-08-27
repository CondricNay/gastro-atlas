import { describe, expect, it } from "vitest";
import { getTimeRange } from "./timeMap";


describe("getTimeRange", () => {
  it("always returns a start year before or equal to the end year", () => {
    const periods = [
      "Before European contact",
      "Early 16th century",
      "Mid-1500s",
      "18th century",
      "1893",
      "Around 200 BCE",
      "Around 700",
    ];

    for (const period of periods) {
      const { startYear, endYear } = getTimeRange(period);

      expect(startYear).toBeLessThanOrEqual(endYear);
    }
  });
});