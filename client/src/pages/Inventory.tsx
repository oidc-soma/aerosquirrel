import React, { useEffect, useState } from 'react';
import styled from 'styled-components';
import InventoryLists from './InventoryLists';
import { useRecoilState, useRecoilValue, useResetRecoilState, useSetRecoilState } from 'recoil';
import { InventoryAtom } from '../atoms';
import axios from 'axios';
import {toast} from 'react-toastify';
import {useNavigate} from 'react-router-dom';

const InventoryLabel = styled.h1`
  position: absolute;
  left: 8rem;
  top: 1rem;

  font-size: 20px;
`;

function Inventory() {

    const [InventoryLSTData, setInventoryLSTData] = useState({});

    const navigation = useNavigate();
       useEffect(() => {
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
         }
       }, []);


    useEffect(() => {
        const LoginToken = sessionStorage.getItem('token');
        axios
          .get(
            "https://8ab30ea2-e8d1-4c0a-b748-5ec1e2e858c0.mock.pstmn.io/api/v1/resources",
            { headers: { Authorization: `Bearer ${LoginToken}` } }
          )
          .then(function (response) {
            console.log(response.data);
            setInventoryLSTData(response.data);
            SetInventoryAtom(response.data);
          });

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

    return (
      <>
        <div className="InventoryWrapper">
          <InventoryLabel>Inventory</InventoryLabel>
          <InventoryLists data={inventoryValue} />

          {/* <p>{JSON.stringify(inventoryValue)}</p> */}
        </div>
      </>
    );

}

export default Inventory;
