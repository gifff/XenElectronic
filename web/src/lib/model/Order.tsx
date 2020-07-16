import CartItem from './CartItem';

export default class Order {
  id: string;
  customerName: string;
  customerEmail: string;
  customerAddress: string;
  paymentMethod: string;
  paymentAccountNumber: string;
  paymentAmount: number;
  cartItems: CartItem[];

  constructor(id: string = '',
    customerName: string = '',
    customerEmail: string = '',
    customerAddress: string = '',
    paymentMethod: string = '',
    paymentAccountNumber: string = '',
    paymentAmount: number = 0,
    cartItems: CartItem[] = new Array<CartItem>()
  ) {
    this.id = id;
    this.customerName = customerName;
    this.customerEmail = customerEmail;
    this.customerAddress = customerAddress;
    this.paymentMethod = paymentMethod;
    this.paymentAccountNumber = paymentAccountNumber;
    this.paymentAmount = paymentAmount;
    this.cartItems = cartItems;
  }
}
