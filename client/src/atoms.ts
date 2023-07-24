import { atom } from "recoil";

export const InventoryAtom = atom({
  key: "Inventory",
  default: { resources: [{}] },
});
