export interface Rental {
  id: string; 
  name: string; 
  address: Address; // Address details
  geometry: Geometry; // Latitude and Longitude
  images: string[]; // URLs of rental images
  agreeToTerms: boolean; // Agreement to terms
  status: "agreed" | "declined" | "pending"; // Rental status
  description: string; // Rental description
  price: number; // Rental price
  currency: "TND" | "USD" | "EUR"; // Currency
  bedrooms: number; // Number of bedrooms
  bathrooms: number; // Number of bathrooms
  areaSize: number; // Size in square meters
  available: boolean; // Rental availability status
  availableFrom: string; // ISO string for availability start date
  tags: string[]; // Associated tags
  type: "shared" | "independent" | "sale"; // Rental type
  standing: "economy" | "standard" | "luxury"; // Rental standing
  amenities: Amenities; // Rental amenities
  rules: Rules; // Rental rules
  createdAt: string; // ISO string for creation time
  updatedAt: string; // ISO string for last update time
  createdBy: string; // User ID of the creator
  updatedBy: string; // User ID of the last updater
  deletedAt?: string; // ISO string for deletion time (soft delete)
  lastUpdatedBy: string; // User ID of the last updater (audit logging)
}

// Nested Address structure
export interface Address {
  streetNumber: string; // Street number
  street: string; // Street name
  city: string; // City name
  country: string; // Country name
  fullAddress: string; // Full address string
}

// Nested Geometry structure
export interface Geometry {
  lat: string; // Latitude as a string
  lng: string; // Longitude as a string
}

// Nested Amenities structure
export interface Amenities {
  airConditioning: boolean; // Air Conditioning availability
  heating: boolean; // Heating availability
  refrigerator: boolean; // Refrigerator availability
  parking: boolean; // Parking availability
}

// Nested Rules structure
export interface Rules {
  petsAllowed: boolean; // Pets allowed status
  partiesAllowed: boolean; // Parties allowed status
  smokingAllowed: boolean; // Smoking allowed status
}
