window.scrollToSection = function(id) {
  const el = document.getElementById('sec-' + id);
  if (el) el.scrollIntoView({ behavior: 'smooth' });
};
