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
    'tldraw-Savere', {parentFunction}, {appExport}
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
      root.shapes[Id] = {
        
          id: Id,
          type: "rectangle",
          name: "Rectangle",
          parentId: "page",
          childIndex: 2,
          point: [482, -267.56],
          size: [119.32, 129.24],
          rotation: 0,
          style: {
            color: "black",
            size: "small",
            isFilled: false,
            dash: "draw",
            scale: 1,
          },
          label: "",
          labelPoint: [0.5, 0.5],
        };



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
    
      <button style={{ zIndex: 10000}} onClick={ButtonFunc}>테스트용 버튼</button>
      <div className="tldraw" style={{position: 'absolute', top: '200px', width: '100%', height: '600px'}}>
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