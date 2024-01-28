const form = document.querySelector('form');
const submitMessage = document.querySelector('#submit-message');
const loader = document.getElementById('loader');
const inputs = form.querySelectorAll('input');
const checkbox = document.querySelector('.checkbox');
const menuItems = document.querySelector('.menu-items');

form.addEventListener('submit', (e) => {
  e.preventDefault();
  loader.style.display = 'block';
  submitMessage.classList.add('show');
 
  setTimeout(() => {
    form.submit()
    loader.style.display = 'none';
    form.reset();

  }, 2000);
});

inputs.forEach((input) => {
  input.addEventListener('focus', () => {
    submitMessage.classList.remove('show');
  });
});


document.addEventListener('click', (event) => {
  if (
    !checkbox.contains(event.target) &&
    !menuItems.contains(event.target) &&
    menuItems.classList.contains('open')
  ) {
    checkbox.checked = false;
  }
});

checkbox.addEventListener('change', () => {
  menuItems.classList.toggle('open', checkbox.checked);
});