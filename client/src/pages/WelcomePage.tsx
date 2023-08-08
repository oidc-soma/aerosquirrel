import React from "react";
import { styled } from "styled-components";
import { Button } from "react-bootstrap";
import { useNavigate } from "react-router-dom";
import './WelcomePage.css';

function WelcomePage() {

    const navigation = useNavigate();

    const AeroTitle = styled.h1`
    font-size: 23px;
    font-weight: 500;
    margin-top: 4rem;
    margin-bottom: 3rem;
    `
    
    const AdverTitle = styled.h1`
    font-size: 35px;
    font-weight: 600;
    margin-bottom: 3rem;
    `
    const AdverImage = styled.img`
    width: 60%;
    height: 60%;
    border-radius: 15px;
    `
    const DownDiv = styled.div`
    position: relative;
    width: 100%;
    height: 200px;
    background-color: black;
    `

    const DownTitle = styled.h1`
    font-size: 40px;
    color: white;
    position: absolute;
    top: 7vh;;
    left: 36vw;
    `

    const Amp = styled.a`
      color: #00c3ff;
      text-decoration: none;
    `;

    return (
      <>
        <AeroTitle className="WelAni">AeroSquirrel</AeroTitle>
        <AdverTitle className="WelAni">
          Configure and Manage <br />
          Cloud Instance at one glance
        </AdverTitle>
        <p className="WelAni">
          Unlock Infinite Possibilities with Aerosquirrel!
          <br />
          Elevate your convenience to with our solutions.
          <br />
          Like squirrel, Seamlessly manage, store, and scale your data
          <br />
          while experiencing unparalleled reliability and security. <br />
          Join the future of technology and explore limitless innovation today.
        </p>
        <br />
        <br />
        <br />
        <Button className="WelAni"
          onClick={() => {
            navigation("/login");
          }}
          style={{ marginRight: "12px" }}
        >
          Login
        </Button>

        <Button
          className="btn btn-secondary WelAni"
          onClick={() => {
            navigation("/signup");
          }}
        >
          Sign Up
        </Button>
        <br />
        <br />

        <AdverImage
          src={process.env.PUBLIC_URL + "/welcomeadvimg.png"}
        ></AdverImage>
        <br />
        <br />
        <DownDiv className="WelAni">
          <DownTitle>
            Meet the <Amp>New Cloud</Amp>
            <br />
            with AEROSQUIRREL
          </DownTitle>
        </DownDiv>
      </>
    );   
}

export default WelcomePage;