import Editor from "../components/yorkie-tldraw/Editor";
import styled from 'styled-components';
import { useState } from "react";
import AddInstancePrompt from "./AddInstancePrompt";
import { Button } from "react-bootstrap";
import {useEffect} from 'react';
import {useNavigate} from 'react-router-dom';
import {toast} from 'react-toastify';

const Wrapper = styled.div`
    background-color: #F8F9FA;
    width: 100%;
    height: 100%;
`;

function Drawer() {

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
        <Wrapper>
          
          <Editor />
        </Wrapper>
  
      </>
    );
}

export default Drawer;
