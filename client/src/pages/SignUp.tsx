import React, {useState} from 'react';
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
  const [isUsernameexist, setUsernameexist] = useState<boolean>(false);
  const [isUsernameclicked, setUsernameclicked] = useState<boolean>(false);
  const [isEmailexist, setEmailexist] = useState<boolean>(false);
  const [isEmailclicked, setEmailclicked] = useState<boolean>(false);
  const [isPasswordexist, setPasswordexist] = useState<boolean>(false);
  const [isPasswordclicked, setPasswordclicked] = useState<boolean>(false);
  const [isPasswordcheckexist, setPasswordcheckexist] = useState<boolean>(false);
  const [isPasswordcheckclicked, setPasswordcheckclicked] = useState<boolean>(false);
  
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
            {isUsernameexist ? null : (
              <div className="Errcaption">Please input username</div>
            )}
            <br />
            <Form.Control type="email" placeholder="Email"></Form.Control>
            {isEmailexist ? null : (
              <div className="Errcaption">Please input Email address</div>
            )}

            <br />
            <Form.Control type="password" placeholder="Password"></Form.Control>
            {isPasswordexist ? null : (
              <div className="Errcaption">
                Password must be 12 to 20 characters.
              </div>
            )}

            <br />
            <Form.Control
              type="password"
              placeholder="Password Check"
            ></Form.Control>
            {isPasswordcheckexist ? null : (
              <div className="Errcaption">
                Password must be 12 to 20 characters.
              </div>
            )}
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