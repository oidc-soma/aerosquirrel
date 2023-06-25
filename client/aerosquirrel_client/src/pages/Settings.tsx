import React from 'react';
import styled from 'styled-components';

const SettingsLabel = styled.h1`
  position: absolute;
  left: 8rem;
  top: 1rem;

  font-size: 20px;
`;

const PrefLabel = styled.h1`
  position: absolute;
  left: 8rem;
  top: 5rem;
  font-size: 15px;
`

const UserNameLabel = styled.h1`
  position: absolute;
  left: 8rem;
  top: 10rem;
  font-size: 15px;
`;

const UserNameContentLabel = styled.h1`
  position: absolute;
  left: 8rem;
  top: 12rem;
  font-size: 25px;
`;

const Hr = styled.hr`
    position: absolute;
    left: 0px;
    top: 20rem;
    width: 2px;
`;

const UserBillLine = styled.div`
    position: absolute;
    left: 300px;
    top: 20rem;
    background-color: black;
`

function Settings() {

      const DocumentTitle: HTMLTitleElement | null =
        document.querySelector("title");

      if (!DocumentTitle) {
        throw new Error("No document title error");
      }
      DocumentTitle.innerText = "Settings - Aerosquirrel";


    return(
        <>
        <SettingsLabel>Settings</SettingsLabel>
        <PrefLabel>Profile</PrefLabel>
        <UserNameLabel>Username </UserNameLabel>
        <UserNameContentLabel>TestUSER</UserNameContentLabel>
        <UserBillLine>___</UserBillLine>
        
        <hr />
        </>
    )
}

export default Settings;