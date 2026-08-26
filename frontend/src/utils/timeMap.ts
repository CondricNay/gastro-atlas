interface TimePeriod {
  startYear: number;
  endYear: number;
}

const timeMap: Record<string, TimePeriod> = {
  // Tomato

  "Before European contact": {
    startYear: -500,
    endYear: 1492,
  },

  "Early 16th century": {
    startYear: 1500,
    endYear: 1530,
  },

  "Mid-1500s": {
    startYear: 1540,
    endYear: 1560,
  },

  "Following decades": {
    startYear: 1560,
    endYear: 1600,
  },

  "Early European cultivation": {
    startYear: 1500,
    endYear: 1700,
  },

  "Early 1700s": {
    startYear: 1700,
    endYear: 1730,
  },

  "18th century": {
    startYear: 1700,
    endYear: 1799,
  },

  "Early 1900s": {
    startYear: 1900,
    endYear: 1930,
  },

  "1893": {
    startYear: 1893,
    endYear: 1893,
  },

  // Tea

  "Around 200 BCE": {
    startYear: -200,
    endYear: -200,
  },

  "Around 700": {
    startYear: 700,
    endYear: 700,
  },

  "Around 800": {
    startYear: 800,
    endYear: 800,
  },

  "Around 1600": {
    startYear: 1600,
    endYear: 1600,
  },

  "Around 1800": {
    startYear: 1800,
    endYear: 1800,
  },

  "1870": {
    startYear: 1870,
    endYear: 1870,
  },

  // Chili Pepper

  "Around 6000 BCE": {
    startYear: -6000,
    endYear: -6000,
  },

  "Around 2000 BCE": {
    startYear: -2000,
    endYear: -2000,
  },

  "1493": {
    startYear: 1493,
    endYear: 1493,
  },

  "Around 1500": {
    startYear: 1500,
    endYear: 1500,
  },

  // Coffee

  "Before recorded cultivation": {
    startYear: -1000,
    endYear: 0,
  },

  "Around 1400": {
    startYear: 1400,
    endYear: 1400,
  },
};

export function getTimeRange(time: string): TimePeriod {
  const result = timeMap[time];

  if (!result) {
    throw new Error(`Unknown historical time: ${time}`);
  }

  return result;
}