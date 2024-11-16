import viteLogo from "/vite.svg";
import "./App.css";
import { useEffect } from "react";
import { useState } from "react";
import axios from 'axios';

const BE_API = "http://localhost:81/api/be-app2/data"

const getDataUser = async (token, onSuccess, onError) => {
  try {
    const response = await axios.get(BE_API, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (onSuccess) onSuccess(response);
  } catch (err) {
    if (onError) onError(err)
  }
}

function App() {
  const [username, setUsername] = useState('');

  const navigateBack = () => window.location.replace("http://localhost:81");

  useEffect(() => {
    const queryParams = new URLSearchParams(window.location.search);
    const token = queryParams.get('token');

    getDataUser(
      token,
      (response) => {
        if (response.status === 200) {
          setUsername(response.data);
        }
      },
      (err) => {
        console.log(err);
      }
    )
  }, []);

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
      </div>
      <h1>WELCOME TO APP2</h1>
      <div className="text-info">
        {username}
      </div>
      <div className="card">
        <button onClick={navigateBack}>Logout</button>
      </div>
      <p className="read-the-docs">@Ginsebu</p>
    </>
  );
}

export default App;
