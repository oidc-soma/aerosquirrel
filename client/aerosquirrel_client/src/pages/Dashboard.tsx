import React from 'react';
import Cards from '../components/cards/Cards';
import './Dashboard.css';
import styled from 'styled-components';
import ReactDOM from 'react-dom';
import DashboardUpperCards from '../components/cards/DashboardUpperCards';
import { PieChart, Pie, Legend, Tooltip, ResponsiveContainer, CartesianGrid } from 'recharts';
import {BarChart, Bar, Cell, XAxis, YAxis } from 'recharts';

const DashboardLabel = styled.h1`
  position: absolute;
  left: 8rem;
  top: 1rem;
  font-size: 20px;
`;

const data01 = [
  { name: "On Premise", value: 400 },
  { name: "Amazon Web Services", value: 300 },
  { name: "Oracle OCI", value: 300 },
];

const barchartdummy = [
  {
    name: "2023 01",
    EC2: 4000,
    DynomoDB: 2400,
    ElasticIP: 2400,
  },
  {
    name: "2023 02",
    EC2: 3000,
    DynomoDB: 1398,
    ElasticIP: 2210,
  },
  {
    name: "2023 03",
    EC2: 2000,
    DynomoDB: 9800,
    ElasticIP: 2290,
  },
  {
    name: "2023 04",
    EC2: 2780,
    DynomoDB: 3908,
    ElasticIP: 2000,
  },
  {
    name: "2023 05",
    EC2: 1890,
    DynomoDB: 4800,
    ElasticIP: 2181,
  },
  {
    name: "2023 06",
    EC2: 2390,
    DynomoDB: 3800,
    ElasticIP: 2500,
  },
];

function Dashboard() {
    const DocumentTitle: HTMLTitleElement | null = document.querySelector("title");

    if(!DocumentTitle)
    {
        throw new Error('No document title error');
    }
    DocumentTitle.innerText = "Dashboard - Aerosquirrel";

    return (
      <>
        <div className="dashboardscreen">
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
            <div className="mcb-label">Mostly Cost Breakdown</div>

            <ResponsiveContainer width="100%" height="100%">
              <PieChart width={400} height={400}>
                <Pie
                  dataKey="value"
                  isAnimationActive={false}
                  data={data01}
                  cx="50%"
                  cy="50%"
                  innerRadius={40}
                  outerRadius={80}
                  fill="#8884d8"
                  label
                />
                <Tooltip />
              </PieChart>
            </ResponsiveContainer>
          </div>
          <div
            className="costexplorerGrid"
            style={{
              marginLeft: "140px",
              paddingTop: "30px",
              paddingRight: "5rem",
            }}
          >
            <div className="mcb-label">Cost Explorer</div>
            <ResponsiveContainer width="100%" height="100%">
              <BarChart
                width={500}
                height={300}
                data={barchartdummy}
                margin={{
                  top: 5,
                  right: 30,
                  left: 20,
                  bottom: 5,
                }}
              >
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="name" />
                <YAxis />
                <Tooltip />
                <Legend />
                <Bar dataKey="DynomoDB" fill="#8884d8" />
                <Bar dataKey="EC2" fill="#82ca9d" />
                <Bar dataKey="ElasticIP" fill="#000000" />
              </BarChart>
            </ResponsiveContainer>
            </div>
        </div>
      </>
    );
}

export default Dashboard;