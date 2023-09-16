import logo from "./logo.svg";
import "./App.css";
import { GoogleLogin } from "@react-oauth/google";
import jwt_decode from "jwt-decode";

function App() {
  const onSuccess = (credentialResponse) => {
    console.log("credentialResponse", credentialResponse);
    const token = jwt_decode(credentialResponse.credential);
    // this credential should be sent in requests to the backend
    // https://developers.google.com/identity/gsi/web/guides/verify-google-id-token
    console.log("token", token);
  };

  const onFailure = (error) => {
    console.log(error);
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <a href="https://www.npmjs.com/package/@react-oauth/google">Docs</a>
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
        <div>
          <GoogleLogin onSuccess={onSuccess} onFailure={onFailure} />
        </div>
      </header>
    </div>
  );
}

export default App;
