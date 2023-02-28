import { useState } from 'react';
import './App.css';
import { Greet } from "../wailsjs/go/main/App";

import TopNav from './TopNav';
import CardGroup from './CardGroup';

function App() {
    const [resultText, setResultText] = useState("Fase 1 👇");
    const [name, setName] = useState('');
    const updateName = (e) => setName(e.target.value);
    const updateResultText = (result) => setResultText(result);

    function greet() {
        Greet(name).then(updateResultText);
    }

    function holi(){
        return Prueba()
    }

    return (
        <div id="App">
            <TopNav />
            <div className='p-5 text-center bg-light'>
                <h1 className='mb-3'> Proyecto 1, 201480017 {} da</h1>
            </div>
            <CardGroup />
            <div id="result" className="result">{resultText}d</div>
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text" />
                <button className="btn" style={{ color: "info" }} onClick={greet}>Say Hi</button>
            </div>
        </div>
    )
}

export default App
