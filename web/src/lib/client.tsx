import config from "./config";

class Client {
  private baseUrl: string;
  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  createCart(): Promise<Response> {
    return fetch(`${this.baseUrl}/carts`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/xenelectronic.v1+json'
      }
    })
  }

  listCategories(): Promise<Response> {
    return fetch(`${this.baseUrl}/categories`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/xenelectronic.v1+json'
      }
    })
  }

  listProductByCategory(categoryId: number): Promise<Response> {
    return fetch(`${this.baseUrl}/categories/${categoryId}/products`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/xenelectronic.v1+json'
      }
    })
  }

  addProductToCart(cartId: string, productId: number): Promise<Response> {
    return fetch(`${this.baseUrl}/carts/${cartId}/products`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/xenelectronic.v1+json'
      },
      body: JSON.stringify({ product_id: productId }),
    })
  }

  listProductsInCart(cartId: string): Promise<Response> {
    return fetch(`${this.baseUrl}/carts/${cartId}/products`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/xenelectronic.v1+json'
      }
    })
  }
}

export default new Client(config.API_BASE_URL);
