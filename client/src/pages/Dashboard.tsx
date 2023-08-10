// @ts-ignore
// @ts-nocheck

import React from 'react';
import Cards from '../components/cards/Cards';
import './Dashboard.css';
import styled from 'styled-components';
import ReactDOM from 'react-dom';
import {useNavigate} from 'react-router-dom';
import DashboardUpperCards from '../components/cards/DashboardUpperCards';
import { PieChart, Pie, Legend, Tooltip, ResponsiveContainer, CartesianGrid } from 'recharts';
import {BarChart, Bar, Cell, XAxis, YAxis } from 'recharts';
import Editor from '../components/yorkie-tldraw/Editor';
import { useEffect } from 'react';
import axios from 'axios';
import {
  useRecoilState,
  useRecoilValue,
  useResetRecoilState,
  useSetRecoilState,
} from "recoil";
import { CSPAtom, InventoryAtom } from '../atoms';
import { toast } from "react-toastify";
import { result } from 'cypress/types/lodash';


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

function countUniqueIds(objectsArray:{}[]) {
  const uniqueIds = new Set();

  objectsArray.forEach((obj) => {
    if (obj.hasOwnProperty("user_id")) {
      uniqueIds.add(obj.user_id);
    }
  });

  return uniqueIds.size;
}

function Dashboard() {
    const DocumentTitle: HTMLTitleElement | null = document.querySelector("title");
    const [GetInventoryAtom, SetInventoryAtom] = useRecoilState(InventoryAtom);
    const [GetCSPAtom, SetCSPAtom] = useRecoilState(CSPAtom);
    const navigation = useNavigate();
    useEffect(()=> {
      localStorage.setItem("yorkie", "ciprjcqbjhd3s76qlvg0");
    },[]);
    useEffect(()=> {
      if (!sessionStorage.getItem("token")) {
      toast("Please Login or Signup", {
      position: "top-right",
      autoClose: 5000,
      hideProgressBar: false,
      closeOnClick: true,
      pauseOnHover: true,
      draggable: true,
      progress: undefined,
      theme: "light",
    });
    navigation("/welcome");
    }},[]);


    if(!DocumentTitle)
    {
        throw new Error('No document title error');
    }
    DocumentTitle.innerText = "Dashboard - Aerosquirrel";

    useEffect(() => {
          const LoginToken = sessionStorage.getItem("token");
          axios
            .get(
              "https://8ab30ea2-e8d1-4c0a-b748-5ec1e2e858c0.mock.pstmn.io/api/v1/resources",
              { headers: { Authorization: `Bearer ${LoginToken}` } }
            )
            .then(function (response) {
              SetInventoryAtom(response.data);
            });
            
          axios.get(
            "https://8ab30ea2-e8d1-4c0a-b748-5ec1e2e858c0.mock.pstmn.io/api/v1/configs/csps",
            { headers: { Authorization: `Bearer ${LoginToken}` } }
          ).then(function (response) {
            //console.log(response.data);
            SetCSPAtom(response.data);
          })

          

          
          //SetInventoryAtom(InventoryLSTData);
    }, []);


    const typeCosts = {};

    GetInventoryAtom.resources.forEach((resource) => {
      const resourceType = resource.type;
      const cost = resource.cost;

      if (typeCosts.hasOwnProperty(resourceType)) {
        typeCosts[resourceType] += cost;
      } else {
        typeCosts[resourceType] = cost;
      }
    });

    const resultArray = [];

for (const resourceType in typeCosts) {
  resultArray.push({ name: resourceType, value: typeCosts[resourceType] });
}

    console.log(resultArray);

    console.log(GetInventoryAtom.resources);
    console.log(GetInventoryAtom.resources.length.toString());
          

    let idnum = countUniqueIds(GetInventoryAtom.resources).toString();

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
              SecondaryTitle={idnum}
              Content="Users"
            />
            <DashboardUpperCards
              key={1}
              HeaderTitle="Regions"
              SecondaryTitle={GetInventoryAtom.resources.length.toString()}
              Content="Regions"
            />
            <DashboardUpperCards
              key={1}
              HeaderTitle="Resources"
              SecondaryTitle={GetInventoryAtom.resources.length.toString()}
              Content="Resources"
            />{" "}
            <DashboardUpperCards
              key={1}
              HeaderTitle="Bills"
              SecondaryTitle={GetInventoryAtom.resources.length.toString()}
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
                  data= {resultArray}
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
