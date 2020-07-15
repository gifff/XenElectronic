export default class Product {
  id: number;
  categoryId: number;
  name: string;
  description: string;
  photo: string;
  price: number;

  constructor(id: number = 0, categoryId: number = 0, name: string = '', description: string = '', photo: string = '', price: number = 0) {
    this.id = id;
    this.categoryId = categoryId;
    this.name = name;
    this.description = description;
    this.photo = photo;
    this.price = price;
  }
};
