import React from 'react'
import {
  Container,
  Grid,
  Header,
  Icon,
  Image,
  Rating,
  Segment,
} from 'semantic-ui-react'
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

const ShopLogo = styled(Image)`
  width: 150px !important;
`

const ShopName = styled.div`
  font-size: 1.1rem;
  font-weight: bolder;
`

const ShopDetail = styled.div`
  margin-top: 5px;
  font-size: 0.8rem;
  color: #b0adad;
`
const RatingPointuation = styled.strong`
  color: orange;
`

function App() {
  return (
    <Container>
      <Grid columns={2}>
        <Grid.Row color="red">
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
      <Grid columns={2} container stackable>
        <Grid.Row>
          <Grid.Column>
            <Segment>
              <Grid>
                <Grid.Column width={4}>
                  <ShopLogo src={shop} />
                </Grid.Column>
                <Grid.Column width={12}>
                  <ShopName>Loja</ShopName>
                  <ShopDetail>
                    <Rating icon="star" defaultRating={5} />
                    <RatingPointuation>4.5</RatingPointuation>
                    <span> - Categoria - 1.4 km</span>
                  </ShopDetail>
                  <ShopDetail>30-40 min - R$ 6.50</ShopDetail>
                </Grid.Column>
              </Grid>
            </Segment>
          </Grid.Column>
          <Grid.Column>
            <Segment>
              <Grid>
                <Grid.Column width={4}>
                  <ShopLogo src={shop} />
                </Grid.Column>
                <Grid.Column width={12}>
                  <ShopName>Loja</ShopName>
                  <ShopDetail>
                    <Rating icon="star" defaultRating={5} />
                    <RatingPointuation>4.5</RatingPointuation>
                    <span> - Categoria - 1.4 km</span>
                  </ShopDetail>
                  <ShopDetail>30-40 min - R$ 6.50</ShopDetail>
                </Grid.Column>
              </Grid>
            </Segment>
          </Grid.Column>
        </Grid.Row>
        <Grid.Row>
          <Grid.Column>
            <Segment>
              <Grid>
                <Grid.Column width={4}>
                  <ShopLogo src={shop} />
                </Grid.Column>
                <Grid.Column width={12}>
                  <ShopName>Loja</ShopName>
                  <ShopDetail>
                    <Rating icon="star" defaultRating={5} />
                    <RatingPointuation>4.5</RatingPointuation>
                    <span> - Categoria - 1.4 km</span>
                  </ShopDetail>
                  <ShopDetail>30-40 min - R$ 6.50</ShopDetail>
                </Grid.Column>
              </Grid>
            </Segment>
          </Grid.Column>
          <Grid.Column>
            <Segment>
              <Grid>
                <Grid.Column width={4}>
                  <ShopLogo src={shop} />
                </Grid.Column>
                <Grid.Column width={12}>
                  <ShopName>Loja</ShopName>
                  <ShopDetail>
                    <Rating icon="star" defaultRating={5} />
                    <RatingPointuation>4.5</RatingPointuation>
                    <span> - Categoria - 1.4 km</span>
                  </ShopDetail>
                  <ShopDetail>30-40 min - R$ 6.50</ShopDetail>
                </Grid.Column>
              </Grid>
            </Segment>
          </Grid.Column>
        </Grid.Row>
        <Grid.Row>
          <Grid.Column>
            <Segment>
              <Grid>
                <Grid.Column width={4}>
                  <ShopLogo src={shop} />
                </Grid.Column>
                <Grid.Column width={12}>
                  <ShopName>Loja</ShopName>
                  <ShopDetail>
                    <Rating icon="star" defaultRating={5} />
                    <RatingPointuation>4.5</RatingPointuation>
                    <span> - Categoria - 1.4 km</span>
                  </ShopDetail>
                  <ShopDetail>30-40 min - R$ 6.50</ShopDetail>
                </Grid.Column>
              </Grid>
            </Segment>
          </Grid.Column>
          <Grid.Column>
            <Segment>
              <Grid>
                <Grid.Column width={4}>
                  <ShopLogo src={shop} />
                </Grid.Column>
                <Grid.Column width={12}>
                  <ShopName>Loja</ShopName>
                  <ShopDetail>
                    <Rating icon="star" defaultRating={5} />
                    <RatingPointuation>4.5</RatingPointuation>
                    <span> - Categoria - 1.4 km</span>
                  </ShopDetail>
                  <ShopDetail>30-40 min - R$ 6.50</ShopDetail>
                </Grid.Column>
              </Grid>
            </Segment>
          </Grid.Column>
        </Grid.Row>
      </Grid>
    </Container>
  )
}

export default App
