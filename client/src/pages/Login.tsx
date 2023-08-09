import React, { useState, useEffect} from "react";
import styled from "styled-components";
import Form from "react-bootstrap/Form";
import { Button } from "react-bootstrap";
import './Login.css';
import axios from "axios";
import { useNavigate } from "react-router-dom";
import {toast} from 'react-toastify';

const LoginLabel = styled.h1`
  position: absolute;
  top: 4rem;
  font-size: 20px;
`;

function Login() {

    const navigation = useNavigate();
    const [LoginEmail, setLoginEmail] = useState('');
    const [LoginPassword, setLoginPassword] = useState('');

        const EmailInputHandler = (e:any) => {
          setLoginEmail(e.currentTarget.value);
        };
        const PasswordInputHandler = (e:any) => {
          setLoginPassword(e.currentTarget.value);
        };


    const LoginPostExecute = () => {
      axios
        .post(
          "https://8ab30ea2-e8d1-4c0a-b748-5ec1e2e858c0.mock.pstmn.io/api/v1/login",
          { username: LoginEmail, password: LoginPassword }
        )
        .then(function (response) {
          if(response.status===200)
          {
          toast("Login Completed", {
            position: "top-right",
            autoClose: 5000,
            hideProgressBar: false,
            closeOnClick: true,
            pauseOnHover: true,
            draggable: true,
            progress: undefined,
            theme: "light",
          });
          sessionStorage.setItem("token", response.data.token);
          navigation('/dashboard');
        }
        else{
          toast("Something went wrong, Please Try Again", {
            position: "top-right",
            autoClose: 5000,
            hideProgressBar: false,
            closeOnClick: true,
            pauseOnHover: true,
            draggable: true,
            progress: undefined,
            theme: "light",
          });
        }
        });
    };

    useEffect(() => {
     if (sessionStorage.getItem("token")) {
       toast("Already Logged in", {
         position: "top-right",
         autoClose: 5000,
         hideProgressBar: false,
         closeOnClick: true,
         pauseOnHover: true,
         draggable: true,
         progress: undefined,
         theme: "light",
       });
       navigation("/dashboard");
     }
    }, []);

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
              <Button variant="light" onClick={()=> {navigation('/welcome')}}>Cancel</Button>
            </Form.Group>
          </Form>
        </div>
      </>
    );
}

export default Login;
