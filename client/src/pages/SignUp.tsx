import React from 'react';
import styled from 'styled-components';
import Form from 'react-bootstrap/Form';
import { Button } from 'react-bootstrap';
import './SignUp.css';

const SignUpLabel = styled.h1`
  position: absolute;
  left: 10rem;
  top: 4rem;
  font-size: 20px;
`;

function SignUp() {

    return (
      <>
        <SignUpLabel>Sign Up</SignUpLabel>

        <div className="signupformouter">
          <br />
          <br />
          <br />
          <Form>
            <Form.Group controlId="signupForm.ControlInput">
              <Form.Control type="text" placeholder="Username"></Form.Control>
              <br />
              <Form.Control type="email" placeholder="Email"></Form.Control>
              <br />
              <Form.Control
                type="password"
                placeholder="Password"
              ></Form.Control>
              <br />
              <Form.Control
                type="password"
                placeholder="Password Check"
              ></Form.Control>
              <br />
              <div className="buttonouterclass">
                <div className="signupbutton">
                  <Button variant="primary">Sign Up</Button>
                </div>
                <div className="cancelbutton">
                  <Button variant="light">Cancel</Button>
                </div>
              </div>
            </Form.Group>
          </Form>
        </div>
      </>
    );
}

export default SignUp;