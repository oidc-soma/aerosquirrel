import Editor from "../components/yorkie-tldraw/Editor";
import styled from 'styled-components';
import { useState } from "react";
import AddInstancePrompt from "./AddInstancePrompt";
import { Button } from "react-bootstrap";

const Wrapper = styled.div`
    background-color: #F8F9FA;
    width: 100%;
    height: 100%;
`;

function Drawer() {

  

    return (
      <>
        <Wrapper>
          
          <Editor />
        </Wrapper>
  
      </>
    );
}

export default Drawer;