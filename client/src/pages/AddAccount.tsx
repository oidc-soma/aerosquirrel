import React from 'react';
import styled from 'styled-components';
import CloudProvider from '../components/cards/CloudProvider';
import './AddAccount.css';
import AWSIcon from '../aws-icon.png';
import OracleIcon from '../oracle-icon.png';
import OnPremiseIcon from '../on-premise.png';
import {Button} from 'react-bootstrap';
interface ChildProps {
    closePopup: () => void;
}

const ConnectLabel = styled.h1`
 
  font-size: 20px;
`;

function AddAccount({closePopup}:ChildProps){

    const handleClick = () => {
        closePopup();
    }

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
                  <input type="text" className="ADDARINPUT"></input>
                </div>
              </div>
              {/* <span className="badge rounded-pill badge-success">Continue</span> */}
            </li>
            <li className="list-group-item d-flex justify-content-between align-items-center">
              <div className="d-flex align-items-center">
                <img src={AWSIcon} width={40}></img>
                <div className="ms-3">
                  <p className="fw-bold mb-1">{"Amazon Web Services(AWS)"}</p>
                  <input type="text" className="ADDARINPUT"></input>
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
                  <input type="text" className="ADDARINPUT"></input>
                </div>
              </div>
              {/* <span className="badge rounded-pill badge-success">Continue</span> */}
            </li>
            <li className="list-group-item d-flex justify-content-between align-items-center">
              <div className="d-flex align-items-center">
                <img width={40}></img>
                <div className="ms-3">
                  <p className="fw-bold mb-1">{"Yorkie"}</p>
                  <input type="text" className="ADDARINPUT"></input>
                </div>
              </div>
              {/* <span className="badge rounded-pill badge-success">Continue</span> */}
            </li>
          </ul>

          <div>{"On Premise Icon Designed by phatplus"}</div>
          <Button>OK</Button>
          <Button>Cancel</Button>
        </div>
      </>
    );
}

export default AddAccount;
