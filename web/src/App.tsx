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
import { makeStyles } from '@material-ui/core';
import {
  Route,
  Switch,
  useRouteMatch,
  useLocation,
} from 'react-router-dom';

import HomeIcon from '@material-ui/icons/Home';
import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';

import Cart from './pages/Cart';
import ViewOrder from './pages/ViewOrder';
import ProductList from './pages/ProductList';

import ListItemLink from './components/ListItemLink';

import client from './lib/client';
import Category from './lib/model/Category';

const useStyles = makeStyles((theme) => ({
  container: {
    paddingBottom: theme.spacing(8),
    paddingTop: theme.spacing(2),
  },
}))

function App() {
  const categoryRouteMatch = useRouteMatch<{ categoryId: string }>('/categories/:categoryId');
  const location = useLocation();
  const categoryIdFromRoute = categoryRouteMatch !== null ? Number(categoryRouteMatch.params.categoryId) : 0
  const [cookies, setCookie] = useCookies(['cartId']);
  const [categories, setCategories] = useState<Category[]>([]);
  const [selectedCategory, setSelectedCategory] = useState<number>(!isNaN(categoryIdFromRoute) ? categoryIdFromRoute : 0);
  const classes = useStyles();

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
        <ListItemLink
          to="/"
          selected={location.pathname === '/'}
          key={"Home"}>
          <ListItemIcon><HomeIcon /></ListItemIcon>
          <ListItemText primary={"Home"} />
        </ListItemLink>
        <ListItemLink
          to="/cart"
          selected={location.pathname === '/cart'}
          key={"Cart"}>
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
    <Container maxWidth="lg" className={classes.container}>
      <ResponsiveDrawer title="XenElectronic" drawer={drawer}>
        <Switch>
          <Route exact path="/">
            <Typography variant="body1">Select category first</Typography>
          </Route>
          <Route exact path="/cart">
            <Cart />
          </Route>
          <Route path="/orders/:orderId">
            <ViewOrder />
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
  );
}

export default App;
