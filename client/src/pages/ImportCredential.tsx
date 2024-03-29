import React, { useState } from "react";
import styled from "styled-components";
import CloudProvider from "../components/cards/CloudProvider";
import axios from "axios";
import { Button } from "react-bootstrap";
import {
  ToastsContainer,
  ToastsStore,
  ToastsContainerPosition,
} from "react-toasts";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { useNavigate } from "react-router-dom";
interface ChildProps {
  closePopup: () => void;
}

const ConnectLabel = styled.h1`
  font-size: 20px;
`;

const AddFileInstLabel = styled.div`
  color: red;
  font-size: 12px;
`;

function ImportCredential({ closePopup }: ChildProps) {
  const NofileToast = () => {
    ToastsStore.error("Please Select .toml file to upload.");
  };

  const [FileExist, setFileExist] = useState(null);
  const [FileContent, setFileContent] = useState("");
  const navigation = useNavigate();

  const handleFileMount = (e: any) => {
    setFileExist(e.target.files[0]);
    const reader = new FileReader();
    reader.onload = (e: any) => {
      const content = e.target.result;
      setFileContent(content);
    };
    let content = reader.readAsText(e.target.files[0]);
    console.log(FileContent);
  };

  const handleUpload = () => {
    if (FileExist) {
      const formData = new FormData();
      formData.append("file", FileExist);
      const LoginToken = sessionStorage.getItem("token");

      const parameters = {
        csp: FileContent,
      };

      axios
        .get(
          "https://8ab30ea2-e8d1-4c0a-b748-5ec1e2e858c0.mock.pstmn.io/api/v1/resources/import",
          {
            params: parameters,
            headers: { Authorization: `Bearer ${LoginToken}` },
          }
        )
        .then((res) => {
          if (res.status === 200) {
            toast("Imported Credential successfully.", {
              position: "top-right",
              autoClose: 5000,
              hideProgressBar: false,
              closeOnClick: true,
              pauseOnHover: true,
              draggable: true,
              progress: undefined,
              theme: "light",
            });
            closePopup();
          } else {
            toast("Something went wrong, please try again.", {
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
          console.log(res.data);
        })
        .catch((error) => {
          toast("Something went wrong, please try again.", {
            position: "top-right",
            autoClose: 5000,
            hideProgressBar: false,
            closeOnClick: true,
            pauseOnHover: true,
            draggable: true,
            progress: undefined,
            theme: "light",
          });
          closePopup();
        });
    } else {
      toast("Please Select File to Upload", {
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
  };

  const handleClick = () => {
    closePopup();
  };

  return (
    <>
      <>
        <div className="AddAccountBackground">
          <ConnectLabel>
            Select and Upload .toml file to import credentials.
          </ConnectLabel>
          <button type="button" className="btn-close" onClick={handleClick}>
            <span className="icon-cross"></span>
            <span className="visually-hidden">Close</span>
          </button>

          <br />
          <br />
          <input type="file" onChange={handleFileMount} />
          <br />
          <br />
          {!FileExist && (
            <AddFileInstLabel>Please Select file to Upload</AddFileInstLabel>
          )}
          <br />
          <br />

          <Button onClick={handleUpload}>Upload Credentials</Button>
          <Button onClick={handleClick} className="btn btn-secondary">
            Close
          </Button>
        </div>
      </>
    </>
  );
}

export default ImportCredential;
