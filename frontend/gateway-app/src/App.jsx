import viteLogo from "/vite.svg";
import "./App.css";
import { useState } from "react";
import axios from 'axios';

const BE_API = "http://localhost:81/api/be-app1/generate-token?username="

const postDataUsername = async (username, onSuccess, onError) => {
  try {
    const response = await axios.post(BE_API + username);
    if (onSuccess) onSuccess(response);
  } catch (err) {
    if (onError) onError(err)
  }
}

function App() {
  const [username, setUsername] = useState('');
  const [error, setError] = useState({ status: false, message: '' });
  const [isLoading, setIsLoading] = useState(false);

  const listApps = [
    {
      name: "app1",
      url: "http://localhost:81/app1/",
    },
    {
      name: "app2",
      url: "http://localhost:81/app2/",
    },
  ];

  const handleOnChange = (e) => setUsername(e.target.value.replaceAll(" ", ""));

  const handleSignIn = (app) => {
    if (username === '') {
      setError({ status: true, message: 'Username is required.' });
      return;
    }

    if (username.length > 0 && username.length < 5) {
      setError({ status: true, message: 'Username min 5 characters.' });
      return;
    }

    setError({ status: false, message: '' });
    setIsLoading(true);

    postDataUsername(
      username,
      (response) => {
        setTimeout(() => setIsLoading(false), 1000);

        if (response.status === 200) {
          navigateTo(app, response.data);
        }
      },
      (err) => {
        console.log(err);
      }
    );

  }

  const navigateTo = (app, token) => {
    if (Object.keys(app).length > 0 && app.url)
      window.location.replace(`${app.url}?token=${token}`);
  };

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
      </div>
      <h1>Gateway App</h1>
      <div className="column">
        <input type="text" placeholder="Username" className="text-input"
          value={username}
          onChange={handleOnChange} />

        {error.status && <span className="error-message">{error.message}</span>}

        {isLoading ? <span className="text-loading">Please wait...</span> : <></>}
      </div>

      <div className="row">
        {listApps.map((app, index) => (
          <button className="btn-navigation"
            key={index}
            onClick={() => handleSignIn(app)}
            disabled={isLoading}>
            SIGN IN TO {app.name.toUpperCase()}
          </button>
        ))}
      </div>
      <p className="read-the-docs">
        @Ginsebu
      </p>
    </>
  );
}

export default App;
