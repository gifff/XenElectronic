import Product from './Product';

export default class CartItem {
  id: number;
  productId: number;
  product: Product;

  constructor(id: number, productId: number, product: Product) {
    this.id = id;
    this.productId = productId;
    this.product = product;
  }
}