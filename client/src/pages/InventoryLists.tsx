import React from "react";
import "./InventoryLists.css";

interface ChildProps {
  data: { resources: any };
}

function InventoryLists({ data }: ChildProps) {
  return (
    <div className="invouterlists">
      <table
        id="InventoryTable"
        className="table-striped table-bordered"
        cellSpacing="0"
        width="100%"
      >
        <thead>
          <tr>
            <th className="th-sm">ID</th>
            <th className="th-sm">UserID</th>
            <th className="th-sm">Name</th>
            <th className="th-sm">Type</th>
            <th className="th-sm">Cost</th>
            <th className="th-sm">Metadata</th>
            <th className="th-sm">Tags</th>
            <th className="th-sm">Link</th>
          </tr>
        </thead>
        <tbody>
          {data.resources.map((item: any) => (
            <tr key={item.id}>
              <td>{item.id}</td>
              <td>{item.user_id}</td>
              <td>{item.name}</td>
              <td>{item.type}</td>
              <td>{item.cost}</td>
              <td>
                {item.metadata &&
                  Object.entries(item.metadata).map(([key, value]: any) => (
                    <div key={key}>
                      {Object.entries(value).map(([vkey, vvalue]: any) => (
                        <div key={key}>
                          <span className="metakey">
                            {JSON.stringify(vkey).replace(/"/g, "")}
                          </span>
                          <span className="metavalue">
                            {JSON.stringify(vvalue).replace(/"/g, "")}
                          </span>
                        </div>
                      ))}
                    </div>
                  ))}
              </td>
              {/* <td>{JSON.stringify(item.tags)}</td> */}
              <td>
                {item.tags &&
                  Object.entries(item.tags).map(([key, value]: any) => (
                    <div key={key}>
                      {Object.entries(value).map(([vkey, vvalue]: any) => (
                        <div key={key}>
                          <span className="tagkey">
                            {JSON.stringify(vkey).replace(/"/g, "")}
                          </span>
                          <span className="tagvalue">
                            {JSON.stringify(vvalue).replace(/"/g, "")}
                          </span>
                        </div>
                      ))}
                    </div>
                  ))}
              </td>
              <td>
                <a href={item.link}>{item.link}</a>
              </td>
            </tr>
          ))}
        </tbody>
        <tfoot></tfoot>
      </table>
    </div>
  );
}
export default InventoryLists;
