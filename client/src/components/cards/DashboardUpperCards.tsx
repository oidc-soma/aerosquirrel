import React from "react";
import Card from "react-bootstrap/Card";
import { ReactComponent as AWS } from "../../AWS.svg";
import { ReactComponent as Oracle } from "../../Oracle.svg";
import usersicon from '../../users-icon.png';
import globeicon from '../../globe-icon.png';
import resourcesicon from '../../resources-icon.png';
import billsicon from '../../bills-icon.png';

type CardsProps = {
  HeaderTitle: string;
  SecondaryTitle: string;
  Content: string;
};

function DashboardUpperCards({ HeaderTitle, SecondaryTitle, Content }: CardsProps) {
    if(Content==='Users')
    {
        
        return (
          <Card border="normal">
            <Card.Header>
              <img src={usersicon} width={20} height={20} />
              {HeaderTitle}
            </Card.Header>

            <Card.Body>
              <Card.Title>{SecondaryTitle}</Card.Title>
            </Card.Body>
          </Card>
        );
    }
    else if(Content==='Regions')
    {
        return (
          <Card border="normal">
            <Card.Header>
              <img src={globeicon} width={20} height={20} />
              {HeaderTitle}
            </Card.Header>

            <Card.Body>
              <Card.Title>{SecondaryTitle}</Card.Title>
            </Card.Body>
          </Card>
        );
    }
    else if(Content==='Resources')
    {
         return (
           <Card border="normal">
             <Card.Header>
               <img src={resourcesicon} width={20} height={20} />
               {HeaderTitle}
             </Card.Header>

             <Card.Body>
               <Card.Title>{SecondaryTitle}</Card.Title>
             </Card.Body>
           </Card>
         );
    }
    else if(Content==='Bills')
    {
        return (
          <Card border="normal">
            <Card.Header>
              <img src={billsicon} width={20} height={20} />
              {HeaderTitle}
            </Card.Header>

            <Card.Body>
              <Card.Title>{SecondaryTitle}</Card.Title>
            </Card.Body>
          </Card>
        );
    }
    else{
    return (
      <Card border="normal">
        <Card.Header>{HeaderTitle}</Card.Header>
        <Card.Body>
          <Card.Title>{SecondaryTitle}</Card.Title>
          <Card.Text>{"Other"}</Card.Text>
        </Card.Body>
      </Card>
    );
    }
  
}

export default DashboardUpperCards;