import React from 'react'
import { Container, Grid, Header, Icon, Item } from 'semantic-ui-react'
import styled from 'styled-components'
import logo from './images/logo.png'
import shop from './images/image.png'

const GridColumn = styled(Grid.Column)`
  height: 50px;
`

const Buttons = styled.div`
  margin: 20px;
`

const Logo = styled.img`
  object-fit: cover;
  width: 120px;
  height: 50px;
  margin-top: 5px;
`

function App() {
  return (
    <Container>
      <Grid columns={2}>
        <Grid.Row color="teal">
          <GridColumn textAlign="left">
            <Logo src={logo} alt="Samir Mamude" />
          </GridColumn>
          <GridColumn textAlign="right">
            <Buttons>
              <Icon name="shop" size="large" />
              <Icon name="user" size="large" />
            </Buttons>
          </GridColumn>
        </Grid.Row>
      </Grid>
      <Header as="h2">Lojas</Header>
      <Container>
        <Item.Group divided>
          <Item>
            <Item.Image src={shop} />
            <Item.Content>
              <Item.Header as="a">Content Header</Item.Header>
              <Item.Meta>
                <span>Date</span>
                <span>Category</span>
              </Item.Meta>
              <Item.Description>Lorem ipsum</Item.Description>
            </Item.Content>
          </Item>
          <Item>
            <Item.Image src={shop} />
            <Item.Content>
              <Item.Header as="a">Content Header</Item.Header>
              <Item.Meta>
                <span>Date</span>
                <span>Category</span>
              </Item.Meta>
              <Item.Description>Lorem ipsum</Item.Description>
            </Item.Content>
          </Item>
          <Item>
            <Item.Image src={shop} />
            <Item.Content>
              <Item.Header as="a">Content Header</Item.Header>
              <Item.Meta>
                <span>Date</span>
                <span>Category</span>
              </Item.Meta>
              <Item.Description>Lorem ipsum</Item.Description>
            </Item.Content>
          </Item>
        </Item.Group>
      </Container>
    </Container>
  )
}

export default App
