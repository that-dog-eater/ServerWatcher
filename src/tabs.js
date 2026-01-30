// tabs.js
export function initTabs() {
  const navLinks = document.querySelectorAll('.nav-link');
  const tabs = document.querySelectorAll('.tab-content');

  navLinks.forEach(link => {
    link.addEventListener('click', e => {
      e.preventDefault();

      // Toggle active class
      navLinks.forEach(l => l.classList.remove('active'));
      link.classList.add('active');

      // Show the correct tab
      const targetId = link.getAttribute('href').substring(1);
      tabs.forEach(tab => tab.classList.add('hidden'));
      const targetTab = document.getElementById(targetId);
      if (targetTab) targetTab.classList.remove('hidden');
    });
  });
}
