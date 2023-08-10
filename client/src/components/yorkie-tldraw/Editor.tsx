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
import { createShapes } from "@tldraw/tldraw/dist/state/commands";
import { useMultiplayerState } from '../../hooks/useMultiplayerState';
import CustomCursor from '../../CursorComponent';
import { useState, useCallback } from 'react';
import { doc } from '../../hooks/useMultiplayerState';
import Button from "react-bootstrap/Button";
import "bootstrap/dist/css/bootstrap.min.css";
import AddInstancePrompt from "../../pages/AddInstancePrompt";
import {
  useRecoilState,
  useRecoilValue,
  useResetRecoilState,
  useSetRecoilState,
} from "recoil";
import { InstanceNameAtom } from "../../atoms";
import {toast} from 'react-toastify';

export default function Editor() {
     const [GetInstanceName, SetInstanceName] =
       useRecoilState(InstanceNameAtom);
     const InstanceName = GetInstanceName;
  const [AppRes, setAppRes] = useState<TldrawApp>();

    const [showAddInstancePopup, setshowAddInstancePopup] = useState(false);

    const showAddInstancePopupFunction = () => {
      setshowAddInstancePopup(true);
    };

    const closeAddInstancePopupFunction = () => {
      setshowAddInstancePopup(false);
    };

  const appExport = (app: TldrawApp) => {
    setAppRes(app);
  }


   const parentFunction = (x: any) => {
     console.log(x);
    //createShapes(x, 
   };
  const fileSystemEvents = useFileSystem();
  const { ...events } = useMultiplayerState(
    // `tldraw-${new Date().toISOString().substring(0, 10).replace(/-/g, "")}`
    'tldraw-storage_k', {parentFunction}, {appExport}
  );

  const randomString = (length = 8) => {
    const chars =
      "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    let str = "";

    for (let i = 0; i < length; i++) {
      str += chars.charAt(Math.floor(Math.random() * chars.length));
    }

    return str;
  };
  const ButtonFunc = () => {
    if(InstanceName!=="null")
    {
          
    doc.update((root) => {
      let Id = randomString(36);
      let InstanceIconId = randomString(36);
      let InstanceTextId = randomString(36);
      let InstanceCompletedId = randomString(36);
      let InstanceTextName = "InstanceName";
      let OriginSrc =
        "data:image/png;base64,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";

      root.assets[InstanceIconId] = {
        id: InstanceIconId,
        name: InstanceName+"_ICON",
        size: [65, 65],
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
        size: [80, 80],
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
          font: "sans",
          isFilled: false,
          scale: 0.940027,
          size: "small",
          textAlign: "middle",
        },
        text: InstanceName,
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
    

    }
    else {
      toast("Please Select the Instance", {
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
  
 

  const component = { 
    Cursor: CustomCursor };

  return (
    <>
      {/* <Button
        variant="outline-primary"
        style={{
          float: "right",
        }}
      >
        Yorkie Configure
      </Button> */}
      <Button
        variant="outline-primary"
        onClick={showAddInstancePopupFunction}
        style={{
          float: "right",
        }}
      >
        Add and Draw New Instance
      </Button>
      <div
        className="tldraw"
        style={{
          position: "absolute",
          top: "40px",
          width: "100%",
          height: "calc( 100% - 40px)",
        }}
      >
        <Tldraw
          components={component}
          autofocus
          disableAssets={false}
          showPages={false}
          {...fileSystemEvents}
          {...events}
        />
        {showAddInstancePopup && (
          <AddInstancePrompt closePrompt={closeAddInstancePopupFunction} AddInstFunction={ButtonFunc} />
        )}
      </div>
    </>
  );
}
