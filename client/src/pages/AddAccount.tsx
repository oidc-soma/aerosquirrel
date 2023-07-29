import React from 'react';
import styled from 'styled-components';
import CloudProvider from '../components/cards/CloudProvider';
import './AddAccount.css';

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

          <CloudProvider />
          <div>{"On Premise Icon Designed by phatplus"}</div>
        </div>
      </>
    );
}

export default AddAccount;
