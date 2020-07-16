export default interface CartItemResponse {
  id: number;
  productId: number;
  product: {
    id: number | undefined;
    categoryId: number | undefined;
    name: string | undefined;
    description: string | undefined;
    photo: string | undefined;
    price: number | undefined;
  };
}
