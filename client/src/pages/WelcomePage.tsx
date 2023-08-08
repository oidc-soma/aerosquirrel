import React from "react";
import { styled } from "styled-components";
import { Button } from "react-bootstrap";

function WelcomePage() {

    const AeroTitle = styled.h1`
    font-size: 23px;
    font-weight: 500;
    margin-top: 4rem;
    margin-bottom: 5rem;
    `

    const AdverTitle = styled.h1`
    font-size: 35px;
    font-weight: 600;
    margin-bottom: 5rem;
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
        <AeroTitle>AeroSquirrel</AeroTitle>
        <AdverTitle>
          Configure and Manage <br />
          Cloud Instance at one glance
        </AdverTitle>
        <p>
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
        <Button style={{ marginRight: "12px" }}>Login</Button>

        <Button className="btn btn-secondary">Sign Up</Button>
        <br />
        <br />
        <br />
        <br />
        <AdverImage
          src={process.env.PUBLIC_URL + "/welcomeadvimg.png"}
        ></AdverImage>
        <br />
        <br />
        <DownDiv><DownTitle>Meet the <Amp>New Cloud</Amp><br />with AEROSQUIRREL</DownTitle></DownDiv>
      </>
    );   
}

export default WelcomePage;