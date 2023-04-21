import React, { useState } from "react";
import axios from "axios";
import "./App.css";

const App = () => {
  const [message, setMessage] = useState("");
  const [userInput, setUserInput] = useState("");

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUserInput(event.target.value);
  };

  const handleButtonClick = async () => {
    const response = await axios.get(`http://0.0.0.0:8080/api/${userInput}`);
    setMessage(response.data.message);
    setUserInput("");
  };

  return (
    <div className="container">
      <input type="text" value={userInput} onChange={handleInputChange} />
      <button className="button" onClick={handleButtonClick}>ENVOIE</button>
      <h1>API ROBIX : {message}</h1>
    </div>
  );

};

export default App;
