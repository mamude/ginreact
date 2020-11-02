export interface Category {
  id: number
  name: string
}

export interface Market {
  id: number
  name: string
  rating: number
  deliveryTax: number
  deliveryTime: string
  distance: number
  categoryBusiness: Category
}

export interface Product {
  id: number
  name: string
  description: string
  price: number
}
