import React from 'react'
import { Route, Switch } from 'react-router-dom'
import HomePage from '../../container/HomePage'
import LoginPage from '../../container/LoginPage'

const Navigation: React.FC = () => {
  return (
    <Switch>
      <Route exact path="/" component={HomePage} />
      <Route exact path="/login" component={LoginPage} />
    </Switch>
  )
}

export default Navigation
