import { Events } from "@wailsio/runtime";
import {initTabs } from "./tabs.js"
import { initAddServer } from "./add_server_tab.js";
import { initServersTab } from "./servers_tab.js";
// import { API } from "../bindings/mvp/backend/api/api.js";

console.log("Frontend Loaded!");


document.addEventListener("DOMContentLoaded", () => {
  initTabs();          // Handles tab switching
  initAddServer();
  initServersTab();
});
