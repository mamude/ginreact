import React from 'react';
import { Route, Switch } from 'react-router-dom';
import LoginPage from './LoginPage';


const Navigation: React.FC = () => {
  return <Switch>
    <Route exact path="/" component={LoginPage} />
  </Switch>;
}

export default Navigation;
