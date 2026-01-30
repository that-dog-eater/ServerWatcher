// add_server.js
import { AddServer } from "../bindings/mvp/backend/api/api.js";

/**
 * Initialize Add Server form watcher and submission
 */
export function initAddServer() {
  console.log("initAddServer func running");

  const nameInput = document.getElementById("name");
  const ipInput = document.getElementById("ip");
  const pemInput = document.getElementById("pemKey");
  const addButton = document.querySelector(".btn-primary");
  const resultContainer = document.getElementById("result");
  const serverRows = document.getElementById("server-rows");

  if (!nameInput || !ipInput || !pemInput || !addButton || !resultContainer) {
    console.error("Add Server form elements missing");
    return;
  }

  // Enable/disable button based on form completion
  function checkForm() {
    if (nameInput.value.trim() && ipInput.value.trim() && pemInput.value.trim()) {
      addButton.disabled = false;
      addButton.classList.remove("opacity-50", "cursor-not-allowed");
    } else {
      addButton.disabled = true;
      addButton.classList.add("opacity-50", "cursor-not-allowed");
    }
  }

  // Watch input changes
  [nameInput, ipInput, pemInput].forEach(input => {
    input.addEventListener("input", checkForm);
  });

  // Initial check
  checkForm();

  // Add server handler
  async function addServerHandler() {
    const name = nameInput.value.trim();
    const ip = ipInput.value.trim();
    const pemKey = pemInput.value.trim();

    try {
      // Send three separate strings to Go
      await AddServer(name, ip, pemKey);

      // Show success
      resultContainer.innerHTML = `<p class="text-green-400">Server added!</p>`;

      // Clear form
      nameInput.value = "";
      ipInput.value = "";
      pemInput.value = "";
      checkForm(); // disable button

    } catch (err) {
      console.error("AddServer error:", err);

      let message = "Unknown error";

      // Case 1: err.message exists
      if (err?.message) {
        try {
          // Try to parse JSON string from Wails
          const parsed = JSON.parse(err.message);
          message = parsed.message || err.message;
        } catch {
          // Not JSON, use as-is
          message = err.message;
        }
      } 
      // Case 2: err is a string
      else if (typeof err === "string") {
        message = err;
      } 
      // Case 3: fallback
      else {
        message = "Unexpected error occurred";
      }

      resultContainer.innerHTML = `
        <p class="text-red-400">Failed to add server: ${message}</p>
      `;
    }
  }
  // Attach click handler
  addButton.addEventListener("click", addServerHandler);
}
