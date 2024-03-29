import React, { useState, useEffect } from "react";
import { Link, Route } from "react-router-dom";
import styled from "styled-components";
import Button from "react-bootstrap/Button";
import "bootstrap/dist/css/bootstrap.min.css";
import Cards from "../components/cards/Cards";
import "./Accounts.css";
import CloudProvider from "../components/cards/CloudProvider";
import AddAccount from "./AddAccount";
import ImportCredential from "./ImportCredential";
import { toast } from "react-toastify";
import { useNavigate } from "react-router-dom";

const AccountsLabel = styled.h1`
  position: absolute;
  left: 8rem;
  top: 1rem;
  font-size: 20px;
`;

const AddAccountButton = "position: absolute; right: 5rem; top: 4rem;";

function Accounts() {
  const [showAddPopup, setshowAddPopup] = useState(false);
  const [showCredPopup, setshowCredPopup] = useState(false);

  const showPopupFunction = () => {
    setshowAddPopup(true);
  };

  const showCredFunction = () => {
    setshowCredPopup(true);
  };

  const closePopupFunction = () => {
    setshowAddPopup(false);
  };

  const closeCredFunction = () => {
    setshowCredPopup(false);
  };

  const LogoutFunction = () => {
    sessionStorage.removeItem("token");
    navigation("/welcome");
  };

  const DocumentTitle: HTMLTitleElement | null =
    document.querySelector("title");

  if (!DocumentTitle) {
    throw new Error("No document title error");
  }
  DocumentTitle.innerText = "Accounts - Aerosquirrel";

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

  return (
    <>
      <AccountsLabel>Accounts</AccountsLabel>
      {/* <Button
        variant="outline-primary"
        style={{ position: "absolute", right: "5rem", top: "4rem" }}
        onClick={showPopupFunction}
      >
        Add/Modify Account
      </Button> */}
      <button className="ADDACButton" onClick={showPopupFunction}>
        Add/Modify Account
      </button>
      <button className="ImportCRDButton" onClick={showCredFunction}>
        Import Credentials
      </button>
      {showAddPopup && <AddAccount closePopup={closePopupFunction} />}
      {showCredPopup && <ImportCredential closePopup={closeCredFunction} />}

      {/* <Cards
        HeaderTitle="Amazon Web Services"
        SecondaryTitle="AWS"
        Content="AWS"
      ></Cards> */}

      <div className="cardWrapper" style={{ height: "100vh", opacity: 1 }}>
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
            HeaderTitle="Oracle OCI Account"
            SecondaryTitle="Oracle OCI"
            Content="Oracle"
          />
          <Cards
            key={2}
            HeaderTitle="AWS Account"
            SecondaryTitle="Amazon Web Services"
            Content="AWS"
          />
          <Cards
            key={3}
            HeaderTitle="Yorkie Account"
            SecondaryTitle="Yorkie"
            Content=""
          />

          {/* {[...Array(5)].map((_, i) => (
          <Cards key={i} HeaderTitle="TEST" SecondaryTitle="TEST" Content="TEST" />
        ))} */}
        </div>
        <a className="logout" onClick={LogoutFunction}>
          Logout
        </a>
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
