// @ts-ignore
// @ts-nocheck

import React, {useEffect, useState} from 'react';
import styled from 'styled-components';
import CloudProvider from '../components/cards/CloudProvider';
import './AddAccount.css';
import AWSIcon from '../aws-icon.png';
import OracleIcon from '../oracle-icon.png';
import OnPremiseIcon from '../on-premise.png';
import {Button} from 'react-bootstrap';
import {CSPAtom} from '../atoms';
import {
  useRecoilState,
  useRecoilValue,
  useResetRecoilState,
  useSetRecoilState,
} from "recoil";
import axios from 'axios';
import {toast} from 'react-toastify';
interface ChildProps {
    closePopup: () => void;
}

const ConnectLabel = styled.h1`
 
  font-size: 20px;
`;

function AddAccount({closePopup}:ChildProps){

    const [GetCSPAtom, SetCSPAtom] = useRecoilState(CSPAtom);

    const [AWSvalue, setAWSvalue] = useState<string|undefined>('');
    const [OCIvalue, setOCIvalue] = useState<string|undefined>('');
    const [ONPvalue, setONPvalue] = useState<string|undefined>('');
    const [YORKIEvalue, setYORKIEvalue] = useState<string>('');

    const handleClick = () => {
        closePopup();
    }

    const handleYorkieChange = (event:any) => {
      setYORKIEvalue(event.target.value);
    }

    const handleOCIChange = (event:any) => {
      setOCIvalue(event.target.value);
    }
    
    const handleONPChange = (event:any) => {
      setONPvalue(event.target.value);
    }

    const handleAWSChange = (event:any) => {
      setAWSvalue(event.target.value);
    }

    const AccountModify = () => {
 
      const CSPtempr = {
        aws: [{ name: GetCSPAtom.aws[0].name }, { token: AWSvalue }],
        oci: [{ name: GetCSPAtom.oci[0].name }, { token: OCIvalue }],
        k8s: [{ name: GetCSPAtom.k8s[0].name }, { token: ONPvalue }],
      }; 
      SetCSPAtom(CSPtempr);
      const yorkietemp = YORKIEvalue;
      localStorage.setItem('yorkie', yorkietemp);
      const LoginToken = sessionStorage.getItem('token');
      
      axios.patch(
        'https://8ab30ea2-e8d1-4c0a-b748-5ec1e2e858c0.mock.pstmn.io/api/v1/configs'
        ,CSPtempr
        ,
        { headers: { Authorization: `Bearer ${LoginToken}` }}
      ).then(
        (response) => {
          if(response.status===200)
          {
             toast("Cloud Account Changed successfully.", {
               position: "top-right",
               autoClose: 5000,
               hideProgressBar: false,
               closeOnClick: true,
               pauseOnHover: true,
               draggable: true,
               progress: undefined,
               theme: "light",
             });
            closePopup();
          }
          else{
             toast("Something went wrong, please try again.", {
               position: "top-right",
               autoClose: 5000,
               hideProgressBar: false,
               closeOnClick: true,
               pauseOnHover: true,
               draggable: true,
               progress: undefined,
               theme: "light",
             });
             closePopup();
          }
          
        }
      )

    }

    useEffect(() => {
      let LoginToken = sessionStorage.getItem('token');
      axios
        .get(
          "https://8ab30ea2-e8d1-4c0a-b748-5ec1e2e858c0.mock.pstmn.io/api/v1/configs/csps",
          { headers: { Authorization: `Bearer ${LoginToken}` } }
        )
        .then(function (response) {
          //console.log(response.data);
          SetCSPAtom(response.data);
        });
        setAWSvalue(GetCSPAtom.aws[1].token);
        setOCIvalue(GetCSPAtom.oci[1].token);
        setONPvalue(GetCSPAtom.k8s[1].token);
        const yk = localStorage.getItem('yorkie');
        if(yk)
        {
          setYORKIEvalue(yk);
        }
    }, []);

    return (
      <>
        <div className="AddAccountBackground">
          <ConnectLabel>
            Add the cloud Account and Connect to Aerosquirrel
          </ConnectLabel>

          <button type="button" className="btn-close" onClick={handleClick}>
            <span className="icon-cross"></span>
            <span className="visually-hidden">Close</span>
          </button>

          <ul className="list-group list-group-light cloudprovider">
            <li className="list-group-item d-flex justify-content-between align-items-center">
              <div className="d-flex align-items-center">
                <img src={OracleIcon} width={40}></img>
                <div className="ms-3">
                  <p className="fw-bold mb-1">
                    {"Oracle Cloud Infrastructure(OCI)"}
                  </p>
                  <input type="text" className="ADDARINPUT" value={OCIvalue} onChange={handleOCIChange}></input>
                </div>
              </div>
              {/* <span className="badge rounded-pill badge-success">Continue</span> */}
            </li>
            <li className="list-group-item d-flex justify-content-between align-items-center">
              <div className="d-flex align-items-center">
                <img src={AWSIcon} width={40}></img>
                <div className="ms-3">
                  <p className="fw-bold mb-1">{"Amazon Web Services(AWS)"}</p>
                  <input type="text" className="ADDARINPUT" value={AWSvalue} onChange={handleAWSChange}></input>
                </div>
              </div>
              <br />

              {/* <span className="badge rounded-pill badge-success">Continue</span> */}
            </li>
            <li className="list-group-item d-flex justify-content-between align-items-center">
              <div className="d-flex align-items-center">
                <img src={OnPremiseIcon} width={40}></img>
                <div className="ms-3">
                  <p className="fw-bold mb-1">{"On-Premise"}</p>
                  <input type="text" className="ADDARINPUT" value={ONPvalue} onChange={handleONPChange}></input>
                </div>
              </div>
              {/* <span className="badge rounded-pill badge-success">Continue</span> */}
            </li>
            <li className="list-group-item d-flex justify-content-between align-items-center">
              <div className="d-flex align-items-center">
                <img width={40}></img>
                <div className="ms-3">
                  <p className="fw-bold mb-1">{"Yorkie"}</p>
                  <input type="text" className="ADDARINPUT" value={YORKIEvalue} onChange={handleYorkieChange}></input>
                </div>
              </div>
              {/* <span className="badge rounded-pill badge-success">Continue</span> */}
            </li>
          </ul>

          <div>{"On Premise Icon Designed by phatplus"}</div>
          <Button onClick={AccountModify}>OK</Button>
          <Button>Cancel</Button>
        </div>
      </>
    );
}

export default AddAccount;
