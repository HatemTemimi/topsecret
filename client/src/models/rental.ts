export interface Rental {
  id: string; // MongoDB ObjectID as a string
  name: string; // Rental name
  fullAddress: string; // Full address string
  streetNumber: string; // Street number
  street: string; // Street name
  city: string; // City name
  country: string; // Country name
  lat: string; // Latitude as a string
  lng: string; // Longitude as a string
  images: string[]; // URLs of rental images
  agree: boolean; // Agreement status
  status: boolean; // Active/Inactive status
  description: string; // Rental description
  price: number; // Rental price
  bedrooms: number; // Number of bedrooms
  bathrooms: number; // Number of bathrooms
  areaSize: number; // Size in square meters
  available: boolean; // Rental availability status
  tags: string[]; // Associated tags
  createdAt: string; // ISO string for creation time
  updatedAt: string; // ISO string for last update time
  createdBy: string; // User ID of the creator
  updatedBy: string; // User ID of the last updater
}