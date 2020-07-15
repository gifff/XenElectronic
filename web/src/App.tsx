import React, { useEffect, useState } from 'react';

import Container from '@material-ui/core/Container';
import Divider from '@material-ui/core/Divider';
import List from '@material-ui/core/List';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import ListSubheader from '@material-ui/core/ListSubheader';
import Typography from '@material-ui/core/Typography';
import ResponsiveDrawer from './components/ResponsiveDrawer';
import { useCookies } from 'react-cookie';
import {
  BrowserRouter as Router,
  Route,
  Switch,
} from 'react-router-dom';

import HomeIcon from '@material-ui/icons/Home';
import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';

import ProductList from './pages/ProductList';

import ListItemLink from './components/ListItemLink';

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
        <ListItemLink to="/" key={"Home"}>
          <ListItemIcon><HomeIcon /></ListItemIcon>
          <ListItemText primary={"Home"} />
        </ListItemLink>
        <ListItemLink to="/cart" key={"Cart"}>
          <ListItemIcon><ShoppingCartIcon /></ListItemIcon>
          <ListItemText primary={"Cart"} />
        </ListItemLink>
        <Divider />
      </List>
      <List subheader={
        <ListSubheader component="div">
          Categories
        </ListSubheader>
      }>
        {categories.map(category => (
          <ListItemLink
            to={`/categories/${category.id}`}
            key={category.name}
            selected={category.id === selectedCategory}
            onClick={() => { setSelectedCategory(category.id) }}>
            <ListItemText primary={category.name} />
          </ListItemLink>
        ))}
      </List>
    </React.Fragment>
  );

  return (
    <Router>
      <Container maxWidth="lg">
        <ResponsiveDrawer title="XenElectronic" drawer={drawer}>
          <Switch>
            <Route exact path="/">
              <Typography variant="body1">Select category first</Typography>
            </Route>
            <Route path="/categories/:categoryId">
              <ProductList />
            </Route>
            <Route path="*">
              <Typography variant="body1">The page you're looking for is not found</Typography>
            </Route>
          </Switch>
        </ResponsiveDrawer>
      </Container>
    </Router>
  );
}

export default App;
