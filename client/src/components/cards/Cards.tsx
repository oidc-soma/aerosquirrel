import React from 'react';
import Card from 'react-bootstrap/Card';
import { ReactComponent as AWS } from '../../AWS.svg';
import { ReactComponent as Oracle} from '../../Oracle.svg';

type CardsProps = {
    HeaderTitle: string,
    SecondaryTitle: string,
    Content: string,
};

function Cards({HeaderTitle, SecondaryTitle, Content}:CardsProps) {


if(Content==="AWS")
{
    return (
      <Card border="info">
        <Card.Header>{HeaderTitle}</Card.Header>
        <Card.Body>
          <Card.Title>
            <AWS></AWS>
          </Card.Title>
          <Card.Text>{SecondaryTitle}</Card.Text>
        </Card.Body>
      </Card>
    );

}else if(Content==="Oracle"){
  return (
    <Card border="info">
      <Card.Header>{HeaderTitle}</Card.Header>
      <Card.Body>
        <Card.Title>
          <Oracle></Oracle>
        </Card.Title>
        <Card.Text>{SecondaryTitle}</Card.Text>
      </Card.Body>
    </Card>
  );


}
else{
return (
  <Card border="info">
    <Card.Header>{HeaderTitle}</Card.Header>
    <Card.Body>
      <Card.Title>{SecondaryTitle}</Card.Title>
      <Card.Text>{"Other"}</Card.Text>
    </Card.Body>
  </Card>
);

}




}

export default Cards;
