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
   const parentFunction = (x: any) => {
     console.log(x);
    //createShapes(x, 
   };
  const fileSystemEvents = useFileSystem();
  const { ...events } = useMultiplayerState(
    // `tldraw-${new Date().toISOString().substring(0, 10).replace(/-/g, "")}`
    'tldraw-Savers', {parentFunction}
  );
  const ButtonFunc = () => {
    doc.update((root) => {
      root.shapes = {
        "3df3c015-07dc-4b6d-2b97-15ab54d4ba41": {
          id: "3df3c015-07dc-4b6d-2b97-15ab54d4ba41",
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
        },
      };
    })
  }
  
 

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