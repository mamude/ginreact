import React from 'react'
import { LayoutProps } from '../../common/interfaces/props'
import { PageContainer } from './style'

const PageWrapper: React.FC<LayoutProps> = props => {
  return <PageContainer>{props.children}</PageContainer>
}

export default PageWrapper
