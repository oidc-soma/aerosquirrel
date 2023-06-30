import React from 'react';
import styled from 'styled-components';
import CloudProvider from '../components/cards/CloudProvider';

const ConnectLabel = styled.h1`
 
  font-size: 20px;
`;

function AddAccount(){


    return(
        <>
            <ConnectLabel>Add the cloud Account and Connect to Aerosquirrel</ConnectLabel>
            <CloudProvider />
            <div>
                {"On Premise Icon Designed by phatplus"}
                </div>
        </>
    );
}

export default AddAccount;