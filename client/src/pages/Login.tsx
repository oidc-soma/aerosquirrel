import React, { useState } from "react";
import styled from "styled-components";
import Form from "react-bootstrap/Form";
import { Button } from "react-bootstrap";
import './Login.css';
import axios from "axios";

const LoginLabel = styled.h1`
  position: absolute;
  top: 4rem;
  font-size: 20px;
`;

function Login() {

    const [LoginEmail, setLoginEmail] = useState('');
    const [LoginPassword, setLoginPassword] = useState('');

        const EmailInputHandler = (e:any) => {
          setLoginEmail(e.currentTarget.value);
        };
        const PasswordInputHandler = (e:any) => {
          setLoginPassword(e.currentTarget.value);
        };


    const LoginPostExecute = () => {
      axios.post("https://d9c25fa3-a939-4ec2-abd9-a479b24bdf39.mock.pstmn.io/api/v1/login", {username: LoginEmail, password: LoginPassword})
      .then(function(response) {
          sessionStorage.setItem("token", response.data.token);
      });
    };

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
              <Form.Control type="text" placeholder="Username" onChange={EmailInputHandler}></Form.Control>
              <br />
              <Form.Control
                type="password"
                placeholder="Password"
                onChange={PasswordInputHandler}
              ></Form.Control>
              <br />
              <br />
              <Button onClick={LoginPostExecute}>Login</Button>
              <Button variant="light">Cancel</Button>
            </Form.Group>
          </Form>
        </div>
      </>
    );
}

export default Login;
