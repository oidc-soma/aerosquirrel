import React from 'react';
import Cards from '../components/cards/Cards';
import './Dashboard.css';
import styled from 'styled-components';
import ReactDOM from 'react-dom';
import DashboardUpperCards from '../components/cards/DashboardUpperCards';

const DashboardLabel = styled.h1`
  position: absolute;
  left: 8rem;
  top: 1rem;

  font-size: 20px;
`;


function Dashboard() {
    const DocumentTitle: HTMLTitleElement | null = document.querySelector("title");

    if(!DocumentTitle)
    {
        throw new Error('No document title error');
    }
    DocumentTitle.innerText = "Dashboard - Aerosquirrel";

    return (
      <>
        <DashboardLabel>Dashboard</DashboardLabel>
        <div
          className="dashboardGrid"
          style={{
            marginLeft: "140px",
            paddingTop: "110px",
            paddingRight: "5rem",
          }}
        >
          <DashboardUpperCards
            key={1}
            HeaderTitle="Users"
            SecondaryTitle="0"
            Content="Users"
          />
          <DashboardUpperCards
            key={1}
            HeaderTitle="Regions"
            SecondaryTitle="0"
            Content="Regions"
          />
          <DashboardUpperCards
            key={1}
            HeaderTitle="Resources"
            SecondaryTitle="0"
            Content="Resources"
          />{" "}
          <DashboardUpperCards
            key={1}
            HeaderTitle="Bills"
            SecondaryTitle="0"
            Content="Bills"
          />
        </div>
        <div
          className="mapchartGrid"
          style={{
            marginLeft: "140px",
            paddingTop: "30px",
            paddingRight: "5rem",
          }}
        >
          
        </div>
        <div
          className="costexplorerGrid"
          style={{
            marginLeft: "140px",
            paddingTop: "30px",
            paddingRight: "5rem",
          }}
        >
          
        </div>
      </>
    );
}

export default Dashboard;