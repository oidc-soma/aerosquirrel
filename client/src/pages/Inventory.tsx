import React, { useEffect, useState } from 'react';
import styled from 'styled-components';
import InventoryLists from './InventoryLists';
import { useRecoilState, useRecoilValue, useResetRecoilState, useSetRecoilState } from 'recoil';
import { InventoryAtom } from '../atoms';
import axios from 'axios';

const InventoryLabel = styled.h1`
  position: absolute;
  left: 8rem;
  top: 1rem;

  font-size: 20px;
`;

function Inventory() {

    const [InventoryLSTData, setInventoryLSTData] = useState({});

    useEffect(() => {
        const LoginToken = sessionStorage.getItem('token');
        axios.get(
          "https://d9c25fa3-a939-4ec2-abd9-a479b24bdf39.mock.pstmn.io/api/v1/resources", {headers: {"Authorization": `Bearer ${LoginToken}`}}
        ).then(function(response){
          console.log(response.data);
          setInventoryLSTData(response.data);
          SetInventoryAtom(response.data);
        })

        //SetInventoryAtom(InventoryLSTData);
    },[]       
    );

    const [GetInventoryAtom, SetInventoryAtom] = useRecoilState(InventoryAtom);
    const inventoryValue = useRecoilValue(InventoryAtom);
    const setinventoryRecoilState = useSetRecoilState(InventoryAtom);
    const resetState = useResetRecoilState(InventoryAtom);

    // setinventoryRecoilState({
    //   'resources': [
    //     {
    //       'id': 1,
    //       'user_id': 1,
    //       'name': "rating api server",
    //       'type': "on-prem",
    //       'cost': 0.5,
    //       'metadata': [{ 'admin': "someone" }, { 'spec': "..." }],
    //       'tags': [
    //         {
    //           'key': "alternative_name",
    //           'value': "rating v1",
    //         },
    //       ],
    //       'link': "ssh://127.0.0.1",
    //     },
    //     {
    //       'id': 1,
    //       'user_id': 1,
    //       'name': "rating api server",
    //       'type': "on-prem",
    //       'cost': 0.5,
    //       'metadata': [{ admin: "someone" }, { spec: "..." }],
    //       'tags': [
    //         {
    //           'key': "alternative_name",
    //           'value': "rating v1",
    //         },
    //       ],
    //       'link': "ssh://127.0.0.1",
    //     },
    //   ],
    // });

    return(
        <>
            <InventoryLabel>Inventory</InventoryLabel>
            <InventoryLists data={inventoryValue} />
            
            {/* <p>{JSON.stringify(inventoryValue)}</p> */}
           
        </>
    );

}

export default Inventory;
