const form = document.querySelector('form');
const submitMessage = document.querySelector('#submit-message');
const loader = document.getElementById('loader');


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