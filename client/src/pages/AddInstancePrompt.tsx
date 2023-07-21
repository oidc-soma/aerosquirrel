import React from "react";
import styled from "styled-components";
import { Button } from "react-bootstrap";
import { doc } from '../hooks/useMultiplayerState';
import { useState } from "react";
import {
  Tldraw,
  useFileSystem,
  TDUserStatus,
  TDAsset,
  TDBinding,
  TDShape,
  TDUser,
  TldrawApp,
} from "@tldraw/tldraw";

export const Select = styled.select`
  -webkit-appearance: none;
  -moz-appearance: none;
  appearance: none;
  margin: 0;
  min-width: 0;
  display: block;
  width: 100%;
  padding: 8px 8px;
  font-size: inherit;
  line-height: inherit;
  border: 1px solid;
  border-radius: 4px;
  color: inherit;
  background-color: transparent;
  &:focus {
    border-color: red;
  }
`;

const OPTIONS = [
  {
    value: "vpc-0b6a762e556d8ee14",
    name: "vpc-0b6a762e556d8ee14 | AWS | VPC | ap-sw | Sandbox  ",
    cloud: "AWS",
    service: "VPC",
    region: "ap-sw",
    account: "Sandbox",
    cost: 19.94,
  },
  {
    value: "bmc0999",
    name: "bmc0999 | OCI | BareMetal Computing | ap-sydney-1 | Sandbox  ",
    cloud: "OCI",
    service: "BareMetal Computing",
    region: "ap-sydney-1",
    account: "Sandbox",
    cost: 17.00,
  },
];

const SelectBoxWrapper = styled.div`
	display: flex;
`;

const SelectBox = (props: any) => {
    	const handleChange = (e: any) => {
        // event handler
        console.log(e.target.value);
      };
  return (
    <SelectBoxWrapper>
      <Select>
        {props.options.map((option: any) => (
          <option value={option.value}>{option.name}</option>
        ))}
      </Select>
      <IconSVG
        width="20"
        height="20"
        viewBox="0 0 20 20"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          fill-rule="evenodd"
          clip-rule="evenodd"
          d="M10 14L16 6H4L10 14Z"
          fill="#1A1A1A"
        />
      </IconSVG>
    </SelectBoxWrapper>
  );
};

const IconSVG = styled.svg`
  margin-left: -28px;
  align-self: center;
  width: 24px;
  height: 24px;
`;

const ConnectLabel = styled.h1`
  font-size: 20px;
`;

interface ChildProps {
  closePrompt: () => void;
  AddInstFunction: () => void;
};

function AddInstancePrompt({closePrompt, AddInstFunction}:ChildProps) {
  const [AppRes, setAppRes] = useState<TldrawApp>();

    const randomString = (length = 8) => {
      const chars =
        "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
      let str = "";

      for (let i = 0; i < length; i++) {
        str += chars.charAt(Math.floor(Math.random() * chars.length));
      }

      return str;
    };

const handleClick = () => {
  closePrompt();
};

  const ButtonFunc = () => {
    doc.update((root) => {
      let Id = randomString(36);
      let InstanceIconId = randomString(36);
      let InstanceTextId = randomString(36);
      let InstanceCompletedId = randomString(36);
      let InstanceTextName = "InstanceName";
      let OriginSrc =
        "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAYAAABXAvmHAAAAAXNSR0IArs4c6QAAAERlWElmTU0AKgAAAAgAAYdpAAQAAAABAAAAGgAAAAAAA6ABAAMAAAABAAEAAKACAAQAAAABAAAAMKADAAQAAAABAAAAMAAAAADbN2wMAAABO0lEQVRoBe1U2w3CMAwMiA3YCqZhEqahWzEDNB9nWdE16UcTx5WRUNw4re9hJ6X4OVLg+0q//C8h994v6+nnq37wGAcBa9fcO3CpKbgO52fNP2pnBuSW+zs9t+q0HLAGn3FXMdy2mOn9VYGqU/rskTG7ssvvtxwoz0/3LA7sYWuFnmFDV5zHATDSKjPmOj8qZthQ270DQQBWWq1xC1kpj7riAJv0uIUgU8c1bqGO4u76tMzALP3OUDNsmFn3LSQOgJFWgDHX+VExw4ba7h0IArDSapUZmKXfmRAMG+bCfQuJA2CkFWDMdX5UzLChtnsHggCstFplBmbpdyYEw4a5cN9C4gAYaQUYc50fFTNsqH0eB8CIrbM4wbC1HFjYS4P3jsOQnWBu9N6vCdZyoPbuFLkgYG2DewesBYz6fzqMaA3IXYFbAAAAAElFTkSuQmCC";

      root.assets[InstanceIconId] = {
        id: InstanceIconId,
        name: "AWS_EC2_icon",
        size: [48, 48],
        src: OriginSrc,
        type: "image",
      };

      root.shapes[InstanceIconId] = {
        assetId: InstanceIconId,
        childIndex: 1,
        id: InstanceIconId,
        name: "Image",
        parentId: InstanceCompletedId,
        point: [1216.33, 755.14],
        size: [60, 60],
        rotation: 0,
        style: {
          color: "black",
          dash: "draw",
          isFilled: false,
          scale: 1,
          size: "small",
        },
        type: "image",
      };

      root.shapes[InstanceTextId] = {
        childIndex: 2,
        id: InstanceTextId,
        name: "Text",
        parentId: InstanceCompletedId,
        point: [1187.26, 847.38],
        rotation: 0,
        style: {
          color: "black",
          dash: "draw",
          font: "script",
          isFilled: false,
          scale: 0.940027,
          size: "small",
          textAlign: "middle",
        },
        text: "Instance Name",
        type: "text",
      };

      root.shapes[InstanceCompletedId] = {
        childIndex: 1,
        children: [InstanceIconId, InstanceTextId],
        id: InstanceCompletedId,
        name: "Group",
        parentId: "page",
        point: [1187.26, 755.14],
        rotation: 0,
        size: [138, 128.24],
        style: {
          color: "black",
          dash: "draw",
          isFilled: false,
          scale: 1,
          size: "small",
        },
        type: "group",
      };
    });

    const root = doc.getRoot();

    // Parse proxy object to record

    const shapeRecord: Record<string, TDShape> = JSON.parse(
      root.shapes.toJSON()
    );

    const bindingRecord: Record<string, TDBinding> = JSON.parse(
      root.bindings.toJSON()
    );
    const assetRecord: Record<string, TDAsset> = JSON.parse(
      root.assets.toJSON()
    );

    // Replace page content with changed(propagated) records
    //console.log("KORA" + root.shapes);
    AppRes?.replacePageContent(shapeRecord, bindingRecord, assetRecord);
  };


  return (
    <>
      <div className="AddAccountBackground">
        <button type="button" className="btn-close" onClick={handleClick}>
          <span className="icon-cross"></span>
          <span className="visually-hidden">Close</span>
        </button>
        <ConnectLabel>
          Select the Cloud Instance to add and express at the Tldraw board.
        </ConnectLabel>
        <br />
        <SelectBox options={OPTIONS}></SelectBox>
        <br />
        <br />
        <Button type="button" className="btn btn-primary" onClick={AddInstFunction}>
          OK
        </Button>
        <Button type="button" className="btn btn-secondary" onClick={handleClick}>
          Cancel
        </Button>
      </div>
    </>
  );
  }

export default AddInstancePrompt;
