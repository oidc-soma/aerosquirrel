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

export default function Editor() {
  const [AppRes, setAppRes] = useState<TldrawApp>();
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
    'tldraw-storage', {parentFunction}, {appExport}
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
    
    
    doc.update((root) => {
      let Id = randomString(36);
      let InstanceIconId = randomString(36);
      let InstanceTextId = randomString(36);
      let InstanceCompletedId = randomString(36);
      let InstanceTextName = "InstanceName";
      let OriginSrc =
        "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAYAAABXAvmHAAAAAXNSR0IArs4c6QAAAERlWElmTU0AKgAAAAgAAYdpAAQAAAABAAAAGgAAAAAAA6ABAAMAAAABAAEAAKACAAQAAAABAAAAMKADAAQAAAABAAAAMAAAAADbN2wMAAABO0lEQVRoBe1U2w3CMAwMiA3YCqZhEqahWzEDNB9nWdE16UcTx5WRUNw4re9hJ6X4OVLg+0q//C8h994v6+nnq37wGAcBa9fcO3CpKbgO52fNP2pnBuSW+zs9t+q0HLAGn3FXMdy2mOn9VYGqU/rskTG7ssvvtxwoz0/3LA7sYWuFnmFDV5zHATDSKjPmOj8qZthQ270DQQBWWq1xC1kpj7riAJv0uIUgU8c1bqGO4u76tMzALP3OUDNsmFn3LSQOgJFWgDHX+VExw4ba7h0IArDSapUZmKXfmRAMG+bCfQuJA2CkFWDMdX5UzLChtnsHggCstFplBmbpdyYEw4a5cN9C4gAYaQUYc50fFTNsqH0eB8CIrbM4wbC1HFjYS4P3jsOQnWBu9N6vCdZyoPbuFLkgYG2DewesBYz6fzqMaA3IXYFbAAAAAElFTkSuQmCC";


      root.assets[InstanceIconId] = {
          "id": InstanceIconId,
          "name": "AWS_EC2_icon",
          "size": [
            48,
            48
          ],
          "src": OriginSrc,
          "type": "image",
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
        type: "image"
      };

      root.shapes[InstanceTextId] = {
        childIndex: 2,
        id: InstanceTextId,
        name: "Text",
        parentId: InstanceCompletedId,
        point: [
          1187.26,
          847.38
        ],
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
      }
      
      root.shapes[InstanceCompletedId] = {
        childIndex: 1,
        children: [
          InstanceIconId,
          InstanceTextId,
        ],
        id: InstanceCompletedId,
        name: "Group",
        parentId: "page",
        point: [
          1187.26,
          755.14
        ],
        rotation: 0,
        size: [
          138,
          128.24
        ],
        style: {
          color: "black",
          dash: "draw",
          isFilled: false,
          scale: 1,
          size: "small",
        },
        type: "group"
      }

    })

    
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
  
 

  const component = { 
    Cursor: CustomCursor };

  return (
    <>
      <Button
        variant="outline-primary"
        style={{
          float: "right",
        }}
      >
        Yorkie Configure
      </Button>
      <Button
        variant="outline-primary"
        onClick={ButtonFunc}
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
      </div>
    </>
  );
}