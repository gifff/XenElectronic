import React, { useEffect, useState } from 'react';

import Container from '@material-ui/core/Container';
import Divider from '@material-ui/core/Divider';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import ListSubheader from '@material-ui/core/ListSubheader';
import Typography from '@material-ui/core/Typography';
import ResponsiveDrawer from './components/ResponsiveDrawer';
import { useCookies } from 'react-cookie';
import ProductList from './pages/ProductList';

import HomeIcon from '@material-ui/icons/Home';
import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';
import client from './lib/client';
import Category from './lib/model/Category';

function App() {
  const [cookies, setCookie] = useCookies(['cartId']);
  const [categories, setCategories] = useState<Category[]>([]);
  const [selectedCategory, setSelectedCategory] = useState<number>(0);

  useEffect(() => {
    if (cookies.cartId === null || cookies.cartId === undefined || cookies.cartId === '') {
      client.createCart()
        .then(response => response.json())
        .then(response => {
          setCookie('cartId', response['cart_id'], { path: '/' });
        });
    }

    if (categories.length === 0) {
      client.listCategories()
        .then(response => response.json())
        .then(response => {
          if (Array.isArray(response)) {
            const newCategories: Category[] = [];
            response.forEach(category => newCategories.push(new Category(category.id, category.name)));
            setCategories(newCategories);
          }
        });
    }
  });

  const drawer = (
    <React.Fragment>
      <List>
        <ListItem button key={"Home"}>
          <ListItemIcon><HomeIcon /></ListItemIcon>
          <ListItemText primary={"Home"} />
        </ListItem>
        <ListItem button key={"Cart"}>
          <ListItemIcon><ShoppingCartIcon /></ListItemIcon>
          <ListItemText primary={"Cart"} />
        </ListItem>
        <Divider />
      </List>
      <List subheader={
        <ListSubheader component="div">
          Categories
        </ListSubheader>
      }>
        {categories.map(category => (
          <ListItem
            button
            key={category.name}
            selected={category.id === selectedCategory}
            onClick={() => { setSelectedCategory(category.id) }}>
            <ListItemText primary={category.name} />
          </ListItem>
        ))}
      </List>
    </React.Fragment>
  );

  return (
    <Container maxWidth="lg">
      <ResponsiveDrawer title="XenElectronic" drawer={drawer}>
        {
          selectedCategory > 0 ? (
            <ProductList categoryId={selectedCategory} />
          ) : (
              <Typography variant='body1'>Select category first</Typography>
            )
        }
      </ResponsiveDrawer>
    </Container>
  );
}

export default App;
