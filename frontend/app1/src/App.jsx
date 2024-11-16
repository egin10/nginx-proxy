import viteLogo from "/vite.svg";
import "./App.css";

function App() {
  const navigateBack = () => window.location.replace("http://localhost:3000");

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
      </div>
      <h1>WELCOME TO APP1</h1>
      <div className="card">
        <button onClick={navigateBack}>Logout</button>
      </div>
      <p className="read-the-docs">@Ginsebu</p>
    </>
  );
}

export default App;
