// servers.js
import {
  GetServers,
  GetLatestServerMetrics
} from "../bindings/mvp/backend/api/api.js";

let lastSnapshot = null;
let pollInterval = null;

export function initServersTab(intervalMs = 3000) {
  console.log("initServersTab running");

  const container = document.getElementById("server-rows");
  if (!container) {
    console.error("server-rows container not found");
    return;
  }

  container.innerHTML = `<p class="text-gray-400">Loading servers...</p>`;

  async function fetchAndRender() {
    try {
      const servers = await GetServers();
      const snapshot = JSON.stringify(servers);

      // Only re-render if server list changed
      if (snapshot === lastSnapshot) return;
      lastSnapshot = snapshot;

      container.innerHTML = "";

      if (!servers || servers.length === 0) {
        container.innerHTML = `
          <p class="text-gray-400">No servers added yet.</p>
        `;
        return;
      }

      for (const server of servers) {
        const row = document.createElement("div");
        row.className =
          "p-3 mb-2 bg-gray-800 rounded border border-gray-700";

        // Base layout first (fast render)
        row.innerHTML = `
          <div class="flex items-center justify-between">
            <div>
              <p class="font-semibold text-white">${server.name}</p>
              <p class="text-sm text-gray-400">${server.ip}</p>
            </div>
            <button class="px-3 py-1 text-sm bg-gray-700 hover:bg-gray-600 rounded">
              PEM
            </button>
          </div>

          <div class="mt-2 text-sm text-gray-400 metrics">
            Loading metrics…
          </div>
        `;

        container.appendChild(row);

        // Fetch metrics async per server
        try {
          const data = await GetLatestServerMetrics(server.name);
          const m = data.metrics;

          row.querySelector(".metrics").innerHTML = `
            <div class="grid grid-cols-3 gap-4 mt-1 text-sm">
              <div>
                <p class="text-gray-400">CPU</p>
                <p class="text-white font-medium">${m.cpu.usage}%</p>
              </div>
              <div>
                <p class="text-gray-400">RAM</p>
                <p class="text-white font-medium">
                  ${m.ram.active_ram_percent}%
                </p>
              </div>
              <div>
                <p class="text-gray-400">Disk</p>
                <p class="text-white font-medium">
                  ${m.disk.used_percent}%
                </p>
              </div>
            </div>

            <p class="mt-1 text-xs text-gray-500">
              Updated: ${new Date(data.timestamp).toLocaleTimeString()}
            </p>
          `;
        } catch (err) {
          row.querySelector(".metrics").innerHTML = `
            <p class="text-yellow-400 text-sm">
              No metrics available yet
            </p>
          `;
        }
      }

    } catch (err) {
      console.error("Failed to load servers:", err);
      container.innerHTML = `
        <p class="text-red-400">Failed to load servers</p>
      `;
    }
  }

  // Initial fetch
  fetchAndRender();

  // Restart polling cleanly
  if (pollInterval) clearInterval(pollInterval);
  pollInterval = setInterval(fetchAndRender, intervalMs);
}
