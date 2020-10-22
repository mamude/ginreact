import React from 'react'
import { Route, Switch } from 'react-router-dom'
import LoginPage from './LoginPage'
import HomePage from './Home'

const Navigation: React.FC = () => {
  return (
    <Switch>
      <Route exact path="/" component={LoginPage} />
      <Route exact path="/home" component={HomePage} />
    </Switch>
  )
}

export default Navigation
