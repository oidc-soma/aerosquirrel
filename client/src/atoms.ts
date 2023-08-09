import { atom } from "recoil";

export const InventoryAtom = atom({
  key: "Inventory",
  default: { resources: [{}] },
});

export const CSPAtom = atom({
  key: "CSP",
  default: {
    aws: [{ name: "" }, { token: "" }],
    oci: [{ name: "" }, { token: "" }],
    k8s: [{ name: "" }, { token: "" }],
  },
});
