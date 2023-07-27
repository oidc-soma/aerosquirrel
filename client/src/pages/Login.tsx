import React from "react";
import styled from "styled-components";
import Form from "react-bootstrap/Form";
import { Button } from "react-bootstrap";
import './Login.css';

const LoginLabel = styled.h1`
  position: absolute;
  top: 4rem;
  font-size: 20px;
`;

function Login() {

    return (
      <>
        <div className="loginformouter">
          <LoginLabel>Login</LoginLabel>
          <br />
          <br />
          <br />
          <br />
          <br />
          <Form>
            <Form.Group controlId="signupForm.ControlInput">
              <Form.Control type="text" placeholder="Username"></Form.Control>
              <br />
              <Form.Control
                type="password"
                placeholder="Password"
              ></Form.Control>
              <br />
              <br />
              <Button>Login</Button>
              <Button variant="light">Cancel</Button>
            </Form.Group>
          </Form>
        </div>
      </>
    );
}

export default Login;