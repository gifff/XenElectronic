import React from 'react';

import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import ResponsiveDrawer from './components/ResponsiveDrawer';
import { makeStyles } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
  paper: {
    padding: theme.spacing(2),
    textAlign: 'center',
    color: theme.palette.text.secondary,
  },
}))

function App() {
  const classes = useStyles();

  return (
    <Container maxWidth="lg">
      <ResponsiveDrawer title="XenElectronic">
        <Grid container spacing={3}>
          {
            [...Array(81)].map(() => (
              <Grid item xs={12} sm={6} md={4}>
                <Paper className={classes.paper}>Product</Paper>
              </Grid>
            ))
          }
        </Grid>
      </ResponsiveDrawer>
    </Container>
  );
}

export default App;
