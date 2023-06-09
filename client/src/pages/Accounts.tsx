import React, {useState} from "react";
import { Link, Route } from "react-router-dom";
import styled from "styled-components";
import Button from "react-bootstrap/Button";
import 'bootstrap/dist/css/bootstrap.min.css';
import Cards from '../components/cards/Cards';
import './Accounts.css';
import CloudProvider from "../components/cards/CloudProvider";
import AddAccount from "./AddAccount";

const AccountsLabel = styled.h1`
position: absolute;
left: 8rem;
top: 1rem;

    font-size: 20px;
`;

const AddAccountButton = 
"position: absolute; right: 5rem; top: 4rem;";

function Accounts() {
  const [showAddPopup, setshowAddPopup] = useState(false);
  
  const showPopupFunction = () => {
    setshowAddPopup(true);
  }

  const closePopupFunction = () => {
    setshowAddPopup(false);
  };

  const DocumentTitle: HTMLTitleElement | null = document.querySelector("title");

    if (!DocumentTitle) {
      throw new Error("No document title error");
    }
    DocumentTitle.innerText = "Accounts - Aerosquirrel";


  return (
    <>
      <AccountsLabel>Accounts</AccountsLabel>
      <Button
        variant="outline-primary"
        style={{ position: "absolute", right: "5rem", top: "4rem" }}
        onClick={showPopupFunction}>
        Add New Account
      </Button>
      {showAddPopup&&<AddAccount closePopup={closePopupFunction}/>}
      {/* <Cards
        HeaderTitle="Amazon Web Services"
        SecondaryTitle="AWS"
        Content="AWS"
      ></Cards> */}

      <div
        className="cardGrid"
        style={{
          marginLeft: "140px",
          paddingTop: "180px",
          paddingRight: "5rem",
        }}
      >
        <Cards
          key={1}
          HeaderTitle="CustomTitle"
          SecondaryTitle="Amazon Web Services"
          Content="AWS"
        />
        <Cards
          key={2}
          HeaderTitle="CustomTitle2"
          SecondaryTitle="Oracle OCI"
          Content="Oracle"
        />
        {/* {[...Array(5)].map((_, i) => (
          <Cards key={i} HeaderTitle="TEST" SecondaryTitle="TEST" Content="TEST" />
        ))} */}
      </div>
      
    </>
    // <nav className="fixed bottom-0 left-0 top-0 z-20 flex w-[88px] flex-col gap-6 bg-purplin-850 px-5 py-8 dark:bg-black-900 pt-16">
    //     <div className="flex items-center gap-8 text-sm font-semibold text-black-400" rel="noopener noreferrer">
    //        <div>
    //         링크 로고 이미지 삽입예정
    //        </div>
    //     </div>
    // </nav>
  );
}

export default Accounts;
