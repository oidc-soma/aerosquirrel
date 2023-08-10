import React from "react";
import Card from "react-bootstrap/Card";

type CardsProps = {
  HeaderTitle: string;
  SecondaryTitle: string;
  Content: string;
};

function AccountCards({ HeaderTitle, SecondaryTitle, Content }: CardsProps) {
  return (
    <Card border="info">
      <Card.Header>{HeaderTitle}</Card.Header>
      <Card.Body>
        <Card.Title>{SecondaryTitle}</Card.Title>
        <Card.Text>{Content}</Card.Text>
      </Card.Body>
    </Card>
  );
}

export default AccountCards;
