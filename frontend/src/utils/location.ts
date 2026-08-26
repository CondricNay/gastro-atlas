export interface Location {
  latitude: number;
  longitude: number;
}

const locations: Record<string, Location> = {
  // Tomato

  "Andes, South America": {
    latitude: -13.5,
    longitude: -72.0,
  },

  "Central America": {
    latitude: 15.0,
    longitude: -90.0,
  },

  Mexico: {
    latitude: 23.6,
    longitude: -102.5,
  },

  Spain: {
    latitude: 40.4,
    longitude: -3.7,
  },

  Italy: {
    latitude: 41.9,
    longitude: 12.5,
  },

  Europe: {
    latitude: 54.5,
    longitude: 15.0,
  },

  Americas: {
    latitude: 15.0,
    longitude: -75.0,
  },

  "Monticello, Virginia": {
    latitude: 37.9,
    longitude: -78.5,
  },

  "United States": {
    latitude: 39.8,
    longitude: -98.6,
  },

  // Tea

  China: {
    latitude: 35.8617,
    longitude: 104.1954,
  },

  Korea: {
    latitude: 35.9078,
    longitude: 127.7669,
  },

  Japan: {
    latitude: 36.2048,
    longitude: 138.2529,
  },

  "United Kingdom": {
    latitude: 55.3781,
    longitude: -3.4360,
  },

  India: {
    latitude: 20.5937,
    longitude: 78.9629,
  },

  "Sri Lanka": {
    latitude: 7.8731,
    longitude: 80.7718,
  },

  // Chili Pepper

  Mesoamerica: {
    latitude: 17.0,
    longitude: -92.0,
  },

  "Southeast Asia": {
    latitude: 5.0,
    longitude: 110.0,
  },

  // Coffee

  Ethiopia: {
    latitude: 9.145,
    longitude: 40.4897,
  },

  Yemen: {
    latitude: 15.5527,
    longitude: 48.5164,
  },

  "Arabian Peninsula": {
    latitude: 23.8859,
    longitude: 45.0792,
  },

  "Ottoman Empire": {
    latitude: 39.0,
    longitude: 35.0,
  },
};

export function getLocationCoordinates(location: string): Location {
  const result = locations[location];

  if (!result) {
    throw new Error(`Unknown location: ${location}`);
  }

  return result;
}