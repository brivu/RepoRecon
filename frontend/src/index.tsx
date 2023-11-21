import React, { useState } from 'react';
import ReactDOM from 'react-dom';

const App: React.FC = () => {
  const [text, setText] = useState('');

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setText(event.target.value);
  };

  const handleSubmit = async (event: React.FormEvent) => {
    
    const response = await fetch('http://localhost:8080/url', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ url: text }),
      });
    
      if (!response.ok) {
        console.error('Failed to save URL');
        return;
      }
    
    event.preventDefault();
    console.log(text);
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

ReactDOM.render(<App />, document.getElementById('root'));