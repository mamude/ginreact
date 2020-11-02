import React from 'react'
import { LayoutProps } from '../../common/interfaces/props'
import { Title } from './style'

const PageTitle: React.FC<LayoutProps> = props => {
  return <Title variant="h6">{props.children}</Title>
}

export default PageTitle
