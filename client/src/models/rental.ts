export interface Rental {
    id: string;
    name: string;
    fullAddress: string;
    streetNumber: string;
    street: string;
    city: string;
    country: string;
    lat: string;
    lng: string;
    images: string[]; // URLs of rental images
    agree: boolean; // Agreement status
    status: boolean; // Availability status
  }
  