const form = document.querySelector('form');
const submitMessage = document.querySelector('#submit-message');

form.addEventListener('submit', (e) => {
  e.preventDefault();
  submitMessage.classList.add('show');
 
  setTimeout(() => {
    form.submit()
    form.reset();
  }, 2000);
});