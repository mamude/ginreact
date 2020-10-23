import React from 'react'
import { Route, Switch } from 'react-router-dom'
import HomePage from '../../container/HomePage/index'

const Navigation: React.FC = () => {
  return (
    <Switch>
      <Route exact path="/" component={HomePage} />
    </Switch>
  )
}

export default Navigation
