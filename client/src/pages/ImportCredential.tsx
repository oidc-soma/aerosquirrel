import React, {useState} from "react";
import styled from "styled-components";
import CloudProvider from "../components/cards/CloudProvider";
import axios from "axios";
import { Button } from "react-bootstrap";
import {ToastsContainer, ToastsStore, ToastsContainerPosition} from 'react-toasts';
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

interface ChildProps {
  closePopup: () => void;
}

const ConnectLabel = styled.h1`
  font-size: 20px;
`;

const AddFileInstLabel = styled.div`
    color: red;
    font-size: 12px;
`

function ImportCredential({ closePopup }: ChildProps) {
  
    const NofileToast = () => {
        ToastsStore.error("Please Select .toml file to upload.");
    }

    const [FileExist, setFileExist] = useState(null);

    const handleFileMount = (e:any) => {
        setFileExist(e.target.files[0]);
    }

    const handleUpload = () => {
        if(FileExist) {
            const formData = new FormData();
            formData.append('file', FileExist);

            axios
              .post(
                "http://{{url}}/api/v1/resources/import?csp=oci",
                formData,
                {
                  headers: {
                    "Content-Type": "multipart/form-data",
                  },
                }
              )
              .then((res) => {
                if(res.status===200)
                {
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
                }
                else {
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
                console.log(res.data)
              })
              .catch((error)=>{
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
              })
        }
        else {
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
    }
  
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
          {!FileExist && <AddFileInstLabel>Please Select file to Upload</AddFileInstLabel>}
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
