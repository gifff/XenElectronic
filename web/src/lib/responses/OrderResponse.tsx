import CartItemResponse from './CartItemResponse';

export default interface OrderResponse {
  id: string | undefined;
  customer_name: string | undefined;
  customer_email: string | undefined;
  customer_address: string | undefined;
  payment_method: string | undefined;
  payment_account_number: string | undefined;
  payment_amount: number | undefined;
  cart_items: CartItemResponse[] | undefined;
}
