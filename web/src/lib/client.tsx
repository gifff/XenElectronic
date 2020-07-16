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

  removeProductFromCart(cartId: string, productId: number): Promise<Response> {
    return fetch(`${this.baseUrl}/carts/${cartId}/products/${productId}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/xenelectronic.v1+json'
      },
    })
  }

  createOrder(cartId: string, customerName: string, customerEmail: string, customerAddress: string): Promise<Response> {
    return fetch(`${this.baseUrl}/orders`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/xenelectronic.v1+json'
      },
      body: JSON.stringify({
        cart_id: cartId,
        customer_name: customerName,
        customer_email: customerEmail,
        customer_address: customerAddress,
      }),
    })
  }

  viewOrder(orderId: string): Promise<Response> {
    return fetch(`${this.baseUrl}/orders/${orderId}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/xenelectronic.v1+json'
      }
    })
  }
}

export default new Client(config.API_BASE_URL);
