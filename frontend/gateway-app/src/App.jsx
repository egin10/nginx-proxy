import viteLogo from "/vite.svg";
import "./App.css";

function App() {
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

  const navigateTo = (app) => {
    if (Object.keys(app).length > 0 && app.url)
      window.location.replace(app.url);
  };

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
      </div>
      <h1>Gateway App</h1>
      <div className="card">
        {listApps.map((app, index) => (
          <button className="btn-navigation" key={index} onClick={() => navigateTo(app)}>
            {app.name.toUpperCase()}
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
