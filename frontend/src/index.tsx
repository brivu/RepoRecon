import React, { useState } from 'react';
import ReactDOM from 'react-dom/client';

const axios = require('axios');

const App: React.FC = () => {
  const [text, setText] = useState('');

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setText(event.target.value);
  };

  const handleSubmit = async (event: React.FormEvent) => {
    
    axios.post('http://localhost:3000/url', {
        url: text
    })
    .then( (response: any) => {
        console.log(response)
    })
    .catch( (error: any) => {
        console.log(error)
    })


    event.preventDefault();
    setText('');
  };

  return (
    <div>
      <h1>Hello World</h1>
      <form onSubmit={handleSubmit}>
        <input type="text" value={text} onChange={handleChange} />
        <button type="submit">Submit</button>
      </form>
      <p>You entered: {text}</p>
    </div>
  );
};

const root = document.getElementById('root');
if (root !== null) {
    (ReactDOM as any).createRoot(root).render(<App />);

}