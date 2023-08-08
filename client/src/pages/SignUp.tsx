import React, {useState} from 'react';
import styled from 'styled-components';
import Form from 'react-bootstrap/Form';
import { Button } from 'react-bootstrap';
import './SignUp.css';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const SignUpLabel = styled.h1`
  position: absolute;
  left: 10rem;
  top: 4rem;
  font-size: 20px;
`;

function SignUp() {
  const navigation = useNavigate();
  const [isUsernameexist, setUsernameexist] = useState<boolean>(false);
  const [isUsernameclicked, setUsernameclicked] = useState<boolean>(false);
  const [isEmailexist, setEmailexist] = useState<boolean>(false);
  const [isEmailclicked, setEmailclicked] = useState<boolean>(false);
  const [isPasswordexist, setPasswordexist] = useState<boolean>(false);
  const [isPasswordclicked, setPasswordclicked] = useState<boolean>(false);
  const [isPasswordcheckexist, setPasswordcheckexist] = useState<boolean>(false);
  const [isPasswordcheckclicked, setPasswordcheckclicked] = useState<boolean>(false);
  const [SignupUsername, setSignupUsername] = useState<string>('');
  const [SignupEmail, setSignupEmail] = useState<string>('');
  const [SignupPassword, setSignupPassword] = useState<string>('');
 
  const SignupSubmitFunction = () => {
    axios.post(
      "https://d9c25fa3-a939-4ec2-abd9-a479b24bdf39.mock.pstmn.io/api/v1/users",
      {
        username: SignupUsername,
        email: SignupEmail,
        password: SignupPassword,
      }
    ).then(function(response) {
        sessionStorage.setItem('token',response.data.token);
    });
    
  }

  return (
    <>
      <SignUpLabel>Sign Up</SignUpLabel>
      <div className="signupformouter">
        <br />
        <br />
        <br />
        <Form>
          <Form.Group controlId="signupForm.ControlInput">
            <Form.Control
              type="text"
              placeholder="Username"
              onClick={() => {
                setUsernameclicked(true);
              }}
              onChange={(e) => {
                if (e.target.value !== null && e.target.value !== "") {
                  setUsernameexist(true);
                } else if (e.target.value === null || e.target.value === "") {
                  setUsernameexist(false);
                }
                setSignupUsername(e.currentTarget.value);
              }}
            ></Form.Control>
            {isUsernameclicked && !isUsernameexist ? (
              <div className="Errcaption">Please input username</div>
            ) : null}
            <br />
            <Form.Control
              type="email"
              placeholder="Email"
              onClick={() => {
                setEmailclicked(true);
              }}
              onChange={(e) => {
                if (e.target.value !== null && e.target.value !== "") {
                  setEmailexist(true);
                } else if (e.target.value === null || e.target.value === "") {
                  setEmailexist(false);
                }
              setSignupEmail(e.currentTarget.value);
              }}
            ></Form.Control>
            {isEmailclicked && !isEmailexist ? (
              <div className="Errcaption">Please input Email address</div>
            ) : null}

            <br />
            <Form.Control
              type="password"
              placeholder="Password"
              onClick={() => {
                setPasswordclicked(true);
              }}
              onChange={(e) => {
                if (e.target.value !== null && e.target.value !== "") {
                  setPasswordexist(true);
                } else if (e.target.value === null || e.target.value === "") {
                  setPasswordexist(false);
                }
                setSignupPassword(e.currentTarget.value);
              }}
            ></Form.Control>
            {isPasswordclicked && !isPasswordexist ? (
              <div className="Errcaption">
                Password must be 12 to 20 characters.
              </div>
            ) : null}

            <br />
            <Form.Control
              type="password"
              placeholder="Password Check"
              onClick={() => {
                setPasswordcheckclicked(true);
              }}
              onChange={(e) => {
                if (e.target.value != null && e.target.value !== "") {
                  setPasswordcheckexist(true);
                } else if (e.target.value === null || e.target.value === "") {
                  setPasswordcheckexist(false);
                }
              }}
            ></Form.Control>
            {isPasswordcheckclicked && !isPasswordcheckexist ? (
              <div className="Errcaption">
                Please input the same password with previous input.
              </div>
            ) : null}
            <br />
            <div className="buttonouterclass">
              <div
                className="signupbutton"
                onClick={
                //   () => {
                //   axios
                //     .post(
                //       "https://d9c25fa3-a939-4ec2-abd9-a479b24bdf39.mock.pstmn.io/api/v1/user",
                //       {
                //         username: "oidc",
                //         email: "oidc@gmail.com",
                //         password: "password",
                //       }
                //     )
                //     .then(function (response) {
                //       console.log(response);
                //     });
                // }
                SignupSubmitFunction
              }
              >
                <Button variant="primary">Sign Up</Button>
              </div>
              <div className="cancelbutton">
                <Button variant="light" onClick={()=>{navigation('/welcome')}}>Cancel</Button>
              </div>
            </div>
          </Form.Group>
        </Form>
      </div>
    </>
  );
}

export default SignUp;
